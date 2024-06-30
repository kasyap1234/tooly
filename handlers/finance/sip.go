package finance

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)
type SIP struct {
	
	MonthlyInvestment float64
	Duration int 
	Interest float64
	Total float64 
}
func SIPCalculator(sip *SIP) float64 {
    annualRate := sip.Interest / 100
    monthlyRate := math.Pow((1+annualRate), 1/12) // Calculate monthly rate

    for i := 0; i < sip.Duration*12; i++ { // Assuming duration is in years, convert to months
        sip.Total += sip.MonthlyInvestment * monthlyRate
    }

    return sip.Total
}

func SIPHandler(w http.ResponseWriter, r *http.Request) {
	var sip SIP ; 
	if err :=json.NewDecoder(r.Body).Decode(&sip); err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	finalValue :=SIPCalculator(&sip)
   w.Write([]byte(fmt.Sprintf("%.2f",finalValue)))

}
