package wallet

import (
	entity "github.com/dianhadi/wallet/internal/entity/wallet"
)

type rWallet interface {
	GetByID(walletID string) (entity.Wallet, error)
	GetIDByUserID(userID string) (string, error)
	Insert(wallet entity.Wallet) error
	Enable(wallet entity.Wallet) (entity.Wallet, error)
	Disable(wallet entity.Wallet) (entity.Wallet, error)
}

type Handler struct {
	repoWallet rWallet
}

type Token struct {
	Token string `json:"token"`
}

func New(repoWallet rWallet) (*Handler, error) {
	return &Handler{
		repoWallet: repoWallet,
	}, nil
}
