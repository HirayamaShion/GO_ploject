package main

import (
	"fmt"
)

func main() {

	slice := []string{"あいう", "あいうえ", "あいうえお", "かきくけこ", "かきく"}
	for _, str := range slice {
		//fmt.Println(str, "len:", utf8.RuneCountInString(str))
		fmt.Println(max(str))
	}

}
