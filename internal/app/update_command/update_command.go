package update_command

import (
	"fmt"

	"github.com/spf13/cobra"
)

type updateCommand struct {
  *cobra.Command
}

func New() *cobra.Command {
  cmd := updateCommand{
    Command: &cobra.Command{
      Use: "update",
      Run: run,
    },
  }
  cmd.setup()

  return cmd.Command
}

func (self *updateCommand) setup() {
  // TODO
}

func run(cmd *cobra.Command, args []string) {
  // TODO
  fmt.Printf("Command %s: not implemented yet\n", cmd.Use)
}
