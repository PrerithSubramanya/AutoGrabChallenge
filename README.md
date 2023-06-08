---

# Simple Banking System

This is a simple banking system implemented in Go. It includes basic operations such as adding a customer to the bank, depositing and withdrawing funds, checking account balances, and transferring money between customers.

## Implementation

The project is implemented using Go's built-in packages and does not require any external dependencies. It includes the following core entities:

- `Customer`: Each `Customer` has a `Name` and a `Balance`. Customers can deposit and withdraw funds, and check their current balance.

- `Bank`: The `Bank` keeps track of multiple customers. It allows customers to be added, and facilitates transfer of funds between customers. It also allows checking of the total balance across all customers.

All operations that alter the state of the bank or its customers are guarded by a mutex to ensure thread-safety.

## How to Run the Project

1. Make sure you have Go installed. You can download it from the [official website](https://golang.org/dl/) if needed.

2. Clone this repository into your local machine.

3. Navigate to the project directory.

4. To run the tests, execute `go test` in the terminal. The output will show the test results.

Please note that this project does not include a command-line interface, GUI, or database persistence. The entire functionality can be tested using the provided test suite.

## Requirements

The banking system must fulfill the following requirements:

- Customers should be able to join the bank by providing a name and an initial deposit.
- The bank should maintain a balance for multiple customers and allow them to deposit, withdraw and check their current balances.
- The bank should prevent customers from withdrawing more money than they have in their accounts.
- Customers should be able to transfer money to another customer.
- The bank should be able to calculate the total balance across all customers.
