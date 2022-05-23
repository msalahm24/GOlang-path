package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()
	em := newempController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	http.Handle("/ems", *em)
	http.Handle("/ems/", *em)
}
