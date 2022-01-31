package main

import (
	"github.com/anilkusc/flow-diet-backend/app"
)

// @title Flow-Diet-Backend API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	app := app.App{}
	app.Start()

}
