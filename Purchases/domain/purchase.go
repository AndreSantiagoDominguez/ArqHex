package domain

type Purchase struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"id_user"`
	ProductID int64  `json:"id_producto"`
	Cantidad  int    `json:"cantidad"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}

func NewPurchase(userID, productID int64, cantidad int) *Purchase {
	return &Purchase{
		UserID:    userID,
		ProductID: productID,
		Cantidad:  cantidad,
		Date:      "23/12/25",
		Status:    "PENDING",
	}
}

func (p *Purchase) SetStatus(status string) {
	p.Status = status
}
