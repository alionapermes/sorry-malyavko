package util

import (
	"bytes"
	// "encoding/binary"
	"encoding/gob"
	// "io"
)

// func MustBinaryRead(reader io.Reader, data any) {
// 	if err := binary.Read(reader, binary.BigEndian, &data); err != nil {
// 		panic(err)
// 	}
// }
//
//
// func MustBinaryWrite(writer io.Writer, data any) {
// 	if err := binary.Write(writer, binary.BigEndian, data); err != nil {
// 		panic(err)
// 	}
// }

func MustBinaryEncode(data any) []byte {
  var buf bytes.Buffer

  encoder := gob.NewEncoder(&buf)
  if err := encoder.Encode(data); err != nil {
    panic(err)
  }

  return buf.Bytes()
}
func MustBinaryDecode(data []byte, target any) {
  var buf bytes.Buffer
  buf.Write(data)

  decoder := gob.NewDecoder(&buf)
  if err := decoder.Decode(target); err != nil {
    panic(err)
  }
}
