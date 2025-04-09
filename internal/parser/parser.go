package parser

import (
	"context"
	"fmt"
	sitter "github.com/smacker/go-tree-sitter"
	"os"
	"path/filepath"
)

func GetComments(filePath string) ([]Comment, error) {
	fileExt := filepath.Ext(filePath)

	language, err := getLanguage(fileExt)
	if err != nil {
		return nil, err
	}

	var sourceCode, readFileErr = os.ReadFile(filePath)
	if readFileErr != nil {
		return nil, fmt.Errorf("\nfailed to read file %s: %v\n", filePath, err)
	}

	parser := sitter.NewParser()
	parser.SetLanguage(language)

	tree, _ := parser.ParseCtx(context.Background(), nil, sourceCode)
	root := tree.RootNode()

	var comments []Comment

	var walk func(*sitter.Node)
	walk = func(node *sitter.Node) {
		if node.Type() == "comment" {
			text := cleanComment(node, sourceCode)
			comments = append(comments, Comment{Text: text, Line: node.StartPoint().Row + 1})
		}

		for i := 0; i < int(node.ChildCount()); i++ {
			walk(node.Child(i))
		}
	}

	walk(root)

	return comments, nil
}

type Comment struct {
	Text string
	Line uint32
}
