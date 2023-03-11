package wallet

import (
	"database/sql"
	"fmt"
	"time"

	entity "github.com/dianhadi/wallet/internal/entity/wallet"
)

type database interface {
	Connect() (*sql.DB, error)
}

type Wallet struct {
	db database
}

func New(db database) (Wallet, error) {
	return Wallet{
		db: db,
	}, nil
}

// GetByID retrieves wallet data from the wallet table by wallet ID
func (w Wallet) GetByID(walletID string) (entity.Wallet, error) {
	db, err := w.db.Connect()
	if err != nil {
		return entity.Wallet{}, err
	}
	defer db.Close()

	var wallet entity.Wallet
	err = db.QueryRow("SELECT id, user_id, balance, status, enabled_at, disabled_at FROM wallets WHERE id = $1", walletID).Scan(&wallet.ID, &wallet.UserID, &wallet.Balance, &wallet.Status, &wallet.EnabledAt, &wallet.DisabledAt)

	return wallet, err
}

// GetByUserID retrieves wallet data from the wallet table by user ID
func (w Wallet) GetIDByUserID(userID string) (string, error) {
	db, err := w.db.Connect()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string
	err = db.QueryRow("SELECT id FROM wallets WHERE user_id = $1", userID).Scan(&id)

	return id, err
}

// GetByToken retrieves wallet data from the wallet table by token
func (w Wallet) GetByToken(token string) (string, string, error) {
	db, err := w.db.Connect()
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	var id, userID string
	err = db.QueryRow("SELECT id, user_id FROM wallets WHERE token = $1", token).Scan(&id, &userID)

	return id, userID, err
}

// Insert row with initial user data
func (w Wallet) Insert(wallet entity.Wallet) error {
	db, err := w.db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO wallets (id,user_id,token,balance,status,created_at) values($1,$2,$3,$4,$5,$6)",
		wallet.ID,
		wallet.UserID,
		wallet.Token,
		wallet.Balance,
		wallet.Status,
		time.Now())
	return err
}

// Update row enable data
func (w Wallet) Enable(wallet entity.Wallet) (entity.Wallet, error) {
	db, err := w.db.Connect()
	if err != nil {
		return entity.Wallet{}, err
	}
	defer db.Close()

	// prepare the SQL statement
	stmt, err := db.Prepare("UPDATE wallets SET status=$1, enabled_at=$2 WHERE id=$3")
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error preparing statement: %s", err)
	}
	defer stmt.Close()

	// execute the SQL statement with the updated data
	wallet.Status = "enabled"
	now := time.Now()
	wallet.EnabledAt = &now
	_, err = stmt.Exec(wallet.Status, wallet.EnabledAt, wallet.ID)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error updating wallet: %s", err)
	}
	return wallet, nil
}

// Update row disable data
func (w Wallet) Disable(wallet entity.Wallet) (entity.Wallet, error) {
	db, err := w.db.Connect()
	if err != nil {
		return entity.Wallet{}, err
	}
	defer db.Close()

	// prepare the SQL statement
	stmt, err := db.Prepare("UPDATE wallets SET status=$1, disabled_at=$2 WHERE id=$3")
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error preparing statement: %s", err)
	}
	defer stmt.Close()

	// execute the SQL statement with the updated data
	wallet.Status = "disabled"
	now := time.Now()
	wallet.DisabledAt = &now
	_, err = stmt.Exec(wallet.Status, wallet.DisabledAt, wallet.ID)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error updating wallet: %s", err)
	}
	return wallet, nil
}
