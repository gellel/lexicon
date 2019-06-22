package lexicon_test

import (
	"testing"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

func TestExampleNewSlice(t *testing.T) {
	lexicon.NewSlice().Add("key", slice.New())
	// Output: &map[key:&[]]
}
