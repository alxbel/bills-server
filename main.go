package main

import (
	"gopkg.in/mgo.v2"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"net/http"
	"bills-server/handlers"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appC := handlers.AppContext{session.DB("bills")}
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler, acceptHandler)

	router := NewRouter()
	router.Get("/bills/:year", commonHandlers.ThenFunc(appC.BillsHandler))
	router.Get("/rows/:billId", commonHandlers.ThenFunc(appC.BillRowsHandler))
	router.Get("/PUCatalog", commonHandlers.ThenFunc(appC.PublicUtilitiesHandler))
	router.Delete("/deleteBill/:id", commonHandlers.ThenFunc(appC.DeleteBillHandler))
	router.Delete("/deleteBillRow/:bid/:rid", commonHandlers.ThenFunc(appC.DeleteRowHandler))

	http.ListenAndServe(":8080", router)
}