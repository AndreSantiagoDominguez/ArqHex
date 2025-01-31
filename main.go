package main

import (
	"proyecto_hex/Products/infraestructure"
	"proyecto_hex/Products/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infraestructure.GoMySQL()
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run() 
}
