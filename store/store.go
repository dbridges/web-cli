package store

import "fmt"

var ErrExists = fmt.Errorf("Already exists")

type Entry struct {
	Path string
	Name string
	URL  string
}

type Store interface {
	Add(path, name, url string) error
	Remove(path, name string) error
	Search(path, name string) []Entry
}
