package posts

import (
	"database/sql"
	"time"
)

type PostModel struct {
	ID        string    `json:"id" db:"-"`
	UserID    string    `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Post      string    `json:"post" db:"post"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type PostRelationshipModel struct {
	PostModel
	Author    ShortUser                   `json:"author"`
	Reactions ReactionDto                 `json:"reactions"`
	Comments  []*CommentRelationshipModel `json:"comments"`
}

type ReactionDto struct {
	Likes    int64 `json:"likes" db:"likes"`
	Dislikes int64 `json:"dislikes" db:"dislikes"`
}

type ShortUser struct {
	ID   string `json:"id" db:"user.id"`
	Name string `json:"name" db:"user.name"`
}

type PostReactionModel struct {
	ID     string `json:"id" db:"-"`
	PostID string `json:"post_id" db:"post_id"`
	UserID string `json:"user_id" db:"user_id"`
	Like   bool   `json:"like" db:"like"`
}

type CommentModel struct {
	ID        string    `json:"id" db:"-"`
	UserID    string    `json:"user_id" db:"user_id"`
	PostID    string    `json:"post_id" db:"post_id"`
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CommentRelationshipModel struct {
	ID        string         `json:"id" db:"-"`
	PostID    string         `json:"post_id" db:"post_id"`
	Comment   string         `json:"comment" db:"comment"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	Author    ShortUser      `json:"author"`
	ParentID  sql.NullString `json:"-"`

	Comments  []*CommentRelationshipModel `json:"comments"`
	Reactions []CommentReactionModel      `json:"reactions"`
}

type CommentReactionModel struct {
	ID        string `json:"id" db:"-"`
	CommentID string `json:"comment_id" db:"comment_id"`
	UserID    string `json:"user_id" db:"user_id"`
	Like      bool   `json:"like" db:"like"`
}
