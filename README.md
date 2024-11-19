
# 2webp

A simple CLI tool to convert images (currently JPG/JPEG) to WebP format. Fast, multithreaded, and easy to use. No CGO dependencies required.

## Installation

### Option 1: Install via `go install`

Install the tool globally using `go install`:

```bash
go install github.com/frops/2webp/cmd/2webp@latest
```

Make sure `$GOPATH/bin` or `$HOME/bin` is in your `PATH`.

### Option 2: Run with Docker

You can also run `2webp` using Docker, which ensures all dependencies are packaged:

1. **Build the Docker image:**
   ```bash
   docker build -t 2webp .
   ```

2. **Run the container with your images:**
   ```bash
   docker run --rm -v /path/to/source:/input -v /path/to/destination:/output 2webp -src /input -dst /output
   ```

---

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

4. **Run via Docker:**
   ```bash
   docker run --rm -v /local/source:/input -v /local/destination:/output 2webp -src /input -dst /output -threads 4
   ```

---

## Features

- **Multithreaded**: Speed up conversions with the `-threads` option.
- **Skip existing files**: Avoid reconversion unless `-force` is specified.
- **No CGO**: Runs without external C dependencies for easier installation.
- **Docker Support**: Fully packaged with all dependencies using Docker.
- **Easy to install**: No external libraries needed for `go install`.

---

Start converting your images to WebP with a single command!
