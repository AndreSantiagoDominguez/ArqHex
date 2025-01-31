package infraestructure

import (
	"fmt"
	"proyecto_hex/Products/domain"
	"proyecto_hex/core"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

// Edit implements domain.Iproduct.
func (mysql *MySQL) Edit(Product *domain.Product, id int64) (uint, error) {
	panic("unimplemented")
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(product *domain.Product) (uint, error) {
	query := "INSERT INTO productos (name, price) VALUES (?,?)"

	res, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint(id), nil
}

func (mysql *MySQL) GetAllProduct() []domain.Product {
	query := "SELECT * FROM productos"
	var product []domain.Product

	rows := mysql.conn.FetchRows(query)

	if rows == nil {
		fmt.Println("No se pudieron obtener los datos.")
		return product
	}

	defer rows.Close()

	for rows.Next() {
		var b domain.Product
		rows.Scan(&b.ID, &b.Name, &b.Price)

		product = append(product, b)
	}

	return product
}

func (mysql *MySQL) EditProduct(id int, product domain.Product) (uint, error) {
	query := "UPDATE productos SET name = ?, price = ? WHERE id = ?"

	res, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price, id)
	if err != nil {
		return 0, err
	}

	rows, _ := res.RowsAffected()
	return uint(rows), nil
}

func (mysql *MySQL) DeleteProduct(id int) (uint, error) {
	query := "DELETE FROM productos WHERE id = ?"

	res, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return 0, err
	}

	rows, _ := res.RowsAffected()
	return uint(rows), nil
}