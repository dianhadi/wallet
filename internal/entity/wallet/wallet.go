package wallet

import (
	"time"

	"github.com/dianhadi/wallet/pkg/utils"
	"github.com/google/uuid"
)

type Wallet struct {
	ID         string     `db:"id" json:"id"`
	UserID     string     `db:"user_id" json:"owned_by"`
	Token      string     `db:"token" json:"token,omitempty"`
	Balance    int64      `db:"balance" json:"balance"`
	Status     string     `db:"status" json:"status"`
	CreatedAt  time.Time  `db:"created_at" json:"-"`
	EnabledAt  *time.Time `db:"enabled_at" json:"enabled_at,omitempty"`
	DisabledAt *time.Time `db:"disabled_at" json:"disabled_at,omitempty"`
}

func New(userID string) Wallet {
	newUuid := uuid.New()
	wallet := Wallet{
		ID:     newUuid.String(),
		UserID: userID,
		Token:  utils.GenerateToken(),
		Status: "",
	}

	return wallet
}

func (w Wallet) IsEnabled() bool {
	if w.EnabledAt != nil {
		return true
	}
	return false
}

func (w Wallet) IsDisabled() bool {
	if w.DisabledAt != nil {
		return true
	}
	return false
}
