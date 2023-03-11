package transaction

import "time"

type Transaction struct {
	ID            string     `db:"id" json:"id"`
	WalletID      string     `db:"wallet_id" json:"-"`
	Amount        int64      `db:"amount" json:"amount"`
	BalanceBefore int64      `db:"balance_before" json:"-"`
	BalanceAfter  int64      `db:"balance_after" json:"-"`
	Type          string     `db:"type" json:"type,omitempty"`
	Status        string     `json:"status,omitempty"`
	ReferenceID   string     `db:"reference_id" json:"reference_id"`
	DepositedBy   *string    `db:"deposited_by" json:"deposited_by,omitempty"`
	DepositedAt   *time.Time `db:"deposited_at" json:"deposited_at,omitempty"`
	WithdrawnBy   *string    `db:"withdrawn_by" json:"withdrawn_by,omitempty"`
	WithdrawnAt   *time.Time `db:"withdrawn_at" json:"withdrawn_at,omitempty"`
}
