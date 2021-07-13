package base

import (
	"context"
	"testing"
)

func TestRepo(t *testing.T) {
	r := NewRepo("https://github.com/togettoyou/go-one-server.git", "template")
	if err := r.Clone(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.CopyTo(context.Background(), "/tmp/test_repo", "go-one-server", nil); err != nil {
		t.Fatal(err)
	}
}
