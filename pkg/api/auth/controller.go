package auth

import "net/http"

type Controller struct{}

func (ctrl Controller) Authenticate(w http.ResponseWriter, r *http.Request) {}
