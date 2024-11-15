package bank_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"bank-acc-management/bank"
)

var _ = Describe("Bank", func() {
	var account bank.Account
	BeforeEach(func() {
		account = bank.Account{}
	})

	Describe("Deposit", func() {
		It("should deposit money into the account", func() {
			err := account.Deposit(100.50)
			Expect(err).To(BeNil())
			Expect(account.GetBalance()).To(Equal(100.50))
		})
		It("should return an error for negative deposit amount", func() {
			err := account.Deposit(-100)
			Expect(err).To(MatchError("Deposit amount must be greater than 0"))
		})
	})

	Describe("Withdraw", func() {
		It("should withdraw money from the account", func() {
			account.Deposit(100.50)
			err := account.Withdraw(50)
			Expect(err).To(BeNil())
			Expect(account.GetBalance()).To(Equal(50.50))
		})
		It("should return an error for negative withdraw amount", func() {
			err := account.Withdraw(-100)
			Expect(err).To(MatchError("Withdraw amount must be greater than 0"))
		})
		It("should return an error for insufficient funds", func() {
			account.Deposit(50)
			err := account.Withdraw(100)
			Expect(err).To(MatchError("Insufficient funds"))
		})
	})

	Describe("GetBalance", func() {
		It("should return the account balance", func() {
			account.Deposit(100.50)
			Expect(account.GetBalance()).To(Equal(100.50))
		})
	})
})
