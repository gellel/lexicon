package lexicon_test

import (
	"fmt"
	"testing"

	"github.com/gellel/lexicon"
)

func Test(t *testing.T) {
	var (
		l = &lexicon.Lex{}
	)
	l.Add("hello", "world")
	fmt.Println(l.Get("hello"))
}
