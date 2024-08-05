package utils

import (
	"io/fs"
	"path/filepath"
)

func WalkDir(root string) ([]string, error) {
	var a []string
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		a = append(a, s)
		return nil
	})
	return a, err
}
