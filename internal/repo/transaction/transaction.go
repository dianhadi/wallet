package transaction

import (
	"context"
	"database/sql"
	"errors"
	"time"

	entity "github.com/dianhadi/wallet/internal/entity/transaction"
)

type database interface {
	Connect() (*sql.DB, error)
}

type Transaction struct {
	db database
}

var (
	ErrBalanceInsufficient = errors.New("balance is insufficient")
)

func New(db database) (Transaction, error) {
	return Transaction{
		db: db,
	}, nil
}

// GetAll retrieves all rows from the transactions table
func (t Transaction) GetAllByWalletID(walletID string) ([]entity.Transaction, error) {
	db, err := t.db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, reference_id, amount, type, deposited_by, deposited_at, withdrawn_by, withdrawn_at FROM transactions WHERE wallet_id=$1", walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction
	for rows.Next() {
		transaction := entity.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.ReferenceID, &transaction.Amount, &transaction.Type, &transaction.DepositedBy, &transaction.DepositedAt, &transaction.WithdrawnBy, &transaction.WithdrawnAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// GetByUserID retrieves wallet data from the wallet table by user ID
func (t Transaction) GetIDByReferenceID(referenceID string) (string, error) {
	db, err := t.db.Connect()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string
	err = db.QueryRow("SELECT id FROM transactions WHERE reference_id = $1", referenceID).Scan(&id)

	return id, err
}

func (t Transaction) Deposit(transaction entity.Transaction) (entity.Transaction, error) {
	db, err := t.db.Connect()
	if err != nil {
		return entity.Transaction{}, err
	}
	defer db.Close()

	opts := &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return entity.Transaction{}, err
	}
	defer tx.Rollback()

	var balance int64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id=$1 FOR UPDATE", transaction.WalletID).Scan(&balance)
	if err != nil {
		return entity.Transaction{}, err
	}

	// update balance and create transaction record
	transaction.BalanceBefore = balance
	transaction.BalanceAfter = balance + transaction.Amount
	now := time.Now()
	transaction.DepositedAt = &now

	_, err = tx.Exec("UPDATE wallets SET balance=$1 WHERE id=$2", transaction.BalanceAfter, transaction.WalletID)
	if err != nil {
		return entity.Transaction{}, err
	}

	_, err = tx.Exec("INSERT INTO transactions (id, wallet_id, type, reference_id, amount, balance_before, balance_after, deposited_by, deposited_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		transaction.ID,
		transaction.WalletID,
		transaction.Type,
		transaction.ReferenceID,
		transaction.Amount,
		transaction.BalanceBefore,
		transaction.BalanceAfter,
		transaction.DepositedBy,
		transaction.DepositedAt,
	)

	err = tx.Commit()
	if err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}

func (t Transaction) Withdraw(transaction entity.Transaction) (entity.Transaction, error) {
	db, err := t.db.Connect()
	if err != nil {
		return entity.Transaction{}, err
	}
	defer db.Close()

	opts := &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return entity.Transaction{}, err
	}
	defer tx.Rollback()

	var balance int64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id=$1 FOR UPDATE", transaction.WalletID).Scan(&balance)
	if err != nil {
		return entity.Transaction{}, err
	}

	// update balance and create transaction record
	transaction.BalanceBefore = balance
	transaction.BalanceAfter = balance - transaction.Amount
	if transaction.BalanceAfter < 0 {
		return entity.Transaction{}, ErrBalanceInsufficient
	}
	now := time.Now()
	transaction.WithdrawnAt = &now

	_, err = tx.Exec("UPDATE wallets SET balance=$1 WHERE id=$2", transaction.BalanceAfter, transaction.WalletID)
	if err != nil {
		return entity.Transaction{}, err
	}

	_, err = tx.Exec("INSERT INTO transactions (id, wallet_id, type, reference_id, amount, balance_before, balance_after, withdrawn_by, withdrawn_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		transaction.ID,
		transaction.WalletID,
		transaction.Type,
		transaction.ReferenceID,
		transaction.Amount,
		transaction.BalanceBefore,
		transaction.BalanceAfter,
		transaction.WithdrawnBy,
		transaction.WithdrawnAt,
	)

	err = tx.Commit()
	if err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}
