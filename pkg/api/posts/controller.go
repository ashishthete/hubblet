package posts

import (
	"encoding/json"
	"huddlet/utils"
	"net/http"
)

type Controller struct{}

func (ctrl Controller) AddPost(w http.ResponseWriter, r *http.Request) {
	var request PostModel
	json.NewDecoder(r.Body).Decode(&request)

	GetServices().AddPost(&request)
	utils.JsonResponse(w, http.StatusOK, request)
}

func (ctrl Controller) AddPostReaction(w http.ResponseWriter, r *http.Request) {
	var request PostModel
	json.NewDecoder(r.Body).Decode(&request)
	utils.JsonResponse(w, http.StatusOK, request)
}
