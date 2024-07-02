package main

import (
	"fmt"
	"net/http"

	"github.com/kasyap1234/tooly/middlewares"

	"github.com/kasyap1234/tooly/handlers/code"
	"github.com/kasyap1234/tooly/handlers/finance"
	"github.com/kasyap1234/tooly/handlers/geometry"
	"github.com/kasyap1234/tooly/handlers/randomiser"
     "github.com/kasyap1234/tooly/handlers/finance"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()
	
	r.Use(middlewares.LoggerMiddleware)
	r.Use(middleware.Recoverer)


	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Mount("/code", codeHandlers())
	r.Mount("/geometry", geometryHandlers())
	r.Mount("/randomiser", randomiserHandlers())
	
	// r.Mount()	
	fmt.Println("server started")
	http.ListenAndServe(":3000", r)
}
func fianceHanlders() chi.Router {
	r :=chi.NewRouter(); 
	r.Post("/SIP",finance.SIPHandler()); 

}
func codeHandlers() chi.Router {
	r := chi.NewRouter()
	r.Post("/validatejson", code.ValidateJsonHandler)
	return r
}
func geometryHandlers() chi.Router {
	r := chi.NewRouter()
	r.Get("/circle", geometry.CircleAreaHandler)
	r.Get("/circle/circumference", geometry.CircleCircumferenceHandler);
	return r 

}
func randomiserHandlers() chi.Router {
 r :=chi.NewRouter()
 r.Get("/toss",randomiser.TossHandler)
 r.Get("/rockpaperscissors",randomiser.RockPaperScissorsHandler)
 return r 


}
