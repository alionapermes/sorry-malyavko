package util

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
	"github.com/alionapermes/sorry-malyavko/internal/model"
)

func MustParseUserlist(studentsChan chan model.Student, reader io.Reader) {
  data, err := io.ReadAll(reader)
  if err != nil {
    panic(err)
  }

  pattern := fmt.Sprintf(`stud0?(\d+)\s+(\w{%d})`, constant.StudentPasswordLength)
  expr := regexp.MustCompile(pattern)

  linesChan := make(chan string)
  readLines(linesChan, data)

  for line := range linesChan {
    res := expr.FindStringSubmatch(line)

    num, err := strconv.Atoi(res[1])
    if err != nil {
      panic(err)
    }

    id := model.StudentID(num)
    password := model.StudentPassword([]byte(res[2]))

    studentsChan <- model.Student{ID: id, Password: password}
  }
}

func readLines(linesChan chan<- string, data []byte) {
  nl := byte('\n')
  from := 0

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
}
