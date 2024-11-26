package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Account structure
type Account struct {
	AccID    int    `json:"acc_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Character structure
type Character struct {
	CharID  int `json:"char_id"`
	AccID   int `json:"acc_id"`
	ClassID int `json:"class_id"`
}

// Score structure
type Score struct {
	ScoreID     int `json:"score_id"`
	CharID      int `json:"char_id"`
	RewardScore int `json:"reward_score"`
}

// Connect to the database
func connectToDB() {
	var err error
	connStr := "postgres://postgres:55429298@localhost:5433/wira?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	fmt.Println("Successfully connected to the database")
}

// Get all accounts
func getAccounts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT acc_id, username, email FROM accounts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.AccID, &account.Username, &account.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accounts = append(accounts, account)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

// Get characters by account ID
func getCharactersByAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accID := vars["acc_id"]

	rows, err := db.Query("SELECT char_id, acc_id, class_id FROM characters WHERE acc_id = $1", accID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var characters []Character
	for rows.Next() {
		var character Character
		if err := rows.Scan(&character.CharID, &character.AccID, &character.ClassID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		characters = append(characters, character)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

// Get scores by character ID
func getScoresByCharacterID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	charID := vars["char_id"]

	rows, err := db.Query("SELECT score_id, char_id, reward_score FROM scores WHERE char_id = $1", charID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var scores []Score
	for rows.Next() {
		var score Score
		if err := rows.Scan(&score.ScoreID, &score.CharID, &score.RewardScore); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		scores = append(scores, score)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scores)
}

// Pagination for accounts
func getAccountsWithPagination(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // default value if not provided
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0 // default value if not provided
	}

	rows, err := db.Query("SELECT acc_id, username, email FROM accounts LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.AccID, &account.Username, &account.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accounts = append(accounts, account)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

// Initialize routes
func initializeRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", getAccounts).Methods("GET")
	r.HandleFunc("/accounts/{acc_id}/characters", getCharactersByAccountID).Methods("GET")
	r.HandleFunc("/characters/{char_id}/scores", getScoresByCharacterID).Methods("GET")
	r.HandleFunc("/accounts/paginate", getAccountsWithPagination).Methods("GET")
	http.Handle("/", r)
}

func main() {
	connectToDB()
	defer db.Close()

	initializeRoutes()
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
