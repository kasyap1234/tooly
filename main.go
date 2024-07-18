package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/cors"
	"github.com/kasyap1234/tooly/middlewares"

	"github.com/kasyap1234/tooly/handlers/code"
	"github.com/kasyap1234/tooly/handlers/finance"
	"github.com/kasyap1234/tooly/handlers/geometry"
	"github.com/kasyap1234/tooly/handlers/notes"
	"github.com/kasyap1234/tooly/handlers/randomiser"
	"github.com/kasyap1234/tooly/handlers/taskmanager"

	"log"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kasyap1234/tooly/db"
)
func connectWithRetry(dsn string, models ...interface{}) {
    var err error
    maxAttempts := 10
    for i := 0; i < maxAttempts; i++ {
          database.ConnectDB(dsn, models...)
        if err == nil {
            return
        }
        log.Printf("Failed to connect to database. Retrying in 5 seconds... (Attempt %d/%d)", i+1, maxAttempts)
        time.Sleep(5 * time.Second)
    }
    log.Fatalf("Could not connect to the database: %v", err)
}
func main() {

	r := chi.NewRouter()
	dsn := "host=localhost user=youruser password=yourpassword dbname=database port=5432 sslmode=disable"


	 connectWithRetry(dsn, &taskmanager.Task{})
	 connectWithRetry(dsn,&notes.Notes{})

     cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},    // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for preflight requests
	})

	// Use CORS middleware
	r.Use(cors.Handler)
	r.Use(middlewares.LoggerMiddleware)
	r.Use(middleware.Recoverer)


	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Mount("/code", codeHandlers())
	r.Mount("/geometry", geometryHandlers())
	r.Mount("/randomiser", randomiserHandlers())
	r.Mount("/finance", financeHandlers())
    r.Mount("/tasksmanager",taskmanagerHandlers())
	r.Mount("/notehandler",notesHandlers())

	
	// r.Mount()	
	fmt.Println("server started")
	http.ListenAndServe(":3000", r)
}
func financeHandlers() chi.Router {
	r :=chi.NewRouter(); 
	r.Post("/SIP",finance.SIPHandler); 
return r 

}
// http://localhost:3000/SIP
func codeHandlers() chi.Router {
	r := chi.NewRouter()
	r.Post("/validatejson", code.ValidateJsonHandler)
	return r
}
func taskmanagerHandlers() chi.Router {
	r := chi.NewRouter()
	r.Get("/tasks", taskmanager.GetTasks(database.DB))
	r.Post("/tasks", taskmanager.CreateTask(database.DB))
	r.Get("/tasks/{id}", taskmanager.GetTask(database.DB))
	r.Put("/tasks/{id}", taskmanager.UpdateTask(database.DB))
	r.Delete("/tasks/{id}", taskmanager.DeleteTask(database.DB))
	return r
}
func geometryHandlers() chi.Router {
	r := chi.NewRouter()
	r.Get("/circle", geometry.CircleAreaHandler)
	r.Get("/circle/circumference", geometry.CircleCircumferenceHandler);
	return r 

}
func notesHandlers() chi.Router {
	r := chi.NewRouter()
	r.Get("/notes", notes.GetNotes(database.DB))
	r.Post("/notes", notes.CreateNotes(database.DB))
	r.Get("/notes/{id}", notes.GetNote(database.DB))
	r.Put("/notes/{id}", notes.UpdateNote(database.DB))
	r.Delete("/notes/{id}", notes.DeleteNote(database.DB))
	return r
}
func randomiserHandlers() chi.Router {
 r :=chi.NewRouter()
 r.Get("/toss",randomiser.TossHandler)
 r.Get("/rockpaperscissors",randomiser.RockPaperScissorsHandler)
 return r 


}
