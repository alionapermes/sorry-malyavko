package its_mine_command

import (
	"fmt"

	"github.com/markphelps/optional"
	"github.com/spf13/cobra"

	"github.com/alionapermes/sorry-malyavko/internal/command"
)

type itsMineCommand struct {
	*cobra.Command

	flagAll    bool
	flagMe     bool
	flagID     uint16
	flagTarget string
}

func New() *cobra.Command {
	cmd := itsMineCommand{Command: &cobra.Command{}}
	cmd.Use = "its-mine"
  // cmd.Run = cmd.run
	cmd.Run = func(cmd *cobra.Command, args []string) {
    fmt.Printf("Command %s: not implemented yet\n", cmd.Use)
  }
	cmd.setup()
	return cmd.Command
}

func (self *itsMineCommand) setup() {
	flagSet := self.PersistentFlags()

	flagSet.BoolVarP(&self.flagAll, "all", "a", false, "dowload from all users")
	flagSet.BoolVarP(&self.flagMe, "me", "m", true, "download from your user")
	flagSet.Uint16VarP(&self.flagID, "id", "n", 0, "download stud<id> user")
	flagSet.StringVarP(
		&self.flagTarget, "target", "t", "", "download a target from specified user")
}

func (self *itsMineCommand) run(cmd *cobra.Command, args []string) {
	var flags command.ItsMineFlags

	if flag := self.Flag("all"); flag.Changed {
		flags.All = optional.NewBool(self.flagAll)
	}

	if flag := self.Flag("me"); flag.Changed {
		flags.Me = optional.NewBool(self.flagMe)
	}

	if flag := self.Flag("id"); flag.Changed {
		flags.ID = optional.NewUint16(self.flagID)
	}

	if flag := self.Flag("all"); flag.Changed {
		flags.Target = optional.NewString(self.flagTarget)
	}

	command.ItsMine(flags)
}
