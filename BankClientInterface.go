package main


type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)
	
	// Withdrawal withdraws given amount from clients account. 
	// return error if clients balance less the withdrawal amount 
	Withdrawal(amount int) error
	
	// Balance returns clients balance
	Balance() int
}