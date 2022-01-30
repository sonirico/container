package main

import (
	"log"

	"github.com/sonirico/container/filter"
)

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}

	arr = filter.SliceFilter(arr, func(x int) bool {
		return x%3 == 0
	})

	log.Println(arr)
}
