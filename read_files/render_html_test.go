package read_files

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "Hello World",
			Body:        "## markdown document",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = Post{
			Title:       "Hello World",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
