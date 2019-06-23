package lexicon_test

import (
	"fmt"

	"github.com/gellel/lexicon"
)

func ExampleNewInt() {
	fmt.Println(lexicon.NewInt(map[string]int{"a": 1}))
	// Output: &map[a:1]
}
