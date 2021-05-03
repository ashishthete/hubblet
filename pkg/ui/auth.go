package ui

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const CookieName = "access_token"

var jwtTokenSecret = "abc123456def"

func CreateToken(user_id string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user_id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtTokenSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractClaims(r *http.Request) (jwt.MapClaims, error) {

	// get our token string from Cookie
	biscuit, err := r.Cookie(CookieName)
	if err != nil {
		log.Println("ExtractClaims", err)
		return nil, err
	}

	token, err := jwt.Parse(biscuit.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid claims type")
	}
	return claims, nil
}

func SetCookieTokenAndRedirect(id string) *http.Cookie {
	tokenString, err := CreateToken(id)
	if err != nil {
		log.Println("Error", err)
	}
	expirationTime := time.Now().Add(time.Hour)
	// Set the cookie with the JWT
	return &http.Cookie{
		Name:     CookieName, // you have to search the cookie by this name
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   false,
	}
}

func SetTokenAndRedirect(w http.ResponseWriter, r *http.Request) {
	cookie := SetCookieTokenAndRedirect("test")
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
}
