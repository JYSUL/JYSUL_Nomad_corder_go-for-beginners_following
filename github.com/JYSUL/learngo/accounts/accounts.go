package accounts

import "errors"

// BankAcount struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw your balance")

// NewAcount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 실제 메모리를 참조해서 변경해야한다.
// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney // or error.Error()
	}
	a.balance -= amount
	return nil
}
