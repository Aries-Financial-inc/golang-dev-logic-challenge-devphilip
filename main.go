package main

import (
	. "fmt"
	"golang-dev-logic-challenge-devphilip/routes"
	"log"
	"net/http"
)

// @title           Options Analysis API
// @version         1.0
// @description     This is an options analysis API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Philip Akpeki
// @contact.url    http://www.swagger.io/support
// @contact.email  philipakpeki@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
func main() {
	router := routes.SetupRouter()
	Println("Starting server on port 8080")
	log.Panic(http.ListenAndServe(":8080", router))
}
