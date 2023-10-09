package util

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
	"github.com/alionapermes/sorry-malyavko/internal/model"
)

func MustParseUserlist(reader io.Reader) <-chan model.Student {
  data, err := io.ReadAll(reader)
  if err != nil {
    panic(err)
  }

  pattern := fmt.Sprintf(`stud0?(\d+)\s+(\w{%d})`, constant.StudentPasswordLength)
  expr := regexp.MustCompile(pattern)

  studentsChan := make(chan model.Student)

  go func() {
    for line := range readLines(data) {
      res := expr.FindStringSubmatch(line)

      num, err := strconv.Atoi(res[1])
      if err != nil {
        panic(err)
      }

      id := model.StudentID(num)
      password := model.StudentPassword([]byte(res[2]))

      studentsChan <- model.Student{ID: id, Password: password}
    }

    close(studentsChan)
  }()

  return studentsChan
}

func readLines(data []byte) <-chan string {
  nl := byte('\n')
  from := 0

  linesChan := make(chan string)

  go func() {
    for idx, char := range data {
      if char == nl {
        linesChan <- string(data[from:idx])
        from = idx + 1
      }
    }

    idxLast := len(data)
    if from < idxLast {
      linesChan <- string(data[from:idxLast])
    }

    close(linesChan)
  }()

  return linesChan
}
