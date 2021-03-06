package util

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/kyoh86/xdg"
)

func Must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func GitURL() string {
	cmd := exec.Command("git", "remote", "-v")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	// Get origin fetch line
	re := regexp.MustCompile(`origin(.+)\(fetch\)`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) != 2 {
		return ""
	}
	match := strings.TrimSpace(matches[1])
	if strings.Contains(match, "@") {
		match = strings.Split(match, "@")[1]
		match = strings.ReplaceAll(match, ":", "/")
		match = strings.TrimSuffix(match, ".git")
		return "http://" + match
	} else if strings.HasPrefix(match, "http") {
		return match
	}
	return ""
}

func ConfigDir() string {
	fpath, err := xdg.FindConfigFile("web-cli")
	if err != nil {
		return path.Join(xdg.ConfigHome(), "web-cli")
	}
	return fpath
}
