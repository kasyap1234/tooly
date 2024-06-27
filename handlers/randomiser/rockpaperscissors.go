package randomiser

import ("math/rand"
		"net/http"
)
func GenerateRockPaperScissors() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(3)]
}
func RockPaperScissorsHandler(w http.ResponseWriter , r *http.Request)string{

	return GenerateRockPaperScissors(); 


}


