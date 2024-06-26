package arithmetic

import (
	"encoding/json"
	"net/http"

	"math/big"
)

func AddFraction(a int, b int, c int, d int) (*big.Int, *big.Int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	sum := new(big.Rat).Add(frac1, frac2)

	return sum.Num(), sum.Denom()

}
func SubFraction(a int, b int, c int, d int) (*big.Int, *big.Int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	sub := new(big.Rat).Sub(frac1, frac2)

	return sub.Num(), sub.Denom()
}
func MulFraction(a int, b int, c int, d int) (*big.Int, *big.Int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	mul := new(big.Rat).Mul(frac1, frac2)

	return mul.Num(), mul.Denom()
}
func DivFraction(a int, b int, c int, d int) (*big.Int, *big.Int) {
	frac1 := big.NewRat(int64(a), int64(b))
	frac2 := big.NewRat(int64(c), int64(d))
	div := new(big.Rat).Quo(frac1, frac2)

	return div.Num(), div.Denom()
}


type Result struct {
	Num big.Int `json:"num"`
	Denom big.Int `json:"denom"`
}



func AddFractionHandler(w http.ResponseWriter, r *http.Request)(Result) {
	var a int 
	var b int 
	var c int 
	var d int 

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	num, denom :=AddFraction(a,b,c,d); 
	return Result{Num : *num , Denom: *denom}

}

func SubFractionHandler(w http.ResponseWriter, r *http.Request)(Result) {
	var a int
	var b int
	var c int
	var d int

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}

	num, denom := SubFraction(a, b, c, d)
	return Result{Num: *num, Denom: *denom}

}
func MulFractionHandler(w http.ResponseWriter, r *http.Request)(Result) {
	var a int
	var b int
	var c int
	var d int

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	num, denom := MulFraction(a, b, c, d)
	return Result{Num: *num, Denom: *denom}

}

func DivFractionHandler(w http.ResponseWriter, r *http.Request)(Result){
	var a int
	var b int
	var c int
	var d int

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return Result{}
	}

	num, denom := DivFraction(a, b, c, d)
	return Result{Num: *num, Denom: *denom}
	
}
