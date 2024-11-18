package converter

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/chai2010/webp"
	"github.com/schollz/progressbar/v3"
)

// Job represents a single file conversion task
type Job struct {
	SourcePath string
	DestPath   string
}

// collectJobs scans the source directory and prepares a list of conversion jobs
func collectJobs(sourceDir, destDir string, force bool) ([]Job, error) {
	var jobs []Job

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a JPG or JPEG file
		if !info.IsDir() && (strings.HasSuffix(strings.ToLower(info.Name()), ".jpg") || strings.HasSuffix(strings.ToLower(info.Name()), ".jpeg")) {
			// Determine destination path
			relPath, _ := filepath.Rel(sourceDir, path)
			destPath := filepath.Join(destDir, strings.TrimSuffix(relPath, filepath.Ext(relPath))+".webp")

			// Skip if file exists and force is not enabled
			if !force {
				if _, err := os.Stat(destPath); err == nil {
					return nil
				}
			}

			// Add the job to the list
			jobs = append(jobs, Job{
				SourcePath: path,
				DestPath:   destPath,
			})
		}

		return nil
	})

	return jobs, err
}

func convertToWebP(sourcePath, destPath string) error {
	// Open the source file
	srcFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Decode the image
	img, err := jpeg.Decode(srcFile)
	if err != nil {
		return err
	}

	// Create the destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Encode the image as WebP
	return webp.Encode(destFile, img, &webp.Options{Lossless: false, Quality: 75})
}

func ProcessDirectory(sourceDir, destDir string, threads int, force bool) error {
	// Collect jobs
	jobs, err := collectJobs(sourceDir, destDir, force)
	if err != nil {
		return err
	}

	// No jobs to process
	if len(jobs) == 0 {
		log.Println("No files to process. Exiting.")
		return nil
	}

	// Progress bar
	bar := progressbar.Default(int64(len(jobs)))

	// Start workers
	jobChan := make(chan Job, len(jobs))
	errorChan := make(chan error, len(jobs))
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobChan {
				// Ensure the destination directory exists
				destDir := filepath.Dir(job.DestPath)
				if err := os.MkdirAll(destDir, 0755); err != nil {
					errorChan <- err
					continue
				}

				// Convert to WebP
				if err := convertToWebP(job.SourcePath, job.DestPath); err != nil {
					errorChan <- err
				}
				bar.Add(1)
			}
		}()
	}

	// Enqueue jobs
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	// Wait for workers to finish
	wg.Wait()
	close(errorChan)

	// Collect errors
	var errors []error
	for err := range errorChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return combineErrors(errors)
	}

	log.Println("All files processed successfully.")
	return nil
}

// combineErrors aggregates multiple errors into one
func combineErrors(errors []error) error {
	var combinedErr string
	for _, err := range errors {
		combinedErr += err.Error() + "\n"
	}
	return fmt.Errorf("multiple errors occurred:\n%s", combinedErr)
}
