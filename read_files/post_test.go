package read_files

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("error open file", func(t *testing.T) {
		_, err := NewPostsFromFS(&StubFailingFS{})

		if err == nil {
			t.Fatal("expected error but don't get one")
		}
	})
	t.Run("create post", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello, world!

The body of posts start after the '---'`
		)
		fs := fstest.MapFS{
			"hello_world.md": {Data: []byte(firstBody)},
		}

		got, err := NewPostsFromFS(fs)
		want := []Post{
			{
				Title:       "Post 1",
				Description: "Description 1",
				Body: `Hello, world!

The body of posts start after the '---'`,
				Tags: []string{"tdd", "go"},
			},
		}
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(got), len(fs))
		}
		assertPost(t, got, want)
	})
}

func assertPost(t testing.TB, got, want []Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
