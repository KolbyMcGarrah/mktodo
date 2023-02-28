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
			Destination: &args.Title,
		},
		&cli.StringFlag{
			Name:        "message",
			Usage:       "set the message for the issue",
			Destination: &args.Message,
		},
		&cli.StringSliceFlag{
			Name:        "label",
			Usage:       "set the label of the issue",
			Destination: &args.Labels,
		},
		&cli.StringFlag{
			Name:        "owner",
			Usage:       "set the owner of the issue",
			Destination: &args.Owner,
		},
		&cli.StringFlag{
			Name:        "repo",
			Usage:       "sets the repo the issue belongs to",
			Destination: &args.Owner,
		},
		&cli.StringFlag{
			Name:        "git-token",
			Usage:       "set the github token",
			Destination: &args.GHToken,
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

	var gh ghInfo
	if args.Repo == "" || args.Owner == "" {
		gh = getGHInfo()
	}
	if args.Repo == "" {
		fmt.Printf("Detected %s as the repository. Would you like to change this? (y to do so): ", gh.repo)
		changeInfo, _ := inputReader.ReadString('\n')
		if strings.TrimSpace(strings.ToUpper(changeInfo)) == YString {
			fmt.Println("Please enter the repository name: ")
			readValue, _ := inputReader.ReadString('\n')
			args.Repo = readValue[:len(readValue)-1]
		} else {
			args.Repo = gh.repo
		}
	}
	if args.Owner == "" {
		fmt.Printf("Detected %s as the organization/owner. Would you like to change this? (y to do so): ", gh.owner)
		changeInfo, _ := inputReader.ReadString('\n')
		if strings.TrimSpace(strings.ToUpper(changeInfo)) == YString {
			fmt.Println("Please enter the organization/owner name: ")
			readValue, _ := inputReader.ReadString('\n')
			args.Owner = readValue[:len(readValue)-1]
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
	match := r.FindStringSubmatch(string(out))
	if len(match) < 2 {
		match = regexp.MustCompile(`.*github.com\/(?P<owner>[\w\-]{1,63})\/(?P<repo>[\w\-]{1,63})`).FindStringSubmatch(string(out))
	}
	return ghInfo{
		repo:  match[2],
		owner: match[1],
	}
}
