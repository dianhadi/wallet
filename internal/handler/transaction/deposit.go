package transaction

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	entity "github.com/dianhadi/wallet/internal/entity/transaction"
	"github.com/dianhadi/wallet/internal/handler/helper"
	"github.com/google/uuid"
)

func (h Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	// header will be set in auth middleware
	walletID := r.Header.Get("Wallet-ID")
	userID := r.Header.Get("User-ID")

	wallet, err := h.repoWallet.GetByID(walletID)
	if err != nil {
		helper.InternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if !wallet.IsEnabled() {
		w.WriteHeader(http.StatusConflict)
		res := helper.Response{
			Status: "fail",
			Data:   helper.Error{Error: "wallet is not enabled yet"},
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if wallet.IsDisabled() {
		w.WriteHeader(http.StatusConflict)
		res := helper.Response{
			Status: "fail",
			Data:   helper.Error{Error: "wallet is already disabled"},
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	r.ParseMultipartForm(32 << 20)
	amount := r.FormValue("amount")
	referenceID := r.FormValue("reference_id")
	w.Header().Set("Content-Type", "application/json")

	if amount == "" || referenceID == "" {
		helper.BadRequest(w, "amount and reference id is required")
		return
	}

	_, err = uuid.Parse(referenceID)
	if err != nil {
		helper.BadRequest(w, "reference id is not valid UUID")
		return
	}

	amountInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		helper.BadRequest(w, "amount is invalid number")
		return
	}

	newUuid := uuid.New()
	transaction := entity.Transaction{
		ID:          newUuid.String(),
		WalletID:    walletID,
		ReferenceID: referenceID,
		Amount:      amountInt,
		Type:        "deposit",
		DepositedBy: &userID,
	}

	refCheck, err := h.repoTransaction.GetIDByReferenceID(transaction.ReferenceID)
	if err != sql.ErrNoRows || refCheck != "" {
		w.WriteHeader(http.StatusConflict)
		res := helper.Response{
			Status: "fail",
			Data:   helper.Error{Error: "reference id is already exist"},
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	transaction, err = h.repoTransaction.Deposit(transaction)
	if err != nil {
		helper.InternalServerError(w, err)
		return
	}
	transaction.Status = "success"
	transaction.Type = ""

	res := helper.Response{
		Status: "success",
		Data:   Transaction{Deposit: &transaction, Withdrawal: nil},
	}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}
