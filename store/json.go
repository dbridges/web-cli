package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type jsonStore struct {
	path    string
	entries []Entry
}

func NewJSONStore(path string) (*jsonStore, error) {
	store := &jsonStore{path: path}
	if err := store.load(); err != nil {
		return nil, err
	}
	return store, nil
}

func (store *jsonStore) Add(path, name, url string) error {
	existing := store.Search(path, name)
	if len(existing) > 0 {
		return ErrExists
	}
	store.entries = append(store.entries, Entry{Path: path, Name: name, URL: url})
	if err := store.save(); err != nil {
		return err
	}
	return nil
}

func (store *jsonStore) Remove(path, name string) error {
	newEntries := make([]Entry, 0)
	for _, entry := range store.entries {
		if entry.Path != path || entry.Name != name {
			newEntries = append(newEntries, entry)
		}
	}
	store.entries = newEntries
	if err := store.save(); err != nil {
		return err
	}
	return nil
}

func (store *jsonStore) Search(path, name string) []Entry {
	entries := make([]Entry, 0)
	for _, e := range store.entries {
		pathMatch := e.Path == path && len(path) > 0 || len(path) == 0
		nameMatch := e.Name == name && len(name) > 0 || len(name) == 0
		if pathMatch && nameMatch {
			entries = append(entries, e)
		}
	}
	return entries
}

func (store *jsonStore) load() error {
	_, err := os.Stat(store.path)
	if err != nil {
		if os.IsNotExist(err) {
			store.entries = make([]Entry, 0)
			return nil
		}
		return fmt.Errorf("Unknown error: %v", err)
	}
	// Load from existing file
	data, err := ioutil.ReadFile(store.path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &store.entries)
	if err != nil {
		return err
	}
	return nil
}

func (store *jsonStore) save() error {
	err := os.MkdirAll(path.Dir(store.path), 0755)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(store.entries, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(store.path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
