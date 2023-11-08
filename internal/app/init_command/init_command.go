package init_command

import (
	"github.com/spf13/cobra"

	"github.com/alionapermes/sorry-malyavko/internal/command"
	"github.com/alionapermes/sorry-malyavko/internal/util"
)

type initCommand struct {
	*cobra.Command
	flagProvider util.FlagProvider
}

func New() *cobra.Command {
	cobraCmd := &cobra.Command{Use: "init"}

	cmd := initCommand{
		Command:      cobraCmd,
		flagProvider: util.NewFlagProvider(cobraCmd),
	}

	cmd.Run = cmd.run
	cmd.setup()

	return cmd.Command
}

func (self *initCommand) setup() {
	self.flagProvider.
		Register("id", "n", uint16(0), "").
		Register("password", "p", "", "").
		Register("default", "d", false, "Use default ssh config")
}

func (self *initCommand) run(cmd *cobra.Command, args []string) {
	flags := command.InitFlags{
    ID: nil,
    Password: nil,
    DefaultSshConfig: false,
  }

  if id := self.flagProvider.Value("id"); id != nil {
    flags.ID = id.(*uint16)
  }

  if password := self.flagProvider.Value("password"); password != nil {
	  flags.Password = password.(*string)
  }

  flags.DefaultSshConfig = *self.flagProvider.ValueOrDefault("default").(*bool)

	command.Init(flags)
}
