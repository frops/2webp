package main

import (
	"flag"
	"log"

	"github.com/frops/2webp/internal/converter"
)

func main() {
	sourceDir := flag.String("src", "", "Source directory containing JPG files (required)")
	destDir := flag.String("dst", "", "Destination directory for WebP files (required)")
	force := flag.Bool("force", false, "Force conversion even if WebP files already exist")
	threads := flag.Int("threads", 4, "Number of concurrent workers")
	flag.Parse()

	if *sourceDir == "" || *destDir == "" {
		flag.Usage()
		log.Fatal("Both -src and -dst flags are required")
	}

	converter.ProcessDirectory(*sourceDir, *destDir, *threads, *force)
}
