package posts

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repository struct{}

func GetRepository() Repository {
	return Repository{}
}

func (repo Repository) insertPost(tx *sqlx.Tx, post *PostModel) (sql.Result, error) {
	insertStmt := `
	INSERT INTO huddlet.user_post(user_id, title, post)
	VALUES(:user_id, :title, :post)
	`
	return tx.NamedExec(insertStmt, post)
}

func (repo Repository) insertPostReaction(tx *sqlx.Tx, post *PostReactionModel) (sql.Result, error) {
	insertStmt := `
	INSERT INTO huddlet.user_post_reaction(user_id, post_id, "like")
	VALUES(:user_id, :post_id, :like)
	`
	return tx.NamedExec(insertStmt, post)
}

func (repo Repository) insertComment(tx *sqlx.Tx, comment *CommentModel) (sql.Result, error) {
	insertStmt := `
	INSERT INTO huddlet.user_comment(user_id, post_id, comment)
	VALUES(:user_id, :post_id, :comment)
	`
	return tx.NamedExec(insertStmt, comment)
}

func (repo Repository) fetchPosts(db *sqlx.DB) ([]*PostRelationshipModel, error) {
	query := `
	SELECT up.id, up.title, up.post, up.created_at, ua.id, ua.name,
		SUM(CASE WHEN upr.like THEN 1 ELSE 0 END) as likes,
		SUM(CASE WHEN upr.like = false THEN 1 ELSE 0 END) as dislikes
	FROM huddlet.user_post up
	INNER JOIN huddlet.user_account ua
		ON ua.id=up.user_id
	LEFT JOIN huddlet.user_post_reaction upr
		ON upr.post_id = up.id
	GROUP BY up.id, ua.id
	ORDER BY up.created_at DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("fetchPosts Error in query", err)
	}
	defer rows.Close()

	var posts []*PostRelationshipModel
	for rows.Next() {
		var post PostRelationshipModel
		err := rows.Scan(&post.ID, &post.Title, &post.Post, &post.CreatedAt, &post.Author.ID, &post.Author.Name, &post.Reactions.Likes, &post.Reactions.Dislikes)
		if err != nil {
			log.Println("fetchPosts Error in scan", err)
		}
		posts = append(posts, &post)
	}
	return posts, err
}

func (repo Repository) fetchComments(db *sqlx.DB, postIds []string) (map[string][]*CommentRelationshipModel, error) {
	query := `
	SELECT uc.id, uc.post_id, uc.comment, uc.created_at, ua.id, ua.name, pcc.parent_id
	FROM huddlet.user_comment uc
	INNER JOIN huddlet.user_account ua
		ON ua.id=uc.user_id
	LEFT JOIN huddlet.parent_child_comment pcc
		ON pcc.child_id = uc.id
	INNER JOIN UNNEST(CAST($1 as text[])) pid
		ON uc.post_id=uuid(pid)
	
	`
	postComments := make(map[string][]*CommentRelationshipModel)

	rows, err := db.Query(query, pq.Array(&postIds))
	if err != nil {
		log.Println("fetchComments Error in query", err)
	}
	defer rows.Close()

	references := make(map[string]*CommentRelationshipModel)

	for rows.Next() {
		var comment CommentRelationshipModel
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.Comment, &comment.CreatedAt, &comment.Author.ID, &comment.Author.Name, &comment.ParentID)
		if err != nil {
			log.Println("fetchComments Error in scan", err)
		}

		if comments, ok := postComments[comment.PostID]; ok {
			postComments[comment.PostID] = append(comments, &comment)
		} else {
			postComments[comment.PostID] = []*CommentRelationshipModel{&comment}
		}

		references[comment.ID] = &comment
	}

	for _, comments := range postComments {
		for _, comment := range comments {
			if comment.ParentID.Valid {
				references[comment.ParentID.String].Comments = append(references[comment.ParentID.String].Comments, comment)
			}
		}
	}

	return postComments, err
}
