package repository

import "proyecto_hex/Purchases/domain"

// IRabbit define los métodos para interactuar con RabbitMQ 
type IRabbit interface {

	SendPurchase(purchase domain.Purchase) error

}
