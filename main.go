package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sir-george2500/g-server/internal/database"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB URL is not found in the env")
	}
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port not found in the env")
	}

	// conect to the database
	conn, dberr := sql.Open("postgres", dbURL)

	if dberr != nil {
		log.Fatal("fail to connect to the database", dberr)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// create a new Router
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	// hanlde cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", apiCfg.handleGetFeeds)
	v1Router.Get("/err", handleErr)
	v1Router.Get("/name", apiCfg.handleGetFeeds)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeeds))
	v1Router.Get("/feeds", apiCfg.handleGetFeeds)
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleCreateFeedsFollow))

	//Mount the router
	router.Mount("/v1", v1Router)
	// start the server
	log.Printf("Serve Staring on Port %v", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
