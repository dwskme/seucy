package handlers

import (
	"net/http"
)

func (u *Handler) Write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		return
	}
}

func (u *Handler) EnglishHandler(writer http.ResponseWriter, _ *http.Request) {
	u.Write(writer, "Hello World")
}

// update profile
// update email, password
// verify email
