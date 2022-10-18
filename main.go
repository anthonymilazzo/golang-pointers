package main

import "fmt"

// Output:
// Without pointers:
// bank_balance1 20
// bank_balance1 21
// bank_balance1 20
// ---
// With pointers:
// bank_balance2 20
// bank_balance1 21
// bank_balance2 21

func add_dollar(balance1 int) {
	balance1 = balance1 + 1
	fmt.Printf("bank_balance1 %d\n", balance1)
}

func add_dollar_pointer(balance2 *int) {
	*balance2 = *balance2 + 1
	fmt.Printf("bank_balance1 %d\n", *balance2)
}

func main() {
	fmt.Println("Without pointers:")

	bank_balance1 := 20
	fmt.Printf("bank_balance1 %d\n", bank_balance1)
	add_dollar(bank_balance1)
	fmt.Printf("bank_balance1 %d\n", bank_balance1)

	fmt.Println("---")

	fmt.Println("With pointers:")

	bank_balance2 := 20
	fmt.Printf("bank_balance2 %d\n", bank_balance2)
	add_dollar_pointer(&bank_balance2)
	fmt.Printf("bank_balance2 %d\n", bank_balance2)
}
