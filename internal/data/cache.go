package data

import (
	"errors"
	"io"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
	"github.com/alionapermes/sorry-malyavko/internal/model"
	"github.com/alionapermes/sorry-malyavko/internal/util"
)

type studentsMap map[model.StudentID]model.StudentPassword

type Cache struct {
	students studentsMap
}

func NewCacheFromBinary(path string) *Cache {
	cache := Cache{
		students: make(studentsMap),
	}

  cache.loadFromBinary()
	return &cache
}

func NewCacheFromUserlist(reader io.Reader) *Cache {
	cache := Cache{
		students: make(studentsMap),
	}

	for student := range util.MustParseUserlist(reader) {
		cache.students[student.ID] = student.Password
	}

	return &cache
}

func (self *Cache) Save() {
	cacheFile := util.MustCreateFile(constant.CachePath)
	defer cacheFile.Close()

  data := util.MustBinaryEncode(self.students)
  compressed := util.Compress(data)

  cacheFile.Write(compressed)
}

func (self *Cache) GetByID(id model.StudentID) (model.Student, error) {
  if password, ok := self.students[id]; ok {
    return model.Student{ID: id, Password: password}, nil
  }

  return model.Student{}, errors.New("Not found")
}

func (self *Cache) GetAll() <-chan model.Student {
	studChan := make(chan model.Student)
	go func() {
    for id, password := range self.students {
      studChan <- model.Student{ID: id, Password: password}
    }
    close(studChan)
  }()
	return studChan
}

func (self *Cache) loadFromBinary() {
  data := util.MustReadFile(constant.CachePath)
  decompressed := util.MustDecompress(data)

  util.MustBinaryDecode(decompressed, &self.students)
}
