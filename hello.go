package main

import (
	"net/http"

	"github.com/wrowston/go-hello-world/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
