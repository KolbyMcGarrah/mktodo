package mktodo

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/kolbymcgarrah/mktodo/internal/config"
	"github.com/kolbymcgarrah/mktodo/internal/github"
	"github.com/urfave/cli/v2"
)

type todo struct {
	config *config.Config
}

func NewTodo() *todo {
	return &todo{
		config: config.NewConfig(),
	}
}

func NewMkTodoCmd() *cli.Command {
	todo := NewTodo()
	cmd := cli.Command{
		Name:   "mk",
		Action: todo.runCmd,
	}
	cmd.Flags = append(cmd.Flags, config.ArgFlags(todo.config.Args)...)
	return &cmd
}

func (t *todo) runCmd(*cli.Context) error {
	ctx := context.Background()
	// placeholder
	_ = ctx
	config.CollectMissingArgs(t.config.Args)

	fmt.Println("Reaching out to Github to create your Issue")

	hc := github.NewHTTPCreator(http.DefaultClient)
	url, err := hc.CreateIssue(ctx, github.Request{
		Owner:   t.config.Args.Owner,
		Repo:    t.config.Args.Repo,
		Body:    t.config.Args.Message,
		Title:   t.config.Args.Title,
		GHToken: t.config.Args.GHToken,
		Labels:  t.config.Args.Labels.Value(),
	})
	if err != nil {
		return err
	}
	// make request

	fmt.Printf("Success!\nYour Issue has been created and can be viewed here: %#v\n", strings.Replace(url, "%", "", -1))
	return nil
}
