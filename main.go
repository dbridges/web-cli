package main

import (
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/dbridges/web-cli/app"
	"github.com/dbridges/web-cli/store"
	"github.com/dbridges/web-cli/util"
)

var Version string

func main() {
	parser := kingpin.New("web", "Associate websites to directory paths and easily open them.")
	parser.Version(Version)

	add := parser.Command("add", "Associate a website to the current working directory")
	addName := add.Arg("name", "Name to associate with").Required().String()
	addURL := add.Arg("url", "URL of website").Required().String()

	remove := parser.Command("remove", "Remove an associated website by name").Alias("rm")
	removeName := remove.Arg("name", "Name to remove").Required().String()

	list := parser.Command("list", "List websites associated with current working directory").Alias("ls")

	cmd, parseErr := parser.Parse(os.Args[1:])

	if parseErr != nil && (!strings.HasPrefix(parseErr.Error(), "expected command") || len(os.Args) != 2) {
		kingpin.Fatalf("%s, try --help", parseErr)
	}

	store, err := store.NewJSONStore()
	util.Must(err)
	app := app.New(store)

	if parseErr != nil {
		app.Open(os.Args[1])
		return
	}

	switch cmd {
	case add.FullCommand():
		app.Add(*addName, *addURL)
	case remove.FullCommand():
		app.Remove(*removeName)
	case list.FullCommand():
		app.List()
	}
}
