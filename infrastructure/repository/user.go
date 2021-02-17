package repository

import (
	"context"

	"github.com/cetRide/api-rideyu/model"
)

func (c *conn) CreateAccount(ctx context.Context, user *model.User) (*model.User, error) {

	sqlStatement := `
		INSERT INTO test (firstname, lastname)
		VALUES ($1, $2)
		RETURNING firstname, lastname`

	row := c.db.QueryRowContext(ctx, sqlStatement, user.FirstName, user.LastName)

	newUser := model.User{}

	err := row.Scan(
		&newUser.FirstName,
		&newUser.LastName,
	)
	return &newUser, err
}
