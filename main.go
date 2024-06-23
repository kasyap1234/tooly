package main 
import (
	"fmt"
	"net/http"
	 "github.com/go-chi/chi"
	
)
func main(){
	 
	r :=chi.NewRouter(); 
	r.Get("/",func (w http.ResponseWriter , r *http.Request){
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":3000", r)
}

