//go:build aix || dragonfly || freebsd || (js && wasm) || wasip1 || linux || netbsd || openbsd || solaris

package util

import (
	"io"
	"os"
	"strings"
)

func MustCreateFile(path string, makeHidden bool) *os.File {
	if !strings.HasPrefix(path, ".") {
		path = "." + path
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

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
