package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NifemiSoneye/rssagg/internal/auth"
	"github.com/NifemiSoneye/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig)handlerCreateUser (w http.ResponseWriter , r *http.Request) {
	type parameters struct {
		Name string `json:name`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w , 400 , fmt.Sprintf("Error parsing JSON :" , err))
		return
	}
	user , err := apiCfg.DB.CreateUser(r.Context() , database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w , 400 , fmt.Sprintf("Couldnt create user :" , err))
		return
	}
	respondWithJSON(w , 201 , databaseUserToUser(user))
}

func (apiCfg *apiConfig)handlerGetUser (w http.ResponseWriter , r *http.Request) {
	apiKey , err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWithError(w , 403 , fmt.Sprintf("Auth error: %v" , err))
		return
	}

	apiCfg.DB.GetUserByApiKey(r.Context() , apiKey)
}