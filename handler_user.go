package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gargdaya/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string         `json:"name"`
		Email sql.NullString `json:"email"`
	}

	// use decoder to get params from request body
	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:    uuid.New(),
		Name:  params.Name,
		Email: params.Email,
	})
	if err != nil {
		respondWithError(w, http.StatusBadGateway, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusBadGateway, "error getting posts for user")
		return
	}
	respondWithJson(w, http.StatusOK, databasePostsToPosts(posts))
}
