package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Logger wraps handlers with logging information
func Logger(inner http.HandlerFunc, name string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// HealthCheckStatusResponse wraps a health check status response
type HealthCheckStatusResponse struct {
	Status string `json:"status"`
}

// HealthCheckStatusHandler handles health check status requests
func HealthCheckStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckStatusResponse{"OK"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", Logger(HealthCheckStatusHandler, "health#check"))

	log.Printf(fmt.Sprintf("Service listening on port %v", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil))
}
