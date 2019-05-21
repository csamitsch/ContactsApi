package controllers

import (
	"github.com/user/contactsApi/models"
	u "github.com/user/contactsApi/utils"
	"net/http"
)

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	contact := &models.Contact{}
	contact.Name = "Christoph"

	var contactList = []models.Contact {
		{
			Name: "John Doe",
		},
		{
			Name: "Jane Doe",
		},
	}

	resp := u.Message(true, "success")
	resp["data"] = contactList
	u.Respond(w, resp)
}