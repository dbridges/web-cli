package store

import (
	"os"
	"path"
	"testing"
	"time"
)

func newJSONTestStore(t *testing.T) *jsonStore {
	path := path.Join(os.TempDir(), string(time.Now().UnixNano())+".json")

	t.Cleanup(func() {
		os.Remove(path)
	})

	store, err := NewJSONStore(path)
	if err != nil {
		t.Fatalf("Error creating JSON Store")
	}

	return store
}

func TestJSONStore(t *testing.T) {
	store := newJSONTestStore(t)

	err := store.Add("test/path", "test_name", "http://test")
	if err != nil {
		t.Fatalf("Expected to Add to store")
	}
	err = store.Add("test/path", "test_name", "http://test")
	if err != ErrExists {
		t.Fatalf("Expected to receive ErrExists")
	}
	err = store.Add("test/path2", "test_name", "http://test2")
	if err != nil {
		t.Fatalf("Expected to Add to store")
	}

	matches := store.Search("test/path", "")
	if len(matches) != 1 {
		t.Fatalf("Expected 1 match")
	}

	matches = store.Search("", "test_name")
	if len(matches) != 2 {
		t.Fatalf("Expected 2 matches")
	}

	matches = store.Search("bad/path", "test_name")
	if len(matches) != 0 {
		t.Fatalf("Expected 0 matches")
	}

	err = store.Remove("test/path", "test_name")
	if err != nil {
		t.Fatalf("Expected to remove path")
	}
	matches = store.Search("", "test_name")
	if len(matches) != 1 {
		t.Fatalf("Expected 1 matches")
	}
}
