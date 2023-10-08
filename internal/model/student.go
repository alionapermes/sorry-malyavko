package model

import "github.com/alionapermes/sorry-malyavko/internal/constant"

type StudentID uint16
type StudentPassword [constant.StudentPasswordLength]byte

type Student struct {
	ID       StudentID
	Password StudentPassword
}

func NewStudent(id uint16, password string) Student {
	return Student{
		ID:       StudentID(id),
		Password: StudentPassword([]byte(password)),
	}
}
