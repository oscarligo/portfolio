package createuser

import (
	"context"
	"fmt"
	"log"
	"os"

	"backend/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func create_user() {

	// Url from environment variable
	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("La variable DB_SOURCE no está definida")
	}

	// Username and password from environment variables
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	// Create connection pool
	if username == "" || password == "" {
		log.Fatal("ERROR: ADMIN_USERNAME and ADMIN_PASSWORD must be set in environment variables")
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("Error while creating connection pool:", err)
	}
	defer pool.Close()

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error while hashing password:", err)
	}

	// Initialize the sqlc repository and insert the record
	queries := repository.New(pool)

	user, err := queries.CreateUser(ctx, repository.CreateUserParams{
		Username:     username,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		log.Fatalf("Error while inserting the user into the DB (does the user '%s' already exist?): %v", username, err)
	}

	fmt.Printf("\nSuccess: the user '%s' has been created successfully in the database.\n\n", user.Username)
}
