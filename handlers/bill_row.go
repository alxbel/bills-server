package handlers
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
	"kommunalka-server/repo"
	"encoding/json"
	"log"
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
