package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
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
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{readMetaLine(titleSeparator),
		readMetaLine(descriptionSeparator),
		strings.Split(readMetaLine(tagsSeparator), ", "),
		newBody(scanner),
	}, nil
}

func newBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		_, _ = fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
