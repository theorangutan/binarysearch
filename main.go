package main

import (
	"flag"
	"fmt"
)

func main() {
	search := flag.Int64("s", 4, "int to search")
	flag.Parse()

	arr := []int64{0, 1, 2, 3, 4, 5}
	length := int64(len(arr))
	location := BinarySearch(arr, 0, length-1, *search)
	fmt.Println(location)
}
