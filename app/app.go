package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dbridges/web-cli/store"
	"github.com/dbridges/web-cli/util"
)

type App struct {
	store store.Store
	cwd   string
}

func New(store store.Store) *App {
	cwd, err := os.Getwd()
	util.Must(err)
	homeDir, err := os.UserHomeDir()
	util.Must(err)
	if strings.HasPrefix(cwd, homeDir) {
		cwd = strings.Replace(cwd, homeDir, "~", 1)
	}
	app := App{store: store, cwd: cwd}
	return &app
}

func (app *App) Open(name string) {
	entries := app.store.Search(app.cwd, name)
	switch len(entries) {
	case 0:
		if name == "git" {
			gitURL := util.GitURL()
			if len(gitURL) == 0 {
				fmt.Println("Unable to parse git remote origin")
			} else {
				app.openURL(gitURL)
			}
		} else {
			fmt.Printf("No entry for '%s'\n", name)
		}
	case 1:
		app.openURL(entries[0].URL)
	default:
		fmt.Println("Duplicate entries found")
	}
}

func (app *App) List() {
	entries := app.store.Search(app.cwd, "")
	git := store.Entry{Name: "git", URL: util.GitURL()}
	if len(git.URL) > 0 {
		entries = entriesWith(entries, git)
	}
	if len(entries) == 0 {
		fmt.Println("No entries found")
	}
	for _, e := range entries {
		fmt.Printf("%s\t%s\n", e.Name, e.URL)
	}
}

func (app *App) Add(name, url string) {
	util.Must(app.store.Add(app.cwd, name, url))
	fmt.Printf("Added %s (%s)\n", name, url)
}

func (app *App) Remove(name string) {
	util.Must(app.store.Remove(app.cwd, name))
	fmt.Printf("Removed %s\n", name)
}

func (app *App) openURL(url string) {
	cmd := exec.Command("open", url)
	util.Must(cmd.Run())
}

// Appends entry to entries if an entry with the same name does not already exist
func entriesWith(entries []store.Entry, entry store.Entry) []store.Entry {
	for _, e := range entries {
		if e.Name == entry.Name {
			return entries
		}
	}
	return append(entries, entry)
}
