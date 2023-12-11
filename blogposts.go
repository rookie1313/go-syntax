package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
	"testing/fstest"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

type Post struct {
	Title       string
	Description string
}

func NewPostsFromFS(fileSystem fstest.MapFS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err //TODO: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}

	return posts, nil
}
func getPost(fileSystem fs.FS, name string) (Post, error) {
	postFile, err := fileSystem.Open(name)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)
	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readLine(titleSeparator)
	description := readLine(descriptionSeparator)

	return Post{title, description}, nil
}
