package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"rss/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


type apiConfig struct{
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found")
	}

	db_url := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("DB not found")
	}

	db, err := sql.Open("postgres",db_url)
	if err!=nil{
		log.Fatal("Cant connect to the database",err)
	}

    queries := database.New(db)

	if err!=nil{
		log.Fatal("Cant create to the database conn",err)
	}

	apiCfg := apiConfig{
		DB:queries,
	}

	fmt.Println("Starting server at", port)

	router := chi.NewRouter()

	// router.Mount(v1Router)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users",apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)




	http.ListenAndServe(":"+port, router)

}
