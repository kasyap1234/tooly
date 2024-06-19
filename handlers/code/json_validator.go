package handlers 

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func ValidateJsonHandler(w http.ResponseWriter, r *http.Request) {
    var data interface{}
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() // Disallow unknown fields

    err := decoder.Decode(&data)
    if err != nil {
        syntaxError, ok := err.(*json.SyntaxError)
        if ok {
            http.Error(w, fmt.Sprintf("Syntax error at byte offset %d", syntaxError.Offset), http.StatusBadRequest)
            return
        }

        typeError, ok := err.(*json.UnmarshalTypeError)
        if ok {
            http.Error(w, fmt.Sprintf("Type error: expected %v but got %v at field %s", typeError.Type, typeError.Value, typeError.Field), http.StatusBadRequest)
            return
        }

        http.Error(w, fmt.Sprintf("JSON error: %v", err), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("JSON is valid!"))
}