package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestDB(t *testing.T, URL string) (*Database, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.URL = URL
	s := New(config)
	if err := s.Connect(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Disconnect()
	}

}
