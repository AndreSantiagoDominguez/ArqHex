package main

import (
	productInfra "proyecto_hex/Products/infraestructure"
	productRoutes "proyecto_hex/Products/infraestructure/routes"

	userInfra "proyecto_hex/Users/infraestructure"
	userRoutes "proyecto_hex/Users/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar MySQL para Products y Users
	productInfra.GoMySQL() 
	userInfra.GoMySQL() 

	r := gin.Default()

	productRoutes.RegisterRoutes(r)
	userRoutes.RegisterUserRoutes(r)
	
	

	r.Run(":8800")
}
