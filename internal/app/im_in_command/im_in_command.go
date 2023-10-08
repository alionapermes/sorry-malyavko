package im_in_command

import (
	"github.com/markphelps/optional"
	"github.com/spf13/cobra"

	"github.com/alionapermes/sorry-malyavko/internal/command"
	"github.com/alionapermes/sorry-malyavko/internal/util"
)

type imInCommand struct {
	*cobra.Command
}

func New() *cobra.Command {
	cmd := imInCommand{
		Command: &cobra.Command{
			Use:  "im-in",
			Args: cobra.MaximumNArgs(1),
			Run:  run,
		},
	}
	cmd.setup()

	return cmd.Command
}

func (self *imInCommand) setup() {
	// TODO
}

func run(cmd *cobra.Command, cmdArgs []string) {
  var args command.ImInArgs

	if len(cmdArgs) == 1 {
    id := util.MustParseUint16(cmdArgs[1])
    args.ID = optional.NewUint16(id)
  }

  command.ImIn(args)
}
