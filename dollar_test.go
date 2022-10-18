package main

import (
	"fmt"
	"testing"
)

func TestAddDollar(t *testing.T) {
	bank_balance1 := 20
	fmt.Printf("bank_balance1 %d\n", bank_balance1)
	add_dollar(bank_balance1)
	fmt.Printf("bank_balance1 %d\n", bank_balance1)
	if bank_balance1 != 20 {
		t.Fatalf("Bank balance is not 20! Value: %d", bank_balance1)
	}
}

func TestAddDollarPointer(t *testing.T) {
	bank_balance2 := 20
	fmt.Printf("bank_balance2 %d\n", bank_balance2)
	add_dollar_pointer(&bank_balance2)
	fmt.Printf("bank_balance2 %d\n", bank_balance2)
	if bank_balance2 != 21 {
		t.Fatalf("Bank balance is not 21! Value: %d", bank_balance2)
	}
}
