package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

type web struct {
	app string
}

func (web *web) run() {
	app := kingpin.New("web", "Associate websites to directory paths and easily open them.")
	kingpin.Version("0.0.1")

	add := app.Command("add", "Associate a website to the current directory")
	addName := add.Arg("name", "Name to associate with").Required().String()
	addURL := add.Arg("url", "URL of website").Required().String()

	remove := app.Command("remove", "Remove associated website").Alias("rm")
	removeName := remove.Arg("name", "Name to remove").Required().String()

	list := app.Command("list", "List websites associate with current directory").Alias("ls")

	cmd, err := app.Parse(os.Args[1:])

	if err != nil {
		if strings.HasPrefix(err.Error(), "expected command") && len(os.Args) == 2 {
			web.open(os.Args[1])
			return
		}
		kingpin.Fatalf("%s, try --help", err)
	}

	switch cmd {
	case add.FullCommand():
		web.add(*addName, *addURL)
	case remove.FullCommand():
		web.remove(*removeName)
	case list.FullCommand():
		web.list()
	}
}

func (web *web) open(name string) {
	fmt.Printf("Opening %s\n", name)
}

func (web *web) list() {
	fmt.Println("Listing")
}

func (web *web) add(name, URL string) {
	fmt.Printf("Adding %s: %s\n", name, URL)
}

func (web *web) remove(name string) {
	fmt.Printf("Removing %s\n", name)
}

func main() {
	web := &web{}
	web.run()
}
