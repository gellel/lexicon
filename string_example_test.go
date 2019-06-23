package lexicon_test

import (
	"fmt"

	"github.com/gellel/lexicon"
)

func ExampleNewString() {
	fmt.Println(lexicon.NewString(map[string]string{"a": "a"}))
	// Output: &map[a:a]
}
