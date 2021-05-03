package posts

import (
	"huddlet/app/db"
	"log"
)

type Services struct{}

func GetServices() Services {
	return Services{}
}

func (service Services) AddPost(post *PostModel) error {
	tx, _ := db.GetPostgresDB().Beginx()
	defer tx.Rollback()
	_, err := GetRepository().insertPost(tx, post)
	if err != nil {
		log.Println("Error inserting data", err)
	}

	err = tx.Commit()
	return err
}

func (service Services) AddComment(comment *CommentModel) error {
	log.Println("comment", comment)
	tx, _ := db.GetPostgresDB().Beginx()
	defer tx.Rollback()
	_, err := GetRepository().insertComment(tx, comment)
	if err != nil {
		log.Println("AddComment: Error inserting data", err)
	}

	err = tx.Commit()
	return err
}

func (service Services) AddPostReaction(reaction *PostReactionModel) error {
	tx, _ := db.GetPostgresDB().Beginx()
	defer tx.Rollback()
	_, err := GetRepository().insertPostReaction(tx, reaction)
	if err != nil {
		log.Println("Error inserting data", err)
	}

	err = tx.Commit()
	return err
}

func (service Services) ListPost() ([]*PostRelationshipModel, error) {
	db := db.GetPostgresDB()
	repo := GetRepository()
	posts, err := repo.fetchPosts(db)
	if err != nil {
		log.Println("Error inserting data", err)
	}
	postIds := []string{}
	for _, post := range posts {
		postIds = append(postIds, post.ID)
	}
	comments, err := repo.fetchComments(db, postIds)
	if err != nil {
		log.Println("Error inserting data", err)
	}

	for _, post := range posts {
		if val, ok := comments[post.ID]; ok {
			post.Comments = val
		}
	}

	return posts, err
}
