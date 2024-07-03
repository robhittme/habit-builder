package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	c "habit-builder/cmds"
)

type Response struct {
	Message string `json:"message"`
}

func InitRoutes() {
	r := http.NewServeMux()
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/habits", HabitHandler)
	fmt.Printf("Server listening on port 4444\n")
	log.Fatal(http.ListenAndServe(":4444", r))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "service healthy"}
	json.NewEncoder(w).Encode(response)
}

func HabitHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		habits, err := c.LoadHabits()
		if err != nil {
			fmt.Println(err, "err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(habits)
		if err != nil {
			fmt.Println(err, "err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
