package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	titleLine := readLine()[7:]

	descriptionLine := readLine()[13:]

	return Post{Title: titleLine, Description: descriptionLine}, nil
}
