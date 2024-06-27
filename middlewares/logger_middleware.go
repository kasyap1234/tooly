package middlewares

import (
	"log"
	"net/http"
	"time"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)
func LoggerMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
        
		start :=time.Now(); 
		
		 ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
        next.ServeHTTP(ww, r)

        duration := time.Since(start)

        log.Printf("Completed %s %s %d %s in %v", r.Method, r.RequestURI, ww.Status(), http.StatusText(ww.Status()), duration)

		
	})
}