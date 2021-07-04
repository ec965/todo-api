package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/ec965/todo-api/config"
	"github.com/ec965/todo-api/models"
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
	// database
	models.Init()

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
		Addr: ":"+config.Port,
		Handler: app,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Serving on", s.Addr)
	log.Fatal(s.ListenAndServe())
}