package util

import (
	"bytes"
	"compress/gzip"
	"io"
	// "os"
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

// func MustCompressFile(file *os.File, data []byte) {
//   gzipWriter, err := gzip.NewWriterLevel(file, gzip.BestCompression)
//   if err != nil {
//     panic(err)
//   }
//
//   gzipWriter.Write(data)
//   gzipWriter.Close()
// }
//
// func MustDecompressFile(path string) []byte {
//   file := MustOpenFile(path)
//   defer file.Close()
//
//   reader, err := gzip.NewReader(file)
//   if err != nil {
//     panic(err)
//   }
//   defer reader.Close()
//
//   data, err := io.ReadAll(reader)
//   if err != nil {
//     panic(err)
//   }
//
//   return data
// }
