package repository

import (
	"fmt"

	"github.com/cetRide/api-rideyu/model"
)

func (c *conn) CreateAccount() model.User {

	sqlStatement := `
		INSERT INTO test (firstname, lastname)
		VALUES ($1, $2)`

	_, err := c.db.Exec(sqlStatement, "Jonathan", "Calhoun")

	if err != nil {
		fmt.Print("error3")
		panic(err)
	}
	return model.User{}
}
