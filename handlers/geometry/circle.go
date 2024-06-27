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
    log.Fatal("invalid radius")
    return 0.00
}

func FindCircumference(radius float64) float64 {
    if radius > 0 {
        return 2 * 3.14 * radius
    }
    log.Fatal("invalid radius")
    return 0.0
}

func CircleAreaHandler(w http.ResponseWriter, r *http.Request) {
    radius := r.URL.Query().Get("radius")
    floatValue, err := strconv.ParseFloat(radius, 64)
    if err != nil {
        http.Error(w, "Invalid radius", http.StatusBadRequest)
        return
    }
    area := FindArea(floatValue)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf(`{"area": %f}`, area)))
}

func CircleCircumferenceHandler(w http.ResponseWriter, r *http.Request) {
    radius := r.URL.Query().Get("radius")
    floatValue, err := strconv.ParseFloat(radius, 64)
    if err != nil {
        http.Error(w, "Invalid radius", http.StatusBadRequest)
        return
    }
    circumference := FindCircumference(floatValue)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf(`{"circumference": %f}`, circumference)))
}
