package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"employeejwt/internal/auth"
	"employeejwt/internal/constants"
	"employeejwt/internal/database"
	"employeejwt/internal/models"
	"employeejwt/internal/repository"
	"employeejwt/internal/validation"

	"golang.org/x/crypto/bcrypt"
)

var userStore *repository.UserStore

func getUserStore() *repository.UserStore {
	if userStore == nil {
		userStore = repository.NewUserStore(database.DB)
	}
	return userStore
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("register: invalid request body for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, constants.ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}
	if err := validation.ValidateRegisterRequest(req); err != nil {
		log.Printf("register: missing required fields for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("register: could not hash password for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "could not hash password", http.StatusInternalServerError)
		return
	}
	log.Println("User registered successfully")
	user := models.User{Name: req.Name, Email: req.Email, Password: string(hash)}
	if err := getUserStore().CreateUser(&user); err != nil {
		log.Printf("register: failed to create user for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
	   resp, err := auth.GenerateToken(user.ID, user.Email)

	   	if err != nil {
	   		http.Error(w, "could not generate token", http.StatusInternalServerError)
	   		return
	   	}

	   w.Header().Set("Content-Type", "application/json")
	   _ = json.NewEncoder(w).Encode(models.AuthResponse{Token: resp})
	*/
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("login: invalid request body for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, constants.ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}
	if err := validation.ValidateLoginRequest(req); err != nil {
		log.Printf("login: missing credentials for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := getUserStore().GetByEmail(req.Email)
	if err != nil {
		log.Printf("login: invalid credentials lookup for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Printf("login: password mismatch for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("login: could not generate token for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}
	log.Println("User logged in successfully")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(models.AuthResponse{Token: token})
}
