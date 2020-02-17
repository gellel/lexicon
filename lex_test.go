package lex_test

import (
	"fmt"
	"testing"

	"github.com/gellel/lex"
)

func Test(t *testing.T) {
	var (
		l = (&lex.Lex{})
	)
	fmt.Println(l)
}
