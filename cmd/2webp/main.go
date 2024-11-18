package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/frops/2webp/internal/converter"
)

func main() {
	sourceDir := flag.String("src", "", "Source directory containing image files (required)")
	destDir := flag.String("dst", "", "Destination directory for WebP files (required)")
	force := flag.Bool("force", false, "Force conversion even if WebP files already exist")
	threads := flag.Int("threads", 4, "Number of concurrent workers")

	flag.Usage = func() {
		fmt.Println("2webp: Convert images to WebP format")
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *sourceDir == "" || *destDir == "" {
		flag.Usage()
		log.Fatal("Both -src and -dst flags are required")
	}

	err := converter.ProcessDirectory(*sourceDir, *destDir, *threads, *force)
	if err != nil {
		log.Fatalf("Error processing directory: %v", err)
	}
}
