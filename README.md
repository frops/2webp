
# 2webp

A simple CLI tool to convert images (currently JPG/JPEG) to WebP format. Fast, multithreaded, and easy to use. Without CGO.

## Installation

Install the tool globally using `go install`:

```bash
go install github.com/yourusername/2webp/cmd/2webp@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Usage

Run the `2webp` command with the following options:

```bash
2webp -src /path/to/source -dst /path/to/destination [options]
```

### Options
- `-src`: Source directory containing images (required).
- `-dst`: Destination directory for WebP files (required).
- `-threads`: Number of concurrent workers (default: 4).
- `-force`: Reconvert files even if they already exist.

### Examples

1. **Convert all images from `input` to `output`:**
   ```bash
   2webp -src ./input -dst ./output
   ```

2. **Force reconversion of all images:**
   ```bash
   2webp -src ./input -dst ./output -force
   ```

3. **Use 8 threads for faster processing:**
   ```bash
   2webp -src ./input -dst ./output -threads 8
   ```

## Features

- **Multithreaded**: Speed up conversions with the `-threads` option.
- **Skip existing files**: Avoid reconversion unless `-force` is specified.
- **Easy to install**: No extra dependencies required, install with `go install`.

---

Start converting your images to WebP with a single command!
