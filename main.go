package main

import (
	// "fmt"
	// "log"
	// "net/http"
	// "os"
	// "regexp"
	// "time"

	// "github.com/gorilla/handlers"
	// "github.com/gorilla/mux"

	// "github.com/ec965/todo-api/config"
	"github.com/ec965/todo-api/models"
	// "github.com/ec965/todo-api/routes"
)

func main() {
	// database
	models.Init()

	// r := mux.NewRouter()
	// // routes
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	ua := r.Header.Get("User-Agent")
	// 	uaRegex := regexp.MustCompile(`(curl|Postman)`)
	// 	if uaRegex.MatchString(ua) {
	// 		w.Write([]byte(time.Now().String()))
	// 	} else {
	// 		timeHtml := "<script>setInterval(function(){ document.getElementById(\"time\").innerHTML = new Date()}, 1000)</script><p id=\"time\">loading...</p>"
	// 		w.Write([]byte(timeHtml))
	// 	}
	// })
	// routes.Init(r)

	// // application middleware
	// var app http.Handler = r
	// app = handlers.ContentTypeHandler(app, "application/x-www-form-urlencoded", "application/json")
	// app = handlers.LoggingHandler(os.Stdout, app)
	// app = handlers.CORS()(app)
	// // app = handlers.RecoveryHandler()(app)

	// s := &http.Server{
	// 	Addr:         ":" + config.Port,
	// 	Handler:      app,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	// fmt.Println("Serving on", s.Addr)
	// log.Fatal(s.ListenAndServe())
}
