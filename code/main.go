package main

import (
	"net/http"

	"github.com/osuksdev/go-experimentation/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
