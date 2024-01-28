package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gargdaya/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// DB Connection and Queries config

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	conn, error := sql.Open("postgres", dbURL)

	if error != nil {
		log.Fatal(error)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	go startScraping(apiCfg.DB, 10, time.Minute*5)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)

	v1Router.Get("/err", handlerErr)

	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.authMiddleware(apiCfg.handlerGetUser))

	v1Router.Post("/feeds", apiCfg.authMiddleware(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)

	v1Router.Get("/posts", apiCfg.authMiddleware(apiCfg.handlerGetPostsForUser))

	v1Router.Post("/feed-follows", apiCfg.authMiddleware(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed-follows", apiCfg.authMiddleware(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed-follows/{feedFollowId}", apiCfg.authMiddleware(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Server running on http://localhost:" + port)
	err := server.ListenAndServe()

	if err != nil {
		println(err.Error())
	}
}
