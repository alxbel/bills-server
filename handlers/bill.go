package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
	"net/http"
	"kommunalka-server/repo"
	"log"
)

func (c *AppContext) BillsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("BillsHandler")
	params := context.Get(r, "params").(httprouter.Params)
	year := params.ByName("year")
	repo := repo.BillRepo{c.DB.C("bills")}
	billsC, err := repo.All(year)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(billsC.Data)
}