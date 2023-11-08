package command

import (
	// "context"
	// "fmt"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
	"github.com/alionapermes/sorry-malyavko/internal/data"
	"github.com/alionapermes/sorry-malyavko/internal/model"
	"github.com/alionapermes/sorry-malyavko/internal/util"
	"github.com/markphelps/optional"
)

type ItsMineFlags struct {
	Me     optional.Bool
	All    optional.Bool
	ID     optional.Uint16
	Target optional.String
}

type itsMineCommand struct {
  Flags ItsMineFlags

	student   model.Student
	sshConfig model.SshConfig
	cache     *data.Cache
}

func ItsMine(flags ItsMineFlags) {
  (&itsMineCommand{Flags: flags}).
    loadAppConfig().
    loadCache().
    downloadTarget()
}

func (self *itsMineCommand) loadAppConfig() *itsMineCommand {
	var appConfig struct {
		DefaultStudent model.Student
		model.SshConfig
	}

	util.MustTomlDecodeFile(constant.ConfigPath, &appConfig)

	self.student = appConfig.DefaultStudent
	self.sshConfig = appConfig.SshConfig

	return self
}

func (self *itsMineCommand) loadCache() *itsMineCommand {
	self.cache = data.NewCacheFromBinary(constant.CachePath)
	return self
}

func (self *itsMineCommand) downloadTarget() *itsMineCommand {
  return nil
  // sshClient := util.MustSshClient(self.student, self.sshConfig)
  // defer sshClient.Close()
  //
  // session, err := sshClient.NewSession()
  // if err != nil {
  //   panic(err)
  // }
  // defer session.Close()
  //
  // var id model.StudentID
  // var target string
  //
  // if _, err := self.Flags.Me.Get(); err != nil {
  //   id = self.student.ID
  // }
  //
  // var id uint16
  // var target string
  //
  // if pID := self.Flags.ID; pID != nil {
  //   id = *pID
  // } else {
  //   id = self
  // }
  //
  // if pTarget := self.Flags.Target; pTarget != nil {
  //   target = *pTarget
  // } else {
  //   target = buildHomePath(id)
  // }
  // zipCmd := ""
  // if err := session.Run(zipCmd); err != nil {
  //   panic(err)
  // }
  //
  // scpClient := util.MustScpClient(sshClient)
  // defer scpClient.Close()
  //
  // file := util.MustCreateFile("dump.zip")
  // defer file.Close()
  //
  // util.MustScpCopyRemote(sshClient, file, target)
  //
  // return self
}

func (self *itsMineCommand) buildHomePath(id uint) {
}
