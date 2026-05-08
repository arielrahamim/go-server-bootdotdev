package main

import (
	"chirpy/internal/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// encoding/json lib ignores lower case struct fields!
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

// Struct to define metrics data
type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
	platform       string
}

// middleware to increment the fileserver hits
// it accepts and returns an http.Handler interface
// so we create a nameless func that implements the interface, bump cfg (which is a pointer to apiconfig)
// and then continues to the next passed input
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

// MetricsHandler serves the metrics endpoint.
func (cfg *apiConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<body>
		<h1>Welcome, Chirpy Admin</h1>
		<p>Chirpy has been visited %d times!</p>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, html, cfg.fileserverHits.Load())
}

// Healthcheck handler (do x when someone hits the endpoint) for the healthcheck endpoint
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ResetMetricsHandler resets the metrics to 0 && delete all users
func (cfg *apiConfig) ResetMetricsHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return 
	}

	err := cfg.dbQueries.DeleteAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Error While Deleting All Users...", 500)
		return
	}
	

	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Metrics reset"))
}

func (cfg *apiConfig) UsersHandler(w http.ResponseWriter, r *http.Request) {
	type req struct {
		Email string `json:"email"`
	}

	payload := req{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	if payload.Email == "" {
		respondWithError(w, 400, "email is required")
		return
	}

	dbUser, err := cfg.dbQueries.CreateUser(r.Context(), payload.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	respUser := User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	}

	respondWithJson(w, http.StatusCreated, respUser)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type jsonErrorStruct struct {
		Body string `json:"error"`
	}

	respondWithJson(w, code, jsonErrorStruct{Body: msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func AsteriksReplacer() (asteriks string) {
	asteriks = "****"
	return
}

func WordsValidator(body string) (cleaned string) {
	bannedWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(body, " ")

	for i, word := range words {
		lower := strings.ToLower(word)
		for _, banned := range bannedWords {
			if lower == banned {
				words[i] = AsteriksReplacer()
			}
		}
	}

	cleaned = strings.Join(words, " ")
	return
}

func (cfg *apiConfig) ChirpsHandler(w http.ResponseWriter, r *http.Request) {
	type req struct {
		Body   string `json:"body"`
		UserID string `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	input := req{}
	err := decoder.Decode(&input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Length Validations
	if len(input.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	// Clean Chirp Body for Bad Words
	input.Body = WordsValidator(input.Body)

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid user_id")
		return
	}

	dbChirps, err := cfg.dbQueries.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   input.Body,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could Not Create Chirp")
		return
	}

	respChirp := Chirp{
		ID:        dbChirps.ID,
		CreatedAt: dbChirps.CreatedAt,
		UpdatedAt: dbChirps.UpdatedAt,
		Body:      dbChirps.Body,
		UserID:    dbChirps.UserID,
	}

	respondWithJson(w, 201, respChirp)
}

// go runtime runs main automatically
func main() {
	godotenv.Load()

	// DB Setup
	dbURL := os.Getenv("DB_URL")
	db, err1 := sql.Open("postgres", dbURL)
	if err1 != nil {
		log.Fatal(nil)
		return
	}

	// SQLC Setup
	dbQueries := database.New(db)

	// 'auth'
	platform := os.Getenv("PLATFORM")

	// one 'box' in memory to save the metrics data in
	// initialized in main
	cfg := &apiConfig{}

	// Update the apiConfig{} struct to hold db config
	cfg.dbQueries = dbQueries

	// users 
	cfg.platform = platform

	// create a new multiplexer (router)
	mux := http.NewServeMux()

	// declare endpoints
	// API
	mux.HandleFunc("GET /api/healthz", HealthzHandler)
	mux.HandleFunc("POST /api/chirps", cfg.ChirpsHandler)
	mux.HandleFunc("POST /api/users", cfg.UsersHandler)
	// Admin
	mux.HandleFunc("GET /admin/metrics", cfg.MetricsHandler)
	mux.HandleFunc("POST /admin/reset", cfg.ResetMetricsHandler)

	// var that hold handler config
	// serve static files from the process current dir
	// the /app/ stripping means remove the /app/ in the URL to match that local file
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))

	// serve the app directory, with our middleware + handler var
	mux.Handle("/app/", cfg.middlewareMetricsInc(handler))

	// assign the multiplexer and address to the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// listen and serve the server
	log.Println("listening on port", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
