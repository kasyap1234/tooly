package arithmetic

import (
	"encoding/json"
	"strconv"
	// "math"
	"net/http"

	"github.com/kasyap1234/tooly/handlers/arithmetic/model"
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
func PercentageOfValueHandler(w http.ResponseWriter, r *http.Request) {
  var percentageInput model.PercentageInput;
if err :=json.NewDecoder(r.Body).Decode(&percentageInput); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
w.Write([]byte("0"))
return ; 
}

value :=strconv.FormatFloat(PercentageOfValue(percentageInput.A, percentageInput.B), 'f', 2, 64)
w.Write([]byte(value))
}



