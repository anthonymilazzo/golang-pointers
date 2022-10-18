package main

import "fmt"

func add_dollar(balance1 int) {
	balance1 = balance1 + 1
	fmt.Printf("balance1 %d\n", balance1)
}

func add_dollar_pointer(balance2 *int) {
	*balance2 = *balance2 + 1
	fmt.Printf("balance2 %d\n", *balance2)
}
