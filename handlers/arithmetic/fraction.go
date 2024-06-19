package arithmetic

import (
	"encoding/json"
	"net/http"
	"fmt"
	"math/big"
)


func AddFraction(a int , b int , c int , d int )(int , int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	sum := new(big.Rat).Add(frac1, frac2)

	return sum.Num(),sum.Den(); 


}
func SubFraction(a int , b int , c int , d int )(int , int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	sub := new(big.Rat).Sub(frac1, frac2)

	return sub.Num(),sub.Den();
}
func MulFraction(a int , b int , c int , d int )(int , int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	mul := new(big.Rat).Mul(frac1, frac2)

	return mul.Num(),mul.Den();
}
func DivFraction(a int , b int , c int , d int )(int , int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	div := new(big.Rat).Quo(frac1, frac2)		

	return div.Num(),div.Den();
}


func AddFractionHandler(w http.ResponseWriter , r *http.Request){
	var a int 
	var b int 
	var c int 
	var d int 

	if err :=json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	return AddFraction(a,b,c,d)
}

func SubFractionHandler(w http.ResponseWriter , r *http.Request){
	var a int 
	var b int 
	var c int 
	var d int	

	if err :=json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	return SubFraction(a,b,c,d)
}
func MulFractionHandler(w http.ResponseWriter , r *http.Request){
	var a int 
	var b int 
	var c int 
	var d int

	if err :=json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	return MulFraction(a,b,c,d)
}

func DivFractionHandler(w http.ResponseWriter , r *http.Request){
	var a int 
	var b int 
	var c int 
	var d int

	if err :=json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	if err :=json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0,0
	}
	return DivFraction(a,b,c,d)
}	
