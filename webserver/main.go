package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
	"webserver/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("PORT is not defined in the .env file")
	}
	if dbURL == "" {
		log.Fatal("DB_URL is not defined in the .env file")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Error with connection to the db", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	fmt.Println(apiCfg)

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 2)
		respondWithJSON(w, 200, "Hello World")
	})
	v1Router.Get("/ready", handleReady)
	v1Router.Get("/error", handleError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("Error with starting the server", err)
		return
	}

}
