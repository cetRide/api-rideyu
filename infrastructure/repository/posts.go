package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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

func (c *conn) FetchComments(ctx context.Context, postId int64) ([]*model.FetchedComment, error) {
	re := strings.NewReplacer("{", "", "}", "")
	sqlStatement := `
	WITH RECURSIVE node_rec AS (
		(SELECT 1 AS depth, ARRAY[id] AS path, *
		 FROM   comments
		 WHERE  parent_comment_id = 0
		)    
		 UNION ALL
		 SELECT r.depth + 1, r.path || n.id, n.*
		 FROM   node_rec r 
		 JOIN   comments    n ON n.parent_comment_id = r.id
		 )
		 SELECT node_rec.path, node_rec.id, node_rec.comment, node_rec.created_at, 
		 node_rec.parent_comment_id, node_rec.user_id,
		 users.username, users.profile_picture
		 FROM   node_rec
		 INNER JOIN users ON node_rec.user_id = users.id
		WHERE post_id = $1
		 ORDER  BY path, created_at`

	var comments []*model.FetchedComment

	rows, err := c.db.QueryContext(ctx, sqlStatement, postId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment model.FetchedComment
		err := rows.Scan(&comment.Path,
			&comment.ID,
			&comment.Comment, &comment.CreatedAt, &comment.ParentCommentId,
			&comment.User_id, &comment.Username, &comment.ProfilePicture,
		)
		if err != nil {
			return nil, err
		}

		comment.Path = re.Replace(comment.Path)

		comments = append(comments, &comment)
	}

	return comments, nil
}

func (c *conn) FetchPosts(ctx context.Context) ([]*model.FetchedPosts, error) {

	sqlStatement := `
	SELECT posts.id, posts.description, users.username, 
	posts.user_id, to_date(posts.created_at::TEXT,'YYYY-MM-DD'), users.profile_picture, posts_media.id, posts_media.file_url
	FROM posts
	INNER JOIN users ON posts.user_id = users.id
	LEFT JOIN posts_media ON posts.id = posts_media.post_id
	ORDER BY posts.created_at`

	rows, err := c.db.QueryContext(ctx, sqlStatement)

	if err != nil {
		return nil, err
	}
	var posts []*model.FetchedPosts
	for rows.Next() {
		var post model.FetchedPosts
		var postMedia model.PostMedia
		err := rows.Scan(&post.ID, &post.Description, &post.Username, &post.User_id, &post.CreatedAt,
			&post.ProfilePicture, &postMedia.Id, &postMedia.FileUrl,
		)
		if err != nil {
			return nil, err
		}
		number_of_posts := len(posts)

		if number_of_posts == 0 || posts[number_of_posts-1].ID != post.ID {
			post.PostMedia = append(post.PostMedia, postMedia)
			posts = append(posts, &post)
		} else {
			posts[number_of_posts-1].PostMedia = append(posts[number_of_posts-1].PostMedia, postMedia)
		}

	}

	return posts, nil
}
