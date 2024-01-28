package main

import (
	"encoding/json"
	"net/http"

	"github.com/gargdaya/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	var params parameters

	err := json.NewDecoder((r.Body)).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})

	respondWithJson(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request){
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadGateway, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}