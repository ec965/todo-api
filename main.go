package main

import (
	"net/http"
	"log"
	"time"
	"os"

	"github.com/gorilla/mux"
    "github.com/gorilla/handlers"

	"github.com/ec965/todo-api/routes"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(time.Now().String()))
}


func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func contentTypeMiddleWare(next http.Handler) http.Handler {
	return handlers.ContentTypeHandler(next)
}

func main() {
	r := mux.NewRouter()
	// router middleware
	r.Use(loggingMiddleware)
	r.Use(contentTypeMiddleWare)
	// routes
	r.HandleFunc("/", DefaultHandler)
	routes.Init(r)


	// application middleware
	var app http.Handler = r
	app = handlers.CORS()(app);
	app = handlers.RecoveryHandler()(app)

	s := &http.Server{
		Addr: ":8080",
		Handler: app,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}