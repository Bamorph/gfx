package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func findFilesByExtension(root string, ext string, pattern *regexp.Regexp) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && pattern.MatchString(info.Name()) {
			fmt.Println(path)
		}

		return nil
	})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_extension>")
		os.Exit(1)
	}

	root := "." // You can change this to the desired directory
	ext := os.Args[1]
	extPattern := regexp.MustCompile("\\." + ext + "$")

	findFilesByExtension(root, ext, extPattern)
}
