package cli

import (
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var Version string

type Opts struct {
	Command string
	Name    string
	URL     string
}

func Parse(args []string) (Opts, error) {
	parser := kingpin.New("web", "Associate websites to directory paths and easily open them.")
	parser.Version(Version)

	add := parser.Command("add", "Associate a website to the current working directory")
	addName := add.Arg("name", "Name to associate with").Required().String()
	addURL := add.Arg("url", "URL of website").Required().String()

	remove := parser.Command("remove", "Remove an associated website by name").Alias("rm")
	removeName := remove.Arg("name", "Name to remove").Required().String()

	list := parser.Command("list", "List websites associated with current working directory").Alias("ls")

	cmd, err := parser.Parse(args[1:])

	if err != nil && (!strings.HasPrefix(err.Error(), "expected command") || len(args) != 2) {
		return Opts{}, err
	}

	switch cmd {
	case add.FullCommand():
		return Opts{Command: "add", Name: *addName, URL: *addURL}, nil
	case remove.FullCommand():
		return Opts{Command: "remove", Name: *removeName}, nil
	case list.FullCommand():
		return Opts{Command: "list"}, nil
	default:
		return Opts{Command: "open", Name: args[1]}, nil
	}
}
