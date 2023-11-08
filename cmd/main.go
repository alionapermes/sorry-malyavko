package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/alionapermes/sorry-malyavko/internal/app/im_in_command"
	"github.com/alionapermes/sorry-malyavko/internal/app/init_command"
	"github.com/alionapermes/sorry-malyavko/internal/app/its_mine_command"
	"github.com/alionapermes/sorry-malyavko/internal/app/update_command"
)

const (
  exitOK = 0
  exitError = 1
)

func main() {
  os.Exit(run())
}

func run() int {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("Runtime error:", err)
    }
  }()

	cli := &cobra.Command{
		Use: "sorry-malyavko",
  }

  cli.AddCommand(init_command.New())
  cli.AddCommand(im_in_command.New())
  cli.AddCommand(update_command.New())
  cli.AddCommand(its_mine_command.New())

	if err := cli.Execute(); err != nil {
    fmt.Println("Error:", err.Error())
    return exitError
	}

  return exitOK
}
