package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

// general internal server error
// log error instead of print in response
func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	res := Response{
		Status: "fail",
		Data:   Error{Error: "error happened in server"},
	}
	log.Println(err.Error())
	json.NewEncoder(w).Encode(res)
	return
}

// bad request error
func BadRequest(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusBadRequest)
	res := Response{
		Status: "fail",
		Data:   Error{Error: err},
	}
	log.Println(err)
	json.NewEncoder(w).Encode(res)
	return
}
