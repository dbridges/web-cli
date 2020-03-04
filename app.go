package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Store interface {
	Add(path, name, url string) error
	Remove(path, name string) error
	Search(path, name string) []Entry
}

type App struct {
	store Store
	cwd   string
}

func (app *App) open(name string) {
	entries := app.store.Search(app.cwd, name)
	switch len(entries) {
	case 0:
		if name == "git" {
			gitURL := GitURL()
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

func (app *App) list() {
	entries := app.store.Search(app.cwd, "")
	git := Entry{Name: "git", URL: GitURL()}
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

func (app *App) add(name, url string) {
	must(app.store.Add(app.cwd, name, url))
	fmt.Printf("Added %s (%s)\n", name, url)
}

func (app *App) remove(name string) {
	must(app.store.Remove(app.cwd, name))
	fmt.Printf("Removed %s\n", name)
}

func (app *App) openURL(url string) {
	cmd := exec.Command("open", url)
	must(cmd.Run())
}

// Appends entry to entries if an entry with the same name does not already exist
func entriesWith(entries []Entry, entry Entry) []Entry {
	for _, e := range entries {
		if e.Name == entry.Name {
			return entries
		}
	}
	return append(entries, entry)
}

func NewApp(store Store) *App {
	cwd, err := os.Getwd()
	must(err)
	homeDir, err := os.UserHomeDir()
	must(err)
	if strings.HasPrefix(cwd, homeDir) {
		cwd = strings.Replace(cwd, homeDir, "~", 1)
	}
	app := App{store: store, cwd: cwd}
	return &app
}
