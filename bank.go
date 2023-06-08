package AutoGrade

import (
	"errors"
	"sync"
)

// Customer struct represents a customer with a name and balance.
type Customer struct {
	Name    string  // Name of the customer
	Balance float64 // Balance of the customer
}

// Bank struct represents a bank with a map of customers.
type Bank struct {
	Customers map[string]*Customer // Map of customers in the bank
	Lock      *sync.RWMutex        // Read-Write lock for thread-safe operations
}

// NewCustomer function creates a new customer with an initial deposit.
func NewCustomer(name string, deposit float64) (*Customer, error) {
	// Ensure the initial deposit is not negative
	if deposit < 0 {
		return nil, errors.New("initial deposit cannot be negative")
	}
	return &Customer{Name: name, Balance: deposit}, nil
}

// NewBank function creates a new Bank.
func NewBank() *Bank {
	return &Bank{Customers: make(map[string]*Customer), Lock: new(sync.RWMutex)}
}

// AddCustomer adds a customer to the bank.
func (b *Bank) AddCustomer(c *Customer) {
	// Locking for thread safety
	b.Lock.Lock()
	defer b.Lock.Unlock()
	b.Customers[c.Name] = c
}

// Deposit adds an amount to the customer's balance.
func (c *Customer) Deposit(amount float64) error {
	// Ensure the deposit amount is not negative
	if amount < 0 {
		return errors.New("deposit amount cannot be negative")
	}
	c.Balance += amount
	return nil
}

// Withdraw subtracts an amount from the customer's balance.
func (c *Customer) Withdraw(amount float64) error {
	// Ensure the withdrawal amount is not negative
	if amount < 0 {
		return errors.New("withdrawal amount cannot be negative")
	}
	// Ensure the customer has sufficient funds
	if c.Balance < amount {
		return errors.New("insufficient funds")
	}
	c.Balance -= amount
	return nil
}

// CheckBalance returns the balance of the customer.
func (c *Customer) CheckBalance() float64 {
	return c.Balance
}

// TotalBankBalance calculates the total balance of all customers in the bank.
func (b *Bank) TotalBankBalance() float64 {
	// Locking for thread safety
	b.Lock.RLock()
	defer b.Lock.RUnlock()

	total := 0.0
	for _, c := range b.Customers {
		total += c.Balance
	}
	return total
}

// TransferFunds transfers an amount from one customer to another.
func (b *Bank) TransferFunds(from, to string, amount float64) error {
	// Locking for thread safety
	b.Lock.Lock()
	defer b.Lock.Unlock()

	// Ensure the transfer amount is not negative
	if amount < 0 {
		return errors.New("transfer amount cannot be negative")
	}

	sender, ok1 := b.Customers[from]
	receiver, ok2 := b.Customers[to]

	// Ensure both the sender and receiver exist
	if !ok1 || !ok2 {
		return errors.New("either sender or receiver does not exist")
	}

	// Withdraw from the sender
	if err := sender.Withdraw(amount); err != nil {
		return err
	}

	// Deposit to the receiver
	err := receiver.Deposit(amount)
	if err != nil {
		return err
	}

	return nil
}
