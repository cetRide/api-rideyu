package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cetRide/api-rideyu/model"
)

func (c *conn) SavePost(ctx context.Context, post *model.Post) (sql.Result, error) {

	sqlStatement := `
		INSERT INTO posts (user_id, description)
		VALUES ($1, $2)
		RETURNING id, user_id, description`

	result, err := c.db.ExecContext(ctx, sqlStatement,
		post.UserId,
		post.Description,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *conn) SaveComment(ctx context.Context, comment *model.Comment) (sql.Result, error) {

	sqlStatement := `
		INSERT INTO comments (user_id, post_id, comment)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, post_id, comment`

	result, err := c.db.ExecContext(ctx, sqlStatement,
		comment.UserId,
		comment.PostId,
		comment.Comment,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *conn) ReplyComment(ctx context.Context, comment *model.Comment, commentId int64) (sql.Result, error) {
	var child_comment_id int64

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	insertComment := `
	INSERT INTO comments (user_id, post_id, comment)
	VALUES ($1, $2, $3)
	RETURNING id`
	row := c.db.QueryRowContext(ctx, insertComment,
		comment.UserId,
		comment.PostId,
		comment.Comment,
	)

	err = row.Scan(&child_comment_id)
	if err != nil {
		theError := tx.Rollback()
		if theError != nil {
			return nil, err
		}
		return nil, err
	}

	insertParentChildComment := `
	INSERT INTO parent_child_comments (parent_comment_id, child_comment_id)
	VALUES ($1, $2)`
	fmt.Printf("the id =[%v]", child_comment_id)
	result, err := tx.ExecContext(ctx, insertParentChildComment, commentId, &child_comment_id)
	if err != nil {
		theError := tx.Rollback()
		if theError != nil {
			return nil, err
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
