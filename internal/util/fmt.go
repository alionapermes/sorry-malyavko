package util

import (
	"fmt"
	"strconv"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
)

func MustScanUint16() (value uint16) {
  fmt.Scanln(&value)
  return
}

func MustParseUint16(str string) uint16 {
  value, err := strconv.Atoi(str)
  if err != nil {
    panic(err)
  }
  return uint16(value)
}

func MustScanStudentPassword() (password string) {
  fmt.Scanln(&password)
  if len(password) != constant.StudentPasswordLength {
    // TODO
  }
  return
}

func ScanHostOr(defaultHost string) (host string) {
  fmt.Scanln(&host)
  if len(host) == 0 {
    return defaultHost
  }

  // TODO: validation

  return
}

func ScanPortOr(defaultPort uint16) (port uint16) {
  fmt.Scanln(&port)
  if port == 0 {
    return defaultPort
  }

  // TODO: validation

  return
}
