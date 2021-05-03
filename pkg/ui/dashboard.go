package ui

import (
	"fmt"
	"html/template"
	"huddlet/pkg/api/posts"
	"huddlet/pkg/api/users"
	"log"
	"net/http"
	"os"
)

func DashBoardPageHandler(w http.ResponseWriter, r *http.Request) {
	conditionsMap := map[string]interface{}{}

	// THIS PAGE should ONLY obe accessible to those users that logged in

	// check if user already logged in
	claims, _ := ExtractClaims(r)
	log.Println("Claims", claims)
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

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./templates/dashboard.html")
		if err != nil {
			log.Println("Error", err)
		}

		posts, err := services.ListPost()
		if err != nil {
			http.Error(w, "user not found", 400)
			log.Println("Error", err)
		}
		conditionsMap["posts"] = posts
		log.Println(conditionsMap)

		if err := tmpl.Execute(w, conditionsMap); err != nil {
			log.Println(err)
		}
		return
	}

	r.ParseForm()
	post := posts.PostModel{
		Post:   r.Form.Get("post"),
		Title:  r.Form.Get("title"),
		UserID: userID,
	}
	err = services.AddPost(&post)
	if err != nil {
		http.Error(w, "Error adding post", 400)
		log.Println("Error", err)
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)

}

func LikePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	fmt.Println(os.Getwd())          //get request method

	claims, _ := ExtractClaims(r)
	log.Println("Claims", claims)
	userID, _ := claims["user_id"].(string)

	if userID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()

		services := posts.GetServices()

		reaction := posts.PostReactionModel{
			PostID: r.Form.Get("post_id"),
			UserID: userID,
			Like:   true,
		}

		err := services.AddPostReaction(&reaction)
		if err != nil {
			log.Println("Error in add user", err)
		}
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	fmt.Println(os.Getwd())          //get request method

	claims, _ := ExtractClaims(r)
	log.Println("Claims", claims)
	userID, _ := claims["user_id"].(string)

	if userID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()

		services := posts.GetServices()

		reaction := posts.PostReactionModel{
			PostID: r.Form.Get("post_id"),
			UserID: userID,
			Like:   false,
		}

		err := services.AddPostReaction(&reaction)
		if err != nil {
			log.Println("Error in add user", err)
		}
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
