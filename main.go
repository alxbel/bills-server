package main

import (
	"gopkg.in/mgo.v2"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"net/http"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appC := appContext{session.DB("bills")}
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler, acceptHandler)

	router := NewRouter()
	router.Get("/bills/:year", commonHandlers.ThenFunc(appC.billsHandler))

	http.ListenAndServe(":8080", router)
}