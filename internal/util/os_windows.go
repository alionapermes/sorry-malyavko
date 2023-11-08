package util

import (
	"io"
	"os"
	"syscall"
)

func MustCreateFile(path string, makeHidden bool) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

  filenameW, err := syscall.UTF16PtrFromString(path)
  if err != nil {
    panic(err)
  }
  syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_HIDDEN)

	return file
}

func MustOpenFile(path string, makeHidden bool) *os.File {
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
