package lexicon_test

import (
	"fmt"

	"github.com/gellel/lexicon"
)

func ExampleNew() {
	fmt.Println(lexicon.New())
	// Output: &map[]
}

func ExampleNewLexicon() {
	fmt.Println(lexicon.NewLexicon(map[string]interface{}{"key": "value"}))
	// Output: &map[key:value]
}
