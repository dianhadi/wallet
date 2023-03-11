package wallet

import (
	"encoding/json"
	"net/http"

	"github.com/dianhadi/wallet/internal/handler/helper"
)

func (h Handler) View(w http.ResponseWriter, r *http.Request) {
	// header will be set in auth middleware
	walletID := r.Header.Get("Wallet-ID")
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

	res := helper.Response{
		Status: "success",
		Data:   wallet,
	}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}
