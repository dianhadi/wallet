package transaction

import (
	entityTransaction "github.com/dianhadi/wallet/internal/entity/transaction"
	entityWallet "github.com/dianhadi/wallet/internal/entity/wallet"
)

type rTransaction interface {
	Deposit(transaction entityTransaction.Transaction) (entityTransaction.Transaction, error)
	Withdraw(transaction entityTransaction.Transaction) (entityTransaction.Transaction, error)
	GetAllByWalletID(walletID string) ([]entityTransaction.Transaction, error)
	GetIDByReferenceID(referenceID string) (string, error)
}

type rWallet interface {
	GetByID(walletID string) (entityWallet.Wallet, error)
}

type Handler struct {
	repoTransaction rTransaction
	repoWallet      rWallet
}

type Transaction struct {
	Deposit    *entityTransaction.Transaction `json:"deposit,omitempty"`
	Withdrawal *entityTransaction.Transaction `json:"withdrawal,omitempty"`
}

type Transactions struct {
	Transactions []entityTransaction.Transaction `json:"transactions,omitempty"`
}

func New(repoTransaction rTransaction, repoWallet rWallet) (*Handler, error) {
	return &Handler{
		repoTransaction: repoTransaction,
		repoWallet:      repoWallet,
	}, nil
}
