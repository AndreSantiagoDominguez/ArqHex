package infraestructure

import (
	"fmt"
	"proyecto_hex/Users/domain"
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

func (mysql *MySQL) EditUser(id int, user domain.User) (uint, error) {
	query := "UPDATE empleados SET first_name = ?, last_name = ?, age = ?, phone_number = ? WHERE id = ?"

	res, err := mysql.conn.ExecutePreparedQuery(query, user.FirstName, user.LastName, user.Age, user.PhoneNumber, id)

	if err != nil {
		return 0, err
	}

	rows, _ := res.RowsAffected()
	return uint(rows), nil
}

func (mysql *MySQL) GetAllUsers() []domain.User {
	query := "SELECT * FROM empleados"
	var user []domain.User

	rows := mysql.conn.FetchRows(query)

	if rows == nil {
		fmt.Println("No se pudieron obtener los datos.")
		return user
	}

	defer rows.Close()

	for rows.Next() {
		var b domain.User
		rows.Scan(&b.Id, &b.LastName, &b.FirstName, &b.Age, &b.PhoneNumber)

		user = append(user, b)
	}

	return user
}


func (mysql *MySQL) Save(user *domain.User) (uint, error) {
	query := "INSERT INTO empleados (first_name, last_name, age, phone_number) VALUES (?, ?, ?, ?)"

	res, err := mysql.conn.ExecutePreparedQuery(query, user.FirstName, user.LastName, user.Age, user.PhoneNumber)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint(id), nil
}

func (mysql *MySQL) DeleteUser(id int) (uint, error) {
	query := "DELETE FROM empleados WHERE id = ?"

	res, err := mysql.conn.ExecutePreparedQuery(query, id)

	if err != nil {
		return 0, err
	}

	rows,_:= res.RowsAffected()
	return uint(rows), nil
}
