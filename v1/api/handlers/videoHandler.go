package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shiva0705/goApi/v1/data"
)

func Get10VideosEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var db = data.DbHandle()
	defer db.Close()
	if err := json.NewEncoder(w).Encode(data.GetVideos(db)); err != nil {
		panic(err)
	}
}
