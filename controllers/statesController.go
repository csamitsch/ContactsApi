package controllers

import (
	u "github.com/user/contactsApi/utils"
	"net/http"
)

var GetStates = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "success")
	resp["data"] = u.ReadStatesFromCsv()
	u.Respond(w, resp)
}
