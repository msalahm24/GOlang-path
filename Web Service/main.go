package main

import (
	"net/http"
	"source/controllers"
)
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000",nil)
}
