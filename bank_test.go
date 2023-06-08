package AutoGrade

import (
	"testing"
)

// TestBankOperations is a unit test for all the operations in the Bank.
func TestBankOperations(t *testing.T) {
	bank := NewBank()

	customer1, err := NewCustomer("Alice", 1000)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	customer2, err := NewCustomer("Bob", 500)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	bank.AddCustomer(customer1)
	bank.AddCustomer(customer2)

	err = customer1.Deposit(500)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if customer1.CheckBalance() != 1500 {
		t.Errorf("Expected balance to be 1500, got %f", customer1.CheckBalance())
	}

	err = customer1.Withdraw(2000)
	if err == nil {
		t.Error("Expected error when withdrawing more than balance")
	}

	err = customer1.Withdraw(1000)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if customer1.CheckBalance() != 500 {
		t.Errorf("Expected balance to be 500, got %f", customer1.CheckBalance())
	}

	if bank.TotalBankBalance() != 1000 {
		t.Errorf("Expected total bank balance to be 1000, got %f", bank.TotalBankBalance())
	}

	err = bank.TransferFunds("Alice", "Bob", 100)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if customer1.CheckBalance() != 400 || customer2.CheckBalance() != 600 {
		t.Errorf("Expected Alice's balance to be 400 and Bob's balance to be 600, got %f and %f respectively", customer1.CheckBalance(), customer2.CheckBalance())
	}

	// Test negative deposit
	err = customer1.Deposit(-500)
	if err == nil {
		t.Error("Expected error when depositing negative amount")
	}

	// Test negative withdrawal
	err = customer1.Withdraw(-500)
	if err == nil {
		t.Error("Expected error when withdrawing negative amount")
	}

	// Test negative transfer
	err = bank.TransferFunds("Alice", "Bob", -100)
	if err == nil {
		t.Error("Expected error when transferring negative amount")
	}
}
