package infraestructure

import (
	"fmt"
	"proyecto_hex/Purchases/domain"
	"proyecto_hex/core"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones:", conn.Err)
	}

	return &MySQL{conn: conn}
}


func (mysql *MySQL) Shopping(purchase *domain.Purchase) (uint, error) {
	query := `
		INSERT INTO purchases (id_user, id_productos, cantidad, date, status)
		VALUES (?, ?, ?, ?, ?)
	`
	
	result, err := mysql.conn.DB.Exec(query,
		purchase.UserID,
		purchase.ProductID,
		purchase.Cantidad,
		purchase.Date,
		purchase.Status,
	)
	if err != nil {
		return 0, fmt.Errorf("error al insertar la compra: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener el ID insertado: %w", err)
	}

	return uint(id), nil
}
