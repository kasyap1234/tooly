package main

import (
	"fmt"
	"net/http"

	
  
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
	r.Mount("/code", codeHanlders())
	r.Mount("/geometry", geometryHandlers())
	r.Mount("/randomiser", randomiserHandlers())
	// r.Mount()	
	fmt.Println("server started")
	http.ListenAndServe(":3000", r)
}

func codeHanlders() chi.Router {
	r := chi.NewRouter()
	r.Get("/validatejson", handlers.ValidateJsonHandler)
	return r
}
func geometryHandlers() chi.Router {
	r := chi.NewRouter()
	r.Post("/circlearea", handlers.CircleAreaHandler)
	r.Post("/circlecircumference", handlers.CircleCircumferenceHandler)
	return r

}
func randomiserHandlers() chi.Router {
 r :=chi.NewRouter()
 r.Get("/toss",handlers.TossHandler)
 r.Get("/rockpaperscissors",handlers.RockPaperScissorsHandler)
 return r 


}
