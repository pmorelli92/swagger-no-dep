package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	// @version 1.0.0
	// @title Pet Service
	// @description Handles pet information
	// @host localhost
	// @schemes http
	router := http.NewServeMux()
	router.HandleFunc("/api/pets/", getPetByID())
	router.HandleFunc("/api/pets", getPets())
	router.HandleFunc("/api/docs", docs())

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

type petType string

const (
	Dog = petType("Dog")
	Cat = petType("Cat")
)

type petResponse struct {
	Id   int     `json:"id" example:"1"`
	Name string  `json:"name" example:"Fenrir"`
	Type petType `json:"type" example:"dog" enums:"dog,cat"`
}

type errorResponse struct {
	Message string `json:"message"`
}

var pets = []petResponse{
	{
		Id:   1,
		Name: "Neon",
		Type: Dog,
	},
	{
		Id:   2,
		Name: "Bills",
		Type: Cat,
	},
}

// @summary Get pets
// @description Gets a list of pets
// @id get-pets
// @produce json
// @Success 200 {object} []main.petResponse
// @Success 405 {object} main.errorResponse
// @Router /api/pets/ [get]
// @tags Pets
func getPets() http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if rq.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(errorResponse{Message: "method not allowed"})
			return
		}

		json.NewEncoder(w).Encode(pets)
	}
}

// @summary Get pet by ID
// @description Gets a pet using the pet ID
// @id get-pet-by-id
// @produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} main.petResponse
// @Success 400 {object} main.errorResponse
// @Success 404 {object} main.errorResponse
// @Success 405 {object} main.errorResponse
// @Router /api/pets/{id} [get]
// @tags Pets
func getPetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if rq.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(errorResponse{Message: "method not allowed"})
			return
		}

		p := strings.Split(rq.URL.Path, "/")
		idStr := p[len(p)-1]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Message: "the path ID is not a valid int"})
			return
		}

		for _, p := range pets {
			if p.Id == id {
				json.NewEncoder(w).Encode(p)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse{Message: "pet not found"})
	}
}

//go:embed swagger/swagger.html
var swaggerHTML string

func docs() http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprint(w, swaggerHTML)
	}
}
