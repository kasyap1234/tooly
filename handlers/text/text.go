package handlers

import (
	"encoding/json"
	
	"net/http"
	"strings"
)

func CapitaliseText(text string ) string {
 return strings.ToUpper(text); 

}
func CapitaliseHandler(w http.ResponseWriter , r *http.Request){
var inputData string 
if err :=json.NewDecoder(r.Body).Decode(&inputData); err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return 
}

w.Write([]byte(CapitaliseText(inputData)))
}

