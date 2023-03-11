package wallet

import (
	"database/sql"
	"encoding/json"
	"net/http"

	entity "github.com/dianhadi/wallet/internal/entity/wallet"
	"github.com/dianhadi/wallet/internal/handler/helper"

	"github.com/google/uuid"
)

func (h Handler) Init(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	userID := r.FormValue("customer_xid")
	w.Header().Set("Content-Type", "application/json")

	if userID == "" {
		helper.BadRequest(w, "customer_xid id is required")
		return
	}

	_, err := uuid.Parse(userID)
	if err != nil {
		helper.BadRequest(w, "customer_xid id is not valid UUID")
		return
	}

	wallet := entity.New(userID)

	userCheck, err := h.repoWallet.GetIDByUserID(wallet.UserID)
	if err != sql.ErrNoRows || userCheck != "" {
		w.WriteHeader(http.StatusConflict)
		res := helper.Response{
			Status: "fail",
			Data:   helper.Error{Error: "user is already exist"},
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = h.repoWallet.Insert(wallet)
	if err != nil {
		helper.InternalServerError(w, err)
		return
	}

	res := helper.Response{
		Status: "success",
		Data:   Token{Token: wallet.Token},
	}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}
