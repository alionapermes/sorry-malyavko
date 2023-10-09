package util

import (
	"fmt"
	"os"

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

func MustSshShell(student model.Student, sshConfig model.SshConfig) {
  client := MustSshClient(student, sshConfig)

  session, err := client.NewSession()
  if err != nil {
    panic(err)
  }
  defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		panic(err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	if err := session.Shell(); err != nil {
		panic(err)
	}

	if err := session.Wait(); err != nil {
		panic(err)
	}
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
