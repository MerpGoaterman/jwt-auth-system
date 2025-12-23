package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/MerpGoaterman/jwt-auth-system/backend/auth"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB(database *sql.DB) {
	db = database
}

// User represents a user in the system
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	TenantID     string    `json:"tenant_id"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents login response with token
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CreateUserRequest represents user creation request
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
}

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash the password
	passwordHash := auth.HashPassword(req.Password)

	// Query user from database
	var user User
	err := db.QueryRow(`
		SELECT id, name, email, password_hash, tenant_id, role, created_at, updated_at
		FROM users WHERE email = ? AND password_hash = ?
	`, req.Email, passwordHash).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash,
		&user.TenantID, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.ID, user.Email, user.TenantID, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return token and user info
	response := LoginResponse{
		Token: token,
		User:  user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT id, name, email, tenant_id, role, created_at, updated_at
		FROM users
	`)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.Name, &user.Email,
			&user.TenantID, &user.Role, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Check if user is admin
	claims, ok := auth.GetUserClaims(r.Context())
	if !ok || claims.Role != "admin" {
		http.Error(w, "Admin access required", http.StatusForbidden)
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Password == "" || req.TenantID == "" || req.Role == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Generate user ID
	userID := uuid.New().String()

	// Hash password
	passwordHash := auth.HashPassword(req.Password)

	// Insert user into database
	_, err := db.Exec(`
		INSERT INTO users (id, name, email, password_hash, tenant_id, role)
		VALUES (?, ?, ?, ?, ?, ?)
	`, userID, req.Name, req.Email, passwordHash, req.TenantID, req.Role)

	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return created user
	user := User{
		ID:       userID,
		Name:     req.Name,
		Email:    req.Email,
		TenantID: req.TenantID,
		Role:     req.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser returns a single user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user User
	err := db.QueryRow(`
		SELECT id, name, email, tenant_id, role, created_at, updated_at
		FROM users WHERE id = ?
	`, userID).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.TenantID, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update user in database
	_, err := db.Exec(`
		UPDATE users SET name = ?, email = ?, tenant_id = ?, role = ?
		WHERE id = ?
	`, req.Name, req.Email, req.TenantID, req.Role, userID)

	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Return updated user
	user := User{
		ID:       userID,
		Name:     req.Name,
		Email:    req.Email,
		TenantID: req.TenantID,
		Role:     req.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Check if user is admin
	claims, ok := auth.GetUserClaims(r.Context())
	if !ok || claims.Role != "admin" {
		http.Error(w, "Admin access required", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]

	// Delete user from database
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetCurrentUser returns the current authenticated user
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetUserClaims(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var user User
	err := db.QueryRow(`
		SELECT id, name, email, tenant_id, role, created_at, updated_at
		FROM users WHERE id = ?
	`, claims.UserID).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.TenantID, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
