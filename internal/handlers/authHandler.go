package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB            *sql.DB
	GenerateToken func(username string) (string, error)
	ValidateToken func(tokenString string) (string, error)
}

func NewAuthHandler(db *sql.DB, generateToken func(string) (string, error), validateToken func(string) (string, error)) *AuthHandler {
	return &AuthHandler{
		DB:            db,
		GenerateToken: generateToken,
		ValidateToken: validateToken,
	}
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HomeHandler handles the homepage endpoint
func (h *AuthHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		w.Write([]byte("Welcome Guest"))
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		w.Write([]byte("Welcome Guest"))
		return
	}

	username, err := h.ValidateToken(tokenString)
	if err != nil {
		w.Write([]byte("Welcome Guest"))
		return
	}

	var dbUsername string
	err = h.DB.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&dbUsername)
	if err != nil {
		w.Write([]byte("Welcome Guest"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s", dbUsername)))
}

// LoginHandler handles user login
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	var dbPassword string
	err := h.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&dbPassword)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !CheckPasswordHash(password, dbPassword) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.GenerateToken(username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"token":"%s"}`, token)))
}

// LogoutHandler handles user logout
func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Logged out successfully. Please remove the token from client."}`))
}

// RegisterHandler handles user registration
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	_, err = h.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"User registered successfully"}`))
}
