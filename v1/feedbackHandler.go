package main

import (
	"encoding/json"
	"net/http"
)

func FeedbackEndpoint(w http.ResponseWriter, r *http.Request) {
	var feedback Feedback
	json.NewDecoder(r.Body).Decode(&feedback)
	RepoUpdateFeedback(feedback)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
