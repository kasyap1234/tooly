package arithmetic 


import (
	"encoding/json"
	"net/http"
	
)
func IsPrime(n int) bool {

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func PrimeHandler(w http.ResponseWriter , r *http.Request)bool {
	var n int 
	
	if err :=json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false; 
	}
	return IsPrime(n)
}




