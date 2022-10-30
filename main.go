package main

import (
	"net/http"

	"github.com/seesharp/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
