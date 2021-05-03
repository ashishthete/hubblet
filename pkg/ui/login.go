package ui

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"huddlet/pkg/api/users"
	"huddlet/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	conditionsMap := map[string]interface{}{}

	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	if r.Method != http.MethodPost {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		conditionsMap["token"] = token
		tmpl.Execute(w, token)
		return
	}

	r.ParseForm()
	conditionsMap["username"] = r.Form.Get("username")
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	services := users.GetServices()
	user, err := services.GetUserByField("username", username)
	if err != nil {
		http.Error(w, "user not found", 400)
		log.Println("Error", err)
	}

	if utils.ComparePasswords(user.Password, password) {
		cookie := SetCookieTokenAndRedirect(user.ID)
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		http.Error(w, "wrong password", 400)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		services := users.GetServices()

		user := users.UserModel{
			Name:     r.Form.Get("email"),
			Email:    r.Form.Get("email"),
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		err := services.AddUser(&user)
		if err != nil {
			log.Println("Error in add user", err)
		}
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   CookieName,
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
