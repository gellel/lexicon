package lex_test

import (
	"fmt"
	"testing"

	"github.com/gellel/lex"
)

var (
    b    lex.Byte        // map[interface{}]byte
    c64  lex.Complex64  // map[interface{}]complex64
    c128 lex.Complex128 // map[interface{}]complex128
    f32  lex.Float32    // map[interface{}]float32
    f64  lex.Float64    // map[interface{}]float64
    i    lex.Int        // map[interface{}]interface{}
    i8   lex.Int8       // map[interface{}]int8
    i16  lex.Inter16      // map[interface{}]int16
    i32  lex.Int32      // map[interface{}]int32
    i64  lex.Int64      // map[interface{}]int64
    r    lex.Rune        // map[interface{}]rune
    s    *lex.Lex         // map[interface{}]interface{}
    u    lex.UInter       // map[interface{}]uint
    u8   lex.UInt8      // map[interface{}]uint8
    u16  lex.UInt16     // map[interface{}]uint16
    u32  lex.UInt32     // map[interface{}]uint32
    u64  lex.UInt64     // map[interface{}]uint64
    v    lex.Interface   // map[interface{}]interface{}
)

func Test(t *testing.T) {
	var (
		l = (&lex.Lex{})
	)
	fmt.Println(l)
}
