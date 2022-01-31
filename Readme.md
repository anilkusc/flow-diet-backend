# Swagger 
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template


add followings to above of main
// @title Orders API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

add followings to the handler
// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]

swag init

add followings to the page of init routers

	_ "swaggo-orders-api/docs" 

	httpSwagger "github.com/swaggo/http-swagger"

add follownig to routes
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

go 
http://localhost:8080/swagger/index.html