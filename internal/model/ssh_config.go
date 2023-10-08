package model

import "fmt"

type SshConfig struct {
	Host string
  Port uint16
}

func (self *SshConfig) HostPort() string {
  return fmt.Sprintf("%s:%d", self.Host, self.Port)
}
