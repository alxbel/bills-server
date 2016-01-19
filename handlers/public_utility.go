package handlers

import (
	"net/http"
	"kommunalka-server/repo"
	"encoding/json"
	"log"
)

func (c *AppContext) PublicUtilitiesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PUHandler")
	repo := repo.PublicUtilityRepo{c.DB.C("pu_catalog")}
	pus, err := repo.All()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(pus.Data)
}
