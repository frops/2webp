
# 2webp

A simple CLI tool to convert images (currently JPG/JPEG) to WebP format. Fast, multithreaded, and easy to use

### Run with Docker

You can also run `2webp` using Docker, which ensures all dependencies are packaged:

1. **Build the Docker image:**
   ```bash
   docker build -t 2webp .
   ```

2. **Run the container with your images:**
   ```bash
   docker run --rm -v /path/to/source:/input -v /path/to/destination:/output frops/2webp -src /input -dst /output
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

## Push to Docker Hub

1. `docker login`
2. `docker tag 2webp frops/2webp`
3. `docker push frops/2webp`

---

## Features

- **Multithreaded**: Speed up conversions with the `-threads` option.
- **Docker Support**: Fully packaged with all dependencies using Docker.
- **Skip existing files**: Avoid reconversion unless `-force` is specified.

---

Start converting your images to WebP with a single command!
