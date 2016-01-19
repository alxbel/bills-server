package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
	"encoding/json"
	"log"
	"bills-server/repo"
)

func (c *AppContext) BillRowsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("BillRowsHandler")
	params := context.Get(r, "params").(httprouter.Params)
	billId := params.ByName("billId")
	repo := repo.BillRowRepo{c.DB.C("bills")}
	rowsC, err := repo.All(billId)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(rowsC.Data)
}

func (c *AppContext) DeleteRowHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := repo.BillRowRepo{c.DB.C("bills")}
	if err := repo.Delete(params.ByName("bid"), params.ByName("rid")); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(true)
}