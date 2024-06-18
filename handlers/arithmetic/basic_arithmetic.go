package arithmetic

import (
	"encoding/json"
	// "math"
	"net/http"
	"github.com/kasyap1234/alltools/handlers/arithmetic/model"
)
func CalculatePercentage(a float64 , b float64 ) float64{

	return (a / b) * 100
}
func PercentageOf(a float64 , b float64 ) float64{
	return b* (a / 100)
}
func CalculatePercentageHandler(w http.ResponseWriter, r *http.Request) float64{
  var percentageInput model.PercentageInput;
if err :=json.NewDecoder(r.Body).Decode(&percentageInput); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return 0 ; 
}

return CalculatePercentage(percentageInput.A, percentageInput.B)

}
func PercentageOfHandler(w http.ResponseWriter, r *http.Request) float64{
  var percentageInput model.PercentageInput;
if err :=json.NewDecoder(r.Body).Decode(&percentageInput); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return 0 ; 
}
return PercentageOf(percentageInput.A, percentageInput.B)
}



