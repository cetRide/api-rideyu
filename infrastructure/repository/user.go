package repository

import (
	"github.com/cetRide/api-rideyu/model"
)

func (c *conn) CreateAccount() model.User {

	sqlStatement := `
		INSERT INTO test (firstname, lastname)
		VALUES ($1, $2)`

	_, err := c.db.Exec(sqlStatement, "Jonathan", "Calhoun")

	if err != nil {
		panic(err)
	}
	return model.User{}
}
