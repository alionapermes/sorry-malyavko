package command

import (
	"fmt"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
	"github.com/alionapermes/sorry-malyavko/internal/data"
	"github.com/alionapermes/sorry-malyavko/internal/model"
	"github.com/alionapermes/sorry-malyavko/internal/util"
)

type InitFlags struct {
	ID               *uint16
	Password         *string
	DefaultSshConfig bool
}

type initCommand struct {
	Flags InitFlags

	student   model.Student
	sshConfig model.SshConfig
	cache     *data.Cache
}

func Init(flags InitFlags) {
	(&initCommand{Flags: flags}).
		readStudent().
		readSshConfig().
		initCache().
		saveCache().
		createAppConfig()
}

func (self *initCommand) readStudent() *initCommand {
  var id uint16
  var password string

  if n := self.Flags.ID; n != nil {
    id = *n
  } else {
    fmt.Print("ID: ")
		id = util.MustScanUint16()
	}

  if p := self.Flags.Password; p != nil {
    password = *p
  } else {
    fmt.Print("Password: ")
		password = util.MustScanStudentPassword()
	}

	self.student = model.NewStudent(id, password)
	return self
}

func (self *initCommand) readSshConfig() *initCommand {
  var host string
  var port uint16

	if self.Flags.DefaultSshConfig {
    host = constant.DefaultHost
    port = constant.DefaultPort
	} else {
    fmt.Printf("Host (%s by default): ", constant.DefaultHost)
		host = util.ScanHostOr(constant.DefaultHost)
    fmt.Printf("Port (%d by default): ", constant.DefaultPort)
		port = util.ScanPortOr(constant.DefaultPort)
	}

	self.sshConfig = model.SshConfig{
    Host: host,
    Port: port,
  }
	return self
}

func (self *initCommand) initCache() *initCommand {
	sshClient := util.MustSshClient(self.student, self.sshConfig)
	defer sshClient.Close()

	sftpClient := util.MustSftpClient(sshClient)
	defer sftpClient.Close()

	userlistFile := util.MustSftpOpen(sftpClient, constant.UserlistPath)
	defer userlistFile.Close()

	self.cache = data.NewCacheFromUserlist(userlistFile)
	return self
}

func (self *initCommand) saveCache() *initCommand {
	self.cache.Save()
	return self
}

func (self *initCommand) createAppConfig() *initCommand {
	type appConfig struct {
		DefaultStudent *model.Student
		*model.SshConfig
	}

	appConfigFile := util.MustCreateFile(constant.ConfigPath)
	defer appConfigFile.Close()

	config := appConfig{
		DefaultStudent: &self.student,
		SshConfig:      &self.sshConfig,
	}

	util.MustTomlEncode(appConfigFile, config)
	return self
}
