package finance

import (
	"encoding/json"
	"fmt"
	
	"net/http"
	"strconv"
	
)

type SIP struct {
	MonthlyInvestment string `json:"MonthlyInvestment"`
	Duration string         `json:"Duration"`
	Interest string			`json:"Interest"`
	Total string  			`json:"Total"`
}
func SIPCalculator(sip *SIP) string{
    
    
   stringInvestment,err := strconv.Atoi(sip.MonthlyInvestment);

   fmt.Println(stringInvestment)
    
   if err!=nil {
	fmt.Printf("Error: %v\n", err)
		return ""
   }

    rate,err :=strconv.ParseFloat(sip.Interest,64); 
	if err!= nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}
	intDuration,err :=strconv.Atoi(sip.Duration); 
	if err!=nil {
		fmt.Printf("Error: %v\n", err)
		return "" 
	}
    // monthlyRate := math.Pow((1+rate), 1/12.0) // Calculate monthly rate
    // intTotal,_ := strconv.ParseFloat(sip.Total,64); 
//    fv = p × ({[1 + i]n – 1} / i) × (1 + i).
    // intTotal +=float64(stringInvestment)*(monthlyRate * float64(intDuration))
   intTotal := float64(stringInvestment) * (((1 + rate) * float64(intDuration)*12 - 1) / rate) * (1 + rate)

    stringValue := strconv.FormatFloat(intTotal, 'f', 2, 64)
	return stringValue; 

}

func SIPHandler(w http.ResponseWriter, r *http.Request) {
	var sip SIP ; 
	if err :=json.NewDecoder(r.Body).Decode(&sip); err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("Error: %v\n", err)
		return
	}
	stringValue :=SIPCalculator(&sip)
   w.Write([]byte(fmt.Sprintf(stringValue)))

}


