package main

import (
	"fmt"
	"os"

	"github.com/kolbymcgarrah/mktodo/cmd/mktodo"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Flags = nil
	app.Commands = []*cli.Command{
		mktodo.NewMkTodoCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("app exiting with errors: ", err)
		os.Exit(1)
	}
}
