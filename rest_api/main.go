package main

import (
	"rest/config"
	"rest/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	config.MigrateDB()
	e := echo.New()

	// Rute untuk registrasi dan login tanpa autentikasi
	e.POST("/api/v1/register", controllers.RegisterHandler)
	e.POST("/api/v1/login", controllers.LoginHandler)

	// Kelompok rute yang memerlukan autentikasi
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("alta")))

	// Routes for Category
	eAuth.GET("/api/v1/categories", controllers.GetAllCategoriesHandler)
	eAuth.GET("/api/v1/categories/:id", controllers.GetCategoryHandler)
	eAuth.POST("/api/v1/categories", controllers.CreateCategoryHandler)
	eAuth.PUT("/api/v1/categories/:id", controllers.UpdateCategoryHandler)
	eAuth.DELETE("/api/v1/categories/:id", controllers.DeleteCategoryHandler)

	// Routes for Product
	eAuth.GET("/api/v1/products", controllers.GetAllProductsHandler)
	eAuth.GET("/api/v1/products/:id", controllers.GetProductHandler)
	eAuth.POST("/api/v1/products", controllers.CreateProductHandler)
	eAuth.PUT("/api/v1/products/:id", controllers.UpdateProductsHandler)
	eAuth.DELETE("/api/v1/products/:id", controllers.DeleteProductHandler)

	// Routes for Transaction
	eAuth.GET("/api/v1/transactions", controllers.GetAllTransactionHandler)
	eAuth.GET("/api/v1/transactions/:id", controllers.GetTransactiontHandler)
	eAuth.POST("/api/v1/transactions", controllers.CreateTransactiontHandler)
	eAuth.PUT("/api/v1/transactions/:id", controllers.UpdateTransactionHandler)
	eAuth.DELETE("/api/v1/transactions/:id", controllers.DeleteTransactionHandler)

	// Memulai server
	if err := e.Start(":8000"); err != nil {
		e.Logger.Fatal("Failed to start server: ", err)
	}
}
