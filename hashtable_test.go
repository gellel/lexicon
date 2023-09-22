package hashtable_test

import (
	"hashtable"
	"testing"
)

func TestHashtable(t *testing.T) {
	table := hashtable.Hashtable[string, string]{
		"hello": "world"}
	t.Log(table)
}
