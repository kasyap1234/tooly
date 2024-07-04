package handlers 

import (
	"encoding/json"
	"net/http"
)

func GenerateImageAI(input string)string {

return input ; 

}
func ImageHandler(w http.ResponseWriter , r *http.Request){
	var input string 

	if err :=json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	w.Write([]byte(GenerateImageAI(input)))
}
