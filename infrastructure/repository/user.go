package repository

import (
	"context"

	"github.com/cetRide/api-rideyu/model"
)

func (c *conn) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {

	sqlStatement := `
		INSERT INTO users (username, email, phone, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, phone`

	row := c.db.QueryRowContext(ctx, sqlStatement,
		user.Username,
		user.Email,
		user.Phone,
		user.Password,
	)

	person := model.User{}

	err := row.Scan(
		&person.ID,
		&person.Username,
		&person.Email,
		&user.Phone,
	)
	return &person, err
}

func (c *conn) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	sqlStatement := `SELECT id, username, email, phone, password FROM users WHERE email = $1`
	row := c.db.QueryRowContext(ctx, sqlStatement, email)

	user := model.User{}

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.Password,
	)
	return &user, err
}
func (c *conn) FindByPhone(ctx context.Context, phone string) (*model.User, error) {
	sqlStatement := `SELECT id, username, email, phone, password FROM users WHERE phone = $1`
	row := c.db.QueryRowContext(ctx, sqlStatement, phone)

	user := model.User{}

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.Password,
	)
	return &user, err
}
func (c *conn) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	sqlStatement := `SELECT id, username, email, phone, password FROM users WHERE username = $1`
	row := c.db.QueryRowContext(ctx, sqlStatement, username)

	user := model.User{}

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.Password,
	)
	return &user, err
}

func (c *conn) FindById(ctx context.Context, id int64) (*model.User, error) {
	sqlStatement := `SELECT * FROM users WHERE username = $1`
	row := c.db.QueryRowContext(ctx, sqlStatement, id)

	user := model.User{}

	err := row.Scan(&user)
	return &user, err
}

func (c *conn) Follow(ctx context.Context, follower int64, following int64) (*model.Follower, error) {

	sqlStatement := `
	INSERT INTO followers (follower, following)
		VALUES ($1,$2)
		RETURNING follower, following`

	row := c.db.QueryRowContext(ctx, sqlStatement, follower, following)

	followerModel := model.Follower{}

	err := row.Scan(&followerModel.Follower, &followerModel.Following)

	return &followerModel, err
}

func (c *conn) UnFollow(ctx context.Context, id int64) (map[string]interface{}, error) {

	sqlStatement := `
	DELETE FROM followers WHERE id = $1`

	_, err := c.db.ExecContext(ctx, sqlStatement, id)

	if err != nil {
		return nil, err
	}
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "Unfollowed successfully"
	return response, nil
}

func (c *conn) ViewListOfFollowers(ctx context.Context, id int64) ([]*model.ListOfFollowers, error) {

	sqlStatement := `
	SELECT followers.id AS follower_id, users.id AS user_id, users.username, users.profile_picture
	FROM followers FULL JOIN users ON followers.follower = users.id WHERE following = $1`

	rows, err := c.db.QueryContext(ctx, sqlStatement, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listOfFollowers []*model.ListOfFollowers
	for rows.Next() {
		var follower model.ListOfFollowers
		if err := rows.Scan(
			&follower.FollowerId,
			&follower.UserId,
			&follower.Username,
			&follower.ProfilePicture,
		); err != nil {
			return nil, err
		}
		listOfFollowers = append(listOfFollowers, &follower)

	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listOfFollowers, nil

}

func (c *conn) ViewListOfFollowing(ctx context.Context, id int64) ([]*model.ListOfFollowers, error) {

	sqlStatement := `
	SELECT followers.id AS follower_id, users.id as user_id, users.username, users.profile_picture
	 FROM followers FULL JOIN users ON followers.following = users.id WHERE follower = $1`

	rows, err := c.db.QueryContext(ctx, sqlStatement, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listOfFollowing []*model.ListOfFollowers
	for rows.Next() {
		var follower model.ListOfFollowers
		if err := rows.Scan(
			&follower.FollowerId,
			&follower.UserId,
			&follower.Username,
			&follower.ProfilePicture,
		); err != nil {
			return nil, err
		}
		
		listOfFollowing = append(listOfFollowing, &follower)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listOfFollowing, nil

}
