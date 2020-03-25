package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/naaltunian/go-mongo/models"
	"github.com/naaltunian/go-mongo/utils"
)

type server struct {
	router *mux.Router
	srv    *http.Server
}

func main() {
	s := server{
		router: mux.NewRouter(),
	}
	s.routes()

	port := os.Getenv("CONFIG_HTTP_PORT")

	utils.ConnectToDatabase()

	log.Println("Using Port " + port)
	s.srv = &http.Server{
		Addr:           ":" + port,
		Handler:        s.router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func (s *server) routes() {
	// checks server status
	s.router.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
	}).Methods("GET")

	getRouter := s.router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/get_users", models.GetUsers)
	getRouter.HandleFunc("/get_user/{id}", models.GetUser)

	postRouter := s.router.Methods("POST", "OPTIONS").Subrouter()
	postRouter.HandleFunc("/create_user", models.CreateUser)
	postRouter.Use(models.ValidateUserMiddleware)
}
