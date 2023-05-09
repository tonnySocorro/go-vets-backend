package handlers

import (
	"log"
	"net/http"
	"os"
	"project/go-vets-backend/middlewares"
	"project/go-vets-backend/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/signUp", middlewares.CheckDB(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
