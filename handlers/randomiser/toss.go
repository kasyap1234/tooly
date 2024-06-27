package randomiser


import ("math/rand"
		"net/http"
		
)
func RandomToss() string {
	choices := []string{"heads", "tails"}
	return choices[rand.Intn(2)]
}
func TossHandler(w http.ResponseWriter , r *http.Request){

	w.Write([]byte(RandomToss()))

}


