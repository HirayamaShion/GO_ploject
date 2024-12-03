package main

import "fmt"

func main() {
	num := 0
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		message := fmt.Sprintf("%d,", i)
		if i == num {
			fmt.Print(i)
			break
		}
		fmt.Print(message)

	}

}
