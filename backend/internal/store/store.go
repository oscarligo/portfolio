package store

import (
	"backend/internal/repository"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// Store struct encapsulates the database connection pool and the generated queries from sqlc, providing a unified interface for database operations.
type Store struct {
	*repository.Queries
	db *pgxpool.Pool
}

// NewStore initializes a new Store with the given database connection pool and returns a pointer to it.
func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		Queries: repository.New(pool),
		db:      pool,
	}
}

// Struct for encapsulating all parameters needed to create a project with its related images and
// technologies in a single transaction
type CreateProjectTxParams struct {
	repository.CreateProjectParams
	ImageURLs     []string `json:"images"`
	TechnologyIDs []int32  `json:"technologies"`
}

// CreateProjectTx executes a series of database operations to create a project, insert its associated
// images, and link it with technologies, all within a single transaction to ensure data integrity.
func (store *Store) CreateProjectTx(ctx context.Context, arg CreateProjectTxParams) (repository.Project, error) {
	var result repository.Project

	// Start
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return result, err
	}
	defer tx.Rollback(ctx)

	qtx := repository.New(tx)

	// Step 1: Create the base project record and capture its ID for subsequent operations
	result, err = qtx.CreateProject(ctx, arg.CreateProjectParams)
	if err != nil {
		return result, fmt.Errorf("error al crear el proyecto base: %w", err)
	}

	// Step 2: Insert the image URLs by iterating over the slice
	for _, url := range arg.ImageURLs {
		_, err = qtx.CreateProjectImage(ctx, repository.CreateProjectImageParams{
			ProjectID: result.ID,
			ImageUrl:  url,
		})
		if err != nil {
			return result, fmt.Errorf("error al insertar la imagen %s: %w", url, err)
		}
	}

	// Step 3: Associate the technologies in the many-to-many relationship table
	for _, techID := range arg.TechnologyIDs {
		err = qtx.AssociateProjectTechnology(ctx, repository.AssociateProjectTechnologyParams{
			ProjectID:    result.ID,
			TechnologyID: techID,
		})
		if err != nil {
			return result, fmt.Errorf("error al asociar la tecnologia con ID %d: %w", techID, err)
		}
	}

	// If all operations completed successfully without errors, commit the changes definitively to PostgreSQL
	err = tx.Commit(ctx)
	return result, err
}

func (store *Store) CreateAdminUserIfNotExist(ctx context.Context, username, password string) error {
	// Hasshing the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Error while hashing password: %w", err)
	}

	// Attempt to insert the admin user into the database.
	_, err = store.Queries.CreateUser(ctx, repository.CreateUserParams{
		Username:     username,
		PasswordHash: string(hashedPassword),
	})

	if err != nil {

		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return nil
		}
		return fmt.Errorf("Error while inserting the admin user: %w", err)
	}

	return nil
}
