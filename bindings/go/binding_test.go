package tree_sitter_arm64_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_arm64 "example.com//bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_arm64.Language())
	if language == nil {
		t.Errorf("Error loading Arm64 grammar")
	}
}
