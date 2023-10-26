package gomap_test

import (
	"testing"

	"github.com/lindsaygelle/gomap"
)

func BenchmarkAdd(b *testing.B) {
	newMap := &gomap.Map[int, int]{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap.Add(i, i)
	}
}

func BenchmarkDelete(b *testing.B) {
	newMap := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap.Add(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap.Delete(i % 1000)
	}
}

func BenchmarkFilter(b *testing.B) {
	newMap := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap.Add(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap.Filter(func(key int, value int) bool {
			return value%2 == 0
		})
	}
}

func BenchmarkGet(b *testing.B) {
	newMap := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap.Add(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap.Get(i % 1000)
	}
}

func BenchmarkIntersection(b *testing.B) {
	newMap1 := &gomap.Map[int, int]{}
	newMap2 := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap1.Add(i, i)
		newMap2.Add(i+500, i+500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap1.Intersection(newMap2)
	}
}

func BenchmarkLength(b *testing.B) {
	newMap := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap.Add(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap.Length()
	}
}

func BenchmarkMap(b *testing.B) {
	newMap := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap.Add(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap.Map(func(key int, value int) int {
			return value * 2
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	newMap1 := &gomap.Map[int, int]{}
	newMap2 := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap1.Add(i, i)
		newMap2.Add(i+1000, i+1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap1.Merge(newMap2)
	}
}

func BenchmarkValues(b *testing.B) {
	newMap := &gomap.Map[int, int]{}
	for i := 0; i < 1000; i++ {
		newMap.Add(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMap.Values()
	}
}
