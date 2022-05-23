package controllers

import (
	"net/http"
	"regexp"
)

type empController struct {
	empIDPattern *regexp.Regexp
}

func (em empController) ServeHTTP(W http.ResponseWriter, r *http.Request){
	W.Write([]byte("Hello from Em con"))
}

func newempController()*empController{
	return &empController{
		empIDPattern: regexp.MustCompile("^/users/(\\d+)/?"),
	}
}
