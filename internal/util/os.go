package util

import (
	"io"
	"os"
)

func MustCreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func MustOpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func MustReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
  defer file.Close()

  data, err := io.ReadAll(file)
  if err != nil {
    panic(err)
  }

	return data
}
