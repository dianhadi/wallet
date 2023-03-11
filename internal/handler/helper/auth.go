package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h Handler) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		auth := r.Header.Get("Authorization")
		if auth == "" || len(auth) < 7 {
			w.WriteHeader(http.StatusUnauthorized)
			res := Response{
				Status: "fail",
				Data:   Error{Error: "Token is required"},
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		if auth[:5] != "Token" {
			w.WriteHeader(http.StatusUnauthorized)
			res := Response{
				Status: "fail",
				Data:   Error{Error: "Token is invalid"},
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		auth = auth[6:]

		walletID, userID, err := h.repoWallet.GetByToken(auth)
		if err != nil || walletID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			res := Response{
				Status: "fail",
				Data:   Error{Error: "Token is invalid"},
			}
			log.Println(err.Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		r.Header.Set("Wallet-ID", walletID)
		r.Header.Set("User-ID", userID)

		next.ServeHTTP(w, r)
	})

}
