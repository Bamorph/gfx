package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func findFilesByExtensions(root string, exts []string) {
	extPatterns := make([]*regexp.Regexp, len(exts))
	for i, ext := range exts {
		extPatterns[i] = regexp.MustCompile("\\." + ext + "$")
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() {
			for _, pattern := range extPatterns {
				if pattern.MatchString(info.Name()) {
					fmt.Println(root + "/" + path)
					break
				}
			}
		}

		return nil
	})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gfx <file_extension1> [<file_extension2> ...]")
		os.Exit(1)
	}

	root := "." // You can change this to the desired directory
	exts := os.Args[1:]

	findFilesByExtensions(root, exts)
}
