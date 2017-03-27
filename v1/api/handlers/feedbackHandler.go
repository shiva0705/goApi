package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shiva0705/goApi/v1/data"
	"github.com/shiva0705/goApi/v1/models"
)

func FeedbackEndpoint(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback

	json.NewDecoder(r.Body).Decode(&feedback)

	var db = data.DbHandle()
	defer db.Close()

	data.UpdateFeedback(db, feedback)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
