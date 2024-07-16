package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/sir-george2500/g-server/docs"
	"github.com/sir-george2500/g-server/internal/database"
	httpSwagger "github.com/swaggo/http-swagger"
)

type apiConfig struct {
	DB *database.Queries
}

// @title g-server API
// @version 1.0
// @description g-server - a full-blown RSS feed aggregator.
// @host localhost:8080
// @BasePath /v1
// @contact.name George S Mulbah
// @contact.url https://github.com/sir-george2500
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set in the environment")
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set in the environment")
	}

	// Connect to the database
	conn, dberr := sql.Open("postgres", dbURL)
	if dberr != nil {
		log.Fatal("Failed to connect to the database:", dberr)
	}
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	// Create a new Router
	router := chi.NewRouter()

	// Handle CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// API routes
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", apiCfg.handleGetFeeds)
	v1Router.Get("/", helloWorld)
	v1Router.Get("/err", handleErr)
	v1Router.Get("/name", apiCfg.handleGetFeeds)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeeds))
	v1Router.Get("/feeds", apiCfg.handleGetFeeds)
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleCreateFeedsFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleGetFeedsFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlGetPostForUser))

	// Serve Swagger UI
	swaggerURL := "http://" + os.Getenv("HOST") + ":" + portString + "/swagger/doc.json"
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))

	// Mount the router
	router.Mount("/v1", v1Router)

	// Start the server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

