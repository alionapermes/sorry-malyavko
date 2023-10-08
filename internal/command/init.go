package command

import (
	"log"

	"github.com/markphelps/optional"

	"github.com/alionapermes/sorry-malyavko/internal/constant"
	"github.com/alionapermes/sorry-malyavko/internal/data"
	"github.com/alionapermes/sorry-malyavko/internal/model"
	"github.com/alionapermes/sorry-malyavko/internal/util"
)

type InitFlags struct {
	ID               optional.Uint16
	Password         optional.String
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
  log.Println("readStudent")
	id, err := self.Flags.ID.Get()
	if err != nil {
		id = util.MustScanUint16()
	}

	password, err := self.Flags.Password.Get()
	if err != nil {
		password = util.MustScanStudentPassword()
	}

	self.student = model.NewStudent(id, password)
	return self
}

func (self *initCommand) readSshConfig() *initCommand {
  log.Println("readSshConfig")
	var config model.SshConfig

	if self.Flags.DefaultSshConfig {
		config = model.SshConfig{
			Host: constant.DefaultHost,
			Port: constant.DefaultPort,
		}
	} else {
		config = model.SshConfig{
			Host: util.ScanHostOr(constant.DefaultHost),
			Port: util.ScanPortOr(constant.DefaultPort),
		}
	}

	self.sshConfig = config
	return self
}

func (self *initCommand) initCache() *initCommand {
  log.Println("initCache")
	sshClient := util.MustSshClient(self.student, self.sshConfig)
	defer sshClient.Close()
  log.Println("ssh client created")

	sftpClient := util.MustSftpClient(sshClient)
	defer sftpClient.Close()
  log.Println("sftp client created")

	userlistFile := util.MustSftpOpen(sftpClient, constant.UserlistPath)
	defer userlistFile.Close()
  log.Println("userlist fetched")

	self.cache = data.NewCacheFromUserlist(userlistFile)
  log.Println("cache OK")
	return self
}

func (self *initCommand) saveCache() *initCommand {
  log.Println("saveCache")
	self.cache.Save()
	return self
}

func (self *initCommand) createAppConfig() *initCommand {
  log.Println("createAppConfig")
	type appConfig struct {
		DefaultStudent *model.Student
		*model.SshConfig
	}

	appConfigFile := util.MustCreateFile(constant.CachePath)
	defer appConfigFile.Close()

	config := appConfig{
		DefaultStudent: &self.student,
		SshConfig:      &self.sshConfig,
	}

	util.MustTomlEncode(appConfigFile, config)
	return self
}
