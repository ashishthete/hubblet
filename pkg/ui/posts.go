package ui

import (
	"huddlet/pkg/api/posts"
	"huddlet/pkg/api/users"
	"log"
	"net/http"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	conditionsMap := map[string]interface{}{}

	// THIS PAGE should ONLY obe accessible to those users that logged in

	// check if user already logged in
	claims, _ := ExtractClaims(r)
	userID, _ := claims["user_id"].(string)

	if userID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user, err := users.GetServices().GetUserByField("id", userID)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	services := posts.GetServices()
	conditionsMap["account"] = user

	r.ParseForm()
	comment := posts.CommentModel{
		PostID:  r.Form.Get("post_id"),
		UserID:  userID,
		Comment: r.Form.Get("comment"),
	}

	err = services.AddComment(&comment)
	if err != nil {
		http.Error(w, "Error adding comment", 400)
		log.Println("Error", err)
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)

}
