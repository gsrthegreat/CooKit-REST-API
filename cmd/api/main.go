package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gsrthegreat/CookIt/internal/auth"
	"github.com/gsrthegreat/CookIt/internal/database"
	"github.com/gsrthegreat/CookIt/internal/handlers"
	"github.com/gsrthegreat/CookIt/internal/middleware"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	fmt.Println("DB connected succesfully!")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-change-in-production"
	}
	auth.InitJWT(jwtSecret)

	authHandler := handlers.NewAuthHandler(db, auth.GenerateToken, auth.ValidateToken)
	pageHandler := handlers.NewPageHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/login", authHandler.LoginHandler)
	mux.HandleFunc("/api/v1/logout", authHandler.LogoutHandler)
	mux.HandleFunc("/api/v1/register", authHandler.RegisterHandler)
	mux.HandleFunc("/api/v1/", authHandler.HomeHandler)

	mux.HandleFunc("/login", pageHandler.LoginPageHandler)
	mux.HandleFunc("/register", pageHandler.RegisterPageHandler)
	mux.HandleFunc("/", pageHandler.HomePageHandler)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	handler := middleware.CORS(mux)

	fmt.Println("server listening at port :8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
