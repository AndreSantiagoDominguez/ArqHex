package controllers

import (
	"net/http"
	"time"

	application "proyecto_hex/Purchases/application"
	"proyecto_hex/Purchases/domain"
	"proyecto_hex/Purchases/infraestructure"

	"github.com/gin-gonic/gin"
)

// CreatePurchaseController gestiona la creación de compras.
type CreatePurchaseController struct {
	app *application.CreatePurchase
}

// NewCreatePurchaseController instancia el controlador utilizando la capa de aplicación.
func NewCreatePurchaseController() *CreatePurchaseController {
	mysql := infraestructure.GetMySQL()
	app := application.NewCreatePurchase(mysql)
	return &CreatePurchaseController{app: app}
}

// addPurchase procesa la solicitud para crear una compra.
// Se espera que el front envíe un JSON con user_id, product_id y quantity.
func (cp_c *CreatePurchaseController) AddPurchase(c *gin.Context) {
	// Estructura auxiliar para el binding del JSON.
	var req struct {
		UserID    int64 `json:"user_id" binding:"required"`
		ProductID int64 `json:"product_id" binding:"required"`
		Quantity  int   `json:"quantity" binding:"required,min=1"`
	}

	// Se realiza el binding y validación de la solicitud.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Solicitud inválida",
			"message": err.Error(),
		})
		return
	}

	// Se crea la compra utilizando el constructor de la capa de dominio.
	purchase := domain.NewPurchase(req.UserID, req.ProductID, req.Quantity)
	// Se establece la fecha actual en formato "YYYY-MM-DD" para coincidir con la columna VARCHAR(10).
	purchase.Date = time.Now().Format("2006-01-02")

	// Se procesa la compra: aquí se envía el mensaje a la cola RabbitMQ a través de la capa de aplicación.
	if err := cp_c.app.Shopping(purchase); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al procesar la compra",
			"message": err.Error(),
		})
		return
	}

	// Se retorna una respuesta exitosa al front.
	// El front posteriormente recibirá el estado actualizado de la compra (confirmada o rechazada) a través de la segunda API.
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Compra enviada a la cola",
		"purchase": purchase,
	})
}
