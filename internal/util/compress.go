package util

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Compress(data []byte) []byte {
  var buf bytes.Buffer

  gzipWriter := gzip.NewWriter(&buf)
  gzipWriter.Write(data)
  gzipWriter.Close()

  return buf.Bytes()
}

func MustDecompress(data []byte) []byte {
  var buf bytes.Buffer
  buf.Write(data)

  gzipReader, err := gzip.NewReader(&buf)
  if err != nil {
    panic(err)
  }
  defer gzipReader.Close()

  decompressed, err := io.ReadAll(gzipReader)
  if err != nil {
    panic(err)
  }

  return decompressed
}
