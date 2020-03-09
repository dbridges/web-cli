package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/dbridges/web-cli/store"
	"github.com/dbridges/web-cli/util"
)

type Runner interface {
	Open(string)
	Add(string, string)
	Remove(string)
	List()
}

type runner struct {
	store store.Store
	cwd   string
}

func NewRunner(store store.Store) Runner {
	cwd, err := os.Getwd()
	util.Must(err)
	homeDir, err := os.UserHomeDir()
	util.Must(err)
	if strings.HasPrefix(cwd, homeDir) {
		cwd = strings.Replace(cwd, homeDir, "~", 1)
	}
	runner := runner{store: store, cwd: cwd}
	return &runner
}

func (runner *runner) Open(name string) {
	entries := runner.store.Search(runner.cwd, name)
	switch len(entries) {
	case 0:
		if name == "git" {
			gitURL := util.GitURL()
			if len(gitURL) == 0 {
				fmt.Println("Unable to parse git remote origin")
			} else {
				runner.openURL(gitURL)
			}
		} else {
			fmt.Printf("No entry for '%s'\n", name)
		}
	case 1:
		runner.openURL(entries[0].URL)
	default:
		fmt.Println("Duplicate entries found")
	}
}

func (runner *runner) List() {
	entries := runner.store.Search(runner.cwd, "")
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

func (runner *runner) Add(name, url string) {
	util.Must(runner.store.Add(runner.cwd, name, url))
	fmt.Printf("Added %s (%s)\n", name, url)
}

func (runner *runner) Remove(name string) {
	util.Must(runner.store.Remove(runner.cwd, name))
	fmt.Printf("Removed %s\n", name)
}

func (runner *runner) openURL(url string) {
	var ex string
	switch runtime.GOOS {
	case "darwin":
		ex = "open"
	case "windows":
		ex = "start"
	default: // Unix
		ex = "xdg-open"
	}
	if ex != "" {
		cmd := exec.Command(ex, url)
		util.Must(cmd.Run())
	} else {
		fmt.Printf("Unsupported operating system %s\n", runtime.GOOS)
	}
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
