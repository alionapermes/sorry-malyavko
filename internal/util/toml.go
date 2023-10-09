package util

import (
	"os"

	"github.com/BurntSushi/toml"
)

func MustTomlEncode(file *os.File, data any) {
	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}

func MustTomlDecodeFile(path string, data any) {
  if _, err := toml.DecodeFile(path, data); err != nil {
    panic(err)
  }
}
