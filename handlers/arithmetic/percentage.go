package arithmetic

import (
	"encoding/json"
	// "math"
	"net/http"
	"github.com/kasyap1234/golang/alltools/tooly/handlers/model"
	
	// "github.com/go-chi/chi"
)

func CalculatePercentageValue(a float64, b float64) float64 {
	return (a / b) * 100
}
func PercentageOfValue(a float64 , b float64 ) float64{
	return b* (a / 100)
}
func CalculatePercentageValueHandler(w http.ResponseWriter, r *http.Request) float64{
  var percentageInput model.PercentageInput;
if err :=json.NewDecoder(r.Body).Decode(&percentageInput); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return 0 ; 
}

return CalculatePercentageValue(percentageInput.A, percentageInput.B)

}
func PercentageOfValueHandler(w http.ResponseWriter, r *http.Request) float64{
  var percentageInput model.PercentageInput;
if err :=json.NewDecoder(r.Body).Decode(&percentageInput); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return 0 ; 
}
return PercentageOfValue(percentageInput.A, percentageInput.B)
}



