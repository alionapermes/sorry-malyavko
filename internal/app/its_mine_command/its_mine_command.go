package its_mine_command

import (
	"fmt"

	"github.com/spf13/cobra"
)

type itsMineCommand struct {
  *cobra.Command
}

func New() *cobra.Command {
  cmd := itsMineCommand{
    Command: &cobra.Command{
      Use: "its_mine",
      Run: run,
    },
  }
  cmd.setup()

  return cmd.Command
}

func (self *itsMineCommand) setup() {
  // TODO
}

func run(cmd *cobra.Command, args []string) {
  // TODO
  fmt.Printf("Command %s: not implemented yet\n", cmd.Use)
}
