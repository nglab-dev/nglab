package utils

import "os"

func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

func MkdirIfNotExist(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
