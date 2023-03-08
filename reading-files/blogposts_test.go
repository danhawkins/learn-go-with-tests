package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/danhawkins/blogposts"
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(`Title: Post 1
Description: Description 1`)},
		"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Description 2`)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "Description 1"})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
