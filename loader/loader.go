package loader

import (
	"os"
	"path/filepath"
	"strings"
)

func LoadFiles(path string, fileTypes []string) (map[string][]string, error) {
	fileMap := make(map[string][]string)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		for _, fileType := range fileTypes {
			if filepath.Ext(path) == fileType {
				fileMap[path] = append(fileMap[path], strings.Split(path, "/")[len(strings.Split(path, "/"))-1])
			}
		}
		return nil
	})

	return fileMap, err
}
