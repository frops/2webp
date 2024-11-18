package converter

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToWebP(t *testing.T) {
	// Create a temporary file for JPG
	tmpDir := t.TempDir()
	sourcePath := filepath.Join(tmpDir, "test.jpg")
	destPath := filepath.Join(tmpDir, "test.webp")

	// Generate a test image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 255, 255})
		}
	}

	// Save as JPG
	srcFile, err := os.Create(sourcePath)
	require.NoError(t, err)
	defer srcFile.Close()

	err = jpeg.Encode(srcFile, img, nil)
	require.NoError(t, err)

	// Test conversion to WebP
	err = convertToWebP(sourcePath, destPath)
	require.NoError(t, err)

	// Ensure the WebP file is created
	_, err = os.Stat(destPath)
	require.NoError(t, err)
}

// func TestCollectJobs(t *testing.T) {
// 	tmpDir := t.TempDir()

// 	// Create a directory structure with files
// 	jpgFile := filepath.Join(tmpDir, "image.jpg")
// 	webpFile := filepath.Join(tmpDir, "image.webp")
// 	err := os.WriteFile(jpgFile, []byte("test"), 0644)
// 	require.NoError(t, err)
// 	err = os.WriteFile(webpFile, []byte("webp"), 0644)
// 	require.NoError(t, err)

// 	destDir := filepath.Join(tmpDir, "output")
// 	err = os.Mkdir(destDir, 0755)
// 	require.NoError(t, err)

// 	// Ensure that without --force, existing WebP files are skipped
// 	jobs, err := collectJobs(tmpDir, destDir, false)
// 	require.NoError(t, err)
// 	require.Len(t, jobs, 0)

// 	// Ensure that with --force, all files are added as jobs
// 	jobs, err = collectJobs(tmpDir, destDir, true)
// 	require.NoError(t, err)
// 	require.Len(t, jobs, 1)
// 	require.Equal(t, jpgFile, jobs[0].SourcePath)
// }

func TestProcessDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	sourceDir := filepath.Join(tmpDir, "src")
	destDir := filepath.Join(tmpDir, "dst")

	// Create directories and a test file
	err := os.MkdirAll(sourceDir, 0755)
	require.NoError(t, err)

	jpgFile := filepath.Join(sourceDir, "image.jpg")
	img := image.NewRGBA(image.Rect(0, 0, 50, 50))
	file, err := os.Create(jpgFile)
	require.NoError(t, err)

	err = jpeg.Encode(file, img, nil)
	require.NoError(t, err)
	file.Close()

	// Test processing the directory
	err = ProcessDirectory(sourceDir, destDir, 2, false)
	require.NoError(t, err)

	// Ensure the WebP file is created
	webpFile := filepath.Join(destDir, "image.webp")
	_, err = os.Stat(webpFile)
	require.NoError(t, err)

	// Test with an invalid source directory
	err = ProcessDirectory("/invalid-path", destDir, 2, false)
	require.Error(t, err)
}
