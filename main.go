package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func getExtension(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return "no extension"
	}
	return strings.ToLower(ext)
}

func walkDirectory(dir string) (map[string]int64, error) {
	extensionSizes := make(map[string]int64)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Log error but continue walking
			fmt.Fprintf(os.Stderr, "Error accessing %s: %v\n", path, err)
			return nil
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Get file info
		info, err := d.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting info for %s: %v\n", path, err)
			return nil
		}

		// Get extension and add size
		ext := getExtension(d.Name())
		extensionSizes[ext] += info.Size()

		return nil
	})

	return extensionSizes, err
}

type ExtensionSize struct {
	Extension string
	Size      int64
}

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting absolute path: %v\n", err)
		absDir = dir
	}

	info, err := os.Stat(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: %s is not a directory\n", dir)
		os.Exit(1)
	}

	fmt.Printf("Scanning directory: %s\n", absDir)
	fmt.Println(strings.Repeat("-", 70))

	extensionSizes, err := walkDirectory(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}

	if len(extensionSizes) == 0 {
		fmt.Println("No files found in the specified directory.")
		return
	}

	var results []ExtensionSize
	var totalSize int64
	for ext, size := range extensionSizes {
		results = append(results, ExtensionSize{Extension: ext, Size: size})
		totalSize += size
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Size > results[j].Size
	})

	fmt.Printf("%-40s %15s %10s\n", "Extension", "Size", "Percentage")
	fmt.Println(strings.Repeat("-", 70))

	for _, item := range results {
		percentage := float64(item.Size) / float64(totalSize) * 100
		fmt.Printf("%-40s %15s %9.2f%%\n", item.Extension, formatBytes(item.Size), percentage)
	}

	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("%-40s %15s %9.2f%%\n", "Total", formatBytes(totalSize), 100.0)
	fmt.Printf("Total extensions found: %d\n", len(results))
}
