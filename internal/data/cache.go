package data

import (
	// "bytes"
	// "encoding/binary"
	// "encoding/gob"
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

	studentsChan := make(chan model.Student)
	util.MustParseUserlist(studentsChan, reader)

	for student := range studentsChan {
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

func (self *Cache) loadFromBinary() {
  data := util.MustReadFile(constant.CachePath)
  decompressed := util.MustDecompress(data)

  util.MustBinaryDecode(decompressed, &self.students)
}
