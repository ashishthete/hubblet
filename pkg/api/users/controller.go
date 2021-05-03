package users

import (
	"encoding/json"
	"huddlet/utils"
	"net/http"
)

var list = []UserModel{
	{
		ID:       "1",
		Name:     "Ashish",
		Username: "Thete",
		Email:    "test@xyx",
	},
}

type Controller struct{}

func (ctrl Controller) List(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, http.StatusOK, list)
}

func (ctrl Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var request UserModel
	json.NewDecoder(r.Body).Decode(&request)
	list = append(list, request)
	utils.JsonResponse(w, http.StatusOK, request)
}
