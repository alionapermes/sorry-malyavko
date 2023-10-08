package init_command

import (
	"github.com/markphelps/optional"
	"github.com/spf13/cobra"

	"github.com/alionapermes/sorry-malyavko/internal/command"
	"github.com/alionapermes/sorry-malyavko/internal/util"
)

type initCommand struct {
  *cobra.Command
}

func New() *cobra.Command {
	cmd := initCommand{
    Command: &cobra.Command{
      Use: "init",
      Run: run,
    },
	}
	cmd.setup()

	return cmd.Command
}

func (self *initCommand) setup() {
	self.PersistentFlags().StringP("id", "n", "", "")
	self.PersistentFlags().StringP("password", "p", "", "")
	self.PersistentFlags().
		BoolP("default", "d", false, "Use default ssh config")
}

func run(cmd *cobra.Command, args []string) {
  flags := command.InitFlags{DefaultSshConfig: false}

  if flag := cmd.Flag("id"); flag != nil {
    id := util.MustParseUint16(flag.Value.String())
    flags.ID = optional.NewUint16(id)
  }

  if flag := cmd.Flag("password"); flag != nil {
    flags.Password = optional.NewString(flag.Value.String())
  }

  if flag := cmd.Flag("default"); flag != nil {
    flags.DefaultSshConfig = true
  }

  command.Init(flags)
}
