package services

import (
	"proyecto_hex/Purchases/application/repository"
	"proyecto_hex/Purchases/domain"
)

type PurchaseEvent struct {
	rmq repository.IRabbit
}

// NewPurchaseEvent crea una nueva instancia de PurchaseEvent.
func NewPurchaseEvent(rmq repository.IRabbit) *PurchaseEvent {
	return &PurchaseEvent{rmq: rmq}
}

// Run envía el evento de compra a la cola.
func (s *PurchaseEvent) Run(purchase domain.Purchase) error {
	return s.rmq.SendPurchase(purchase)
}
