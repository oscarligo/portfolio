package handler

import (
	"backend/internal/store"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Store     *store.Store
	JwtSecret string
}

func NewAuthHandler(store *store.Store, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		Store:     store,
		JwtSecret: jwtSecret,
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// POST /api/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	// Decode json body into loginRequest struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Username and password are not empty
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Search for the user in the database by username
	user, err := h.Store.GetUserByUsername(r.Context(), req.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Username or password is incorrect", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Username or password is incorrect", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(1 * time.Hour) // Expires in 1 hours

	claims := jwt.MapClaims{
		"sub": user.ID,
		"usr": user.Username,
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
	}

	// Create the token with the HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using your JWT_SECRET from the .env
	tokenString, err := token.SignedString([]byte(h.JwtSecret))
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	// Return the token and expiration time in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(loginResponse{
		Token:     tokenString,
		ExpiresAt: expirationTime,
	})
}
