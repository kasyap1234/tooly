package geometry 

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func FindArea(radius float64) float64 {
	if radius > 0 {
		return 3.14 * radius * radius
	}
	log.Fatal("invalid radius"); 

	return 0.00
}
func FindCircumference(radius float64) float64 {
	if radius > 0 {
		return 2 * 3.14 * radius
	}
	log.Fatal("invalid radius"); 

	return 0.0
}

func CircleAreaHandler(w http.ResponseWriter, r *http.Request) float64 {
	radius := r.URL.Query().Get("radius")
	floatValue, err := strconv.ParseFloat(radius, 64)
	if err != nil {
		fmt.Println(err)

	}
	return FindArea(floatValue)
}
func CircleCircumferenceHandler(w http.ResponseWriter, r *http.Request) float64 {
	radius := r.URL.Query().Get("radius")
	floatValue, err := strconv.ParseFloat(radius, 64)
	if err != nil {
		fmt.Println(err)
		return 0.000
	}
	return FindCircumference(floatValue)
}

