package arithmetic

import ("encoding/json"
	"net/http"
)

func sum(input []float64) float64{
	total :=0.0; 

	for _, value := range input {
		total += value
	}
	return total

}
func CalculateAverage(input []float64) float64{
	return sum(input) / float64(len(input))
}
func AverageHandler(w http.ResponseWriter, r *http.Request) float64{
	var input []float64 ; 
	if err :=json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0.0
	}

	return CalculateAverage(input)
}


