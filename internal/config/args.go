package config

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

type Args struct {
	Title   string
	Message string
	Labels  cli.StringSlice
	Owner   string
	Repo    string
	GHToken string
}

func ArgFlags(args *Args) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "title",
			Usage:       "set the title of the issue",
			EnvVars:     []string{"title", "t"},
			Destination: &args.Title,
		},
		&cli.StringFlag{
			Name:        "message",
			Usage:       "set the message for the issue",
			EnvVars:     []string{"message", "m"},
			Destination: &args.Message,
		},
		&cli.StringSliceFlag{
			Name:        "label",
			Usage:       "set the label of the issue",
			EnvVars:     []string{"label", "l"},
			Destination: &args.Labels,
		},
		&cli.StringFlag{
			Name:        "owner",
			Usage:       "set the owner of the issue",
			EnvVars:     []string{"owner", "o"},
			Destination: &args.Owner,
		},
		&cli.StringFlag{
			Name:        "repo",
			Usage:       "sets the repo the issue belongs to",
			EnvVars:     []string{"owner", "o"},
			Destination: &args.Owner,
		},
		&cli.StringFlag{
			Name:        "git-token",
			Usage:       "set the github token",
			EnvVars:     []string{"gittoken", "ght"},
			Destination: &args.Title,
		},
	}
}

const YString = "Y"

func CollectMissingArgs(args *Args) {
	inputReader := bufio.NewReader(os.Stdin)

	if args.Title == "" {
		fmt.Print("Please enter a title for the issue: ")
		args.Title, _ = inputReader.ReadString('\n')
	}
	if len(args.Labels.Value()) == 0 {
		fmt.Print("Would you like to add a label? (y to do so): ")
		addLabel, _ := inputReader.ReadString('\n')
		for strings.TrimSpace(strings.ToUpper(addLabel)) == YString {
			fmt.Print("Please input a label: ")
			label, _ := inputReader.ReadString('\n')
			_ = args.Labels.Set(label)
			fmt.Print("Do you want to add another label? (y to do so): ")
			addLabel, _ = inputReader.ReadString('\n')
		}
	}
	if args.Message == "" {
		fmt.Println("Please enter the Issue text below:")
		args.Message, _ = inputReader.ReadString('\n')
	}
	if args.GHToken == "" {
		fmt.Print("Please enter your GitHub Token: ")
		args.GHToken, _ = inputReader.ReadString('\n')
	}
	gh := getGHInfo()
	if args.Repo == "" {
		fmt.Printf("Detected %s as the repository. Would you like to change this? (y to do so):", gh.repo)
		changeInfo, _ := inputReader.ReadString('\n')
		if strings.TrimSpace(strings.ToUpper(changeInfo)) == YString {
			args.Repo, _ = inputReader.ReadString('\n')
		} else {
			args.Repo = gh.repo
		}
	}
	if args.Owner == "" {
		fmt.Printf("Detected %s as the organization/owner. Would you like to change this? (y to do so):", gh.owner)
		changeInfo, _ := inputReader.ReadString('\n')
		if strings.TrimSpace(strings.ToUpper(changeInfo)) == YString {
			args.Owner, _ = inputReader.ReadString('\n')
		} else {
			args.Owner = gh.owner
		}
	}
}

type ghInfo struct {
	repo  string
	owner string
}

func getGHInfo() ghInfo {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	out, _ := cmd.Output()
	r := regexp.MustCompile(`.*\:(?P<owner>[\w\-]{1,63})\/(?P<repo>[\w\-]{1,63})`)
	return ghInfo{
		repo:  r.FindStringSubmatch(string(out))[2],
		owner: r.FindStringSubmatch(string(out))[1],
	}
}
