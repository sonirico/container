package main

import (
	"github.com/sonirico/container/maps"
	"log"
)

type square struct {
	side int
}

func (s square) Area() int {
	return s.side * s.side
}

func testSlice() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}

	squares := maps.Slice(arr, func(x int) square {
		return square{side: x}
	})

	log.Println(squares[0].Area())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func testMap() {
	m := map[string]int{
		"apples":  1,
		"oranges": 2,
		"bananas": 3,
	}

	reversed := maps.Map(m, func(k1 string, v1 int) (int, string) {
		return v1, reverse(k1)
	})

	log.Println(reversed)
}

func main() {
	testSlice()
	testMap()
}
