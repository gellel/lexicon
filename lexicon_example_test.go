package lexicon_test

import (
	"testing"

	"github.com/gellel/lexicon"
)

func TestExampleNew(t *testing.T) {
	lexicon.New()
	// Output: &map[]
}

func TestExampleNewLexicon(t *testing.T) {
	lexicon.NewLexicon(map[string]interface{}{"key": "value"})
	// Output: &map[key:value]
}
