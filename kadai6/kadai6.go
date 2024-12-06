package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{64, 3, 45, 291, 178}
	fmt.Println(minInt(numbers))

}

func minInt(slice []int) int {
	sort.Sort(sort.IntSlice(slice))
	return slice[0]
}
