package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Println("2つの文字列を入力してください。")

	_, err := fmt.Scan(&str1, &str2)
	if err != nil {
		fmt.Println("入力にエラーがありました", err)
	}

	if str1 == str2 {
		fmt.Println("同じ文字列です。")
	} else {
		fmt.Println("異なる文字列です。")
	}
}
