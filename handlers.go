package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"html"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Gyaneshwar pardhi, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := Todos{
		Todo{Name:"Call me"},
		Todo{Name: "Love me"},
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err);
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	todoId := vars["todoId"]
	fmt.Fprint(w, "To do show:", todoId);
}
