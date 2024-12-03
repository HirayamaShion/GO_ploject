package main

import (
	"fmt"
	"strings"
)

func main() {
	var targetString, searchChar string

	// 入力の促し
	fmt.Println("検索対象の文字列を入力してください：")
	_, err := fmt.Scan(&targetString)
	if err != nil {
		fmt.Println("入力にエラーがありました:", err)
		return
	}

	fmt.Println("検索する1文字を入力してください：")
	_, err = fmt.Scan(&searchChar)
	if err != nil {
		fmt.Println("入力にエラーがありました:", err)
		return
	}

	// 検索文字が1文字か確認
	if len(searchChar) != 1 {
		fmt.Println("検索文字は1文字のみ指定してください。")
		return
	}

	// 検索処理
	index := strings.Index(targetString, searchChar)
	fmt.Println(index)
}
