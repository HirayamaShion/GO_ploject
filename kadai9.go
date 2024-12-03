package main

import (
	"fmt"
	"slices"
)

func main() {

	str := []string{"あいう", "あいうえ", "あいうえお", "かきくけこ", "かきく"}
	//for _, str := range slice {
	//fmt.Println(str, "len:", utf8.RuneCountInString(str))
	fmt.Println(slices.Max(str))
	//}

}
