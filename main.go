package main

import (
	"fmt"
	"os"
	"path"

	"github.com/dbridges/web-cli/cli"
	"github.com/dbridges/web-cli/store"
	"github.com/dbridges/web-cli/util"
)

func main() {
	opts, err := cli.Parse(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s, try --help\n", err)
	}

	store, err := store.NewJSONStore(path.Join(util.ConfigDir(), "web-cli.json"))
	util.Must(err)
	runner := cli.NewRunner(store)

	switch opts.Command {
	case "add":
		runner.Add(opts.Name, opts.URL)
	case "remove":
		runner.Remove(opts.Name)
	case "list":
		runner.List()
	case "open":
		runner.Open(opts.Name)
	}
}
