package command

import (
	"github.com/markphelps/optional"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
	"github.com/alionapermes/sorry-malyavko/internal/data"
	"github.com/alionapermes/sorry-malyavko/internal/model"
	"github.com/alionapermes/sorry-malyavko/internal/util"
)

type ImInArgs struct {
	ID optional.Uint16
}

type imInCommand struct {
	Args ImInArgs

	student   model.Student
	sshConfig model.SshConfig
	cache     *data.Cache
}

func ImIn(args ImInArgs) {
	(&imInCommand{Args: args}).
		loadAppConfig().
		loadCache().
		sshConnect()
}

func (self *imInCommand) loadAppConfig() *imInCommand {
	var appConfig struct {
		DefaultStudent model.Student
		model.SshConfig
	}

	util.MustTomlDecodeFile(constant.ConfigPath, &appConfig)

	self.student = appConfig.DefaultStudent
	self.sshConfig = appConfig.SshConfig

	return self
}

func (self *imInCommand) loadCache() *imInCommand {
	self.cache = data.NewCacheFromBinary(constant.CachePath)
	return self
}

func (self *imInCommand) sshConnect() *imInCommand {
  var stud model.Student
  id, err := self.Args.ID.Get()

  if err != nil {
    stud = self.student
  } else {
    studID := model.StudentID(id)
    stud, err = self.cache.GetByID(studID)

    if err != nil {
      panic(err)
    }
  }

  util.MustSshShell(stud, self.sshConfig)
	return self
}
