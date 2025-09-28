# File Size Aggregator by Extension

A Go program that recursively scans a directory and calculates the total size of files grouped by their extensions.

## Description

This utility walks through a specified directory (and all its subdirectories) to analyze file sizes based on their extensions. It provides a clear summary showing:

- Total size per file extension
- Percentage of total size for each extension
- Human-readable file sizes (B, KB, MB, GB, etc.)
- Files are sorted by total size in descending order

## Installation

```bash
go build -o fexa main.go
```

Or run directly:

```bash
go run main.go [directory]
```

## Usage

### Basic usage (scan current directory):

```bash
./fexa
```

### Scan specific directory:

```bash
./fexa /path/to/directory
```

## Output Example

```
Scanning directory: /home/user/projects
--------------------------------------------------
Extension                      Size Percentage
--------------------------------------------------
.go                         245.67 KB     35.20%
.json                       156.34 KB     22.40%
.md                          89.12 KB     12.77%
.txt                         67.89 KB      9.73%
no extension                 45.23 KB      6.48%
.yaml                        34.56 KB      4.95%
.sh                          23.45 KB      3.36%
.gitignore                   12.34 KB      1.77%
--------------------------------------------------
Total                       698.12 KB    100.00%
Total extensions found: 8
```

## Features

- **Recursive scanning**: Automatically processes all subdirectories
- **Error handling**: Continues scanning even if some files cannot be accessed
- **Human-readable output**: File sizes are formatted in appropriate units (B, KB, MB, GB, etc.)
- **Sorted results**: Extensions are sorted by total size (largest first)
- **Percentage calculation**: Shows what percentage of total size each extension represents
- **No extension handling**: Files without extensions are grouped as "no extension"

## Error Handling

The program handles various error cases:

- Non-existent directories will show an error message
- Files that cannot be accessed are logged to stderr but don't stop the scan
- If a path points to a file instead of a directory, an error is displayed
