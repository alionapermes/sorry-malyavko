package util

import (
	"fmt"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	"github.com/alionapermes/sorry-malyavko/internal/model"
)

func MustSshClient(user model.Student, connCfg model.SshConfig) *ssh.Client {
	clientCfg := &ssh.ClientConfig{
		User: makeUsername(user.ID),
		Auth: []ssh.AuthMethod{
			ssh.Password(string(user.Password[:])),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", connCfg.HostPort(), clientCfg)
	if err != nil {
		panic(err)
	}
	return client
}

func MustSftpClient(sshClient *ssh.Client) *sftp.Client {
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		panic(err)
	}
  return sftpClient
}

func MustSftpOpen(client *sftp.Client, path string) *sftp.File {
  file, err := client.Open(path)
  if err != nil {
    panic(err)
  }
  return file
}

func makeUsername(id model.StudentID) string {
	prefix := "stud"
	if id < 10 {
		prefix += "0"
	}
	return fmt.Sprintf("%s%d", prefix, id)
}
