package controllers

import (
	"net/http"
	"regexp"
)

type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the user controller"))
}

// if you want to constructor you can create it like this|

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile("^/users/(\\d+)/?"),
	}
}

func (uc *UserController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(w)
}
