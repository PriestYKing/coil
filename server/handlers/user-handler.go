package handlers

import (
	"coil/auth"
	"coil/model"
	"coil/rediscache"
	"coil/service"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func LoginHandler(db *sql.DB, ctx *context.Context, rdb *redis.Client, w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := service.LoginService(db, u.Email, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwtToken, err := auth.CreateToken(u.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message":  "Login Successful",
		"username": user.Username,
		"email":    user.Email,
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
	}

	// Storing token in cache for faster lookups in auth
	err = rediscache.StoreJWTToken(rdb, *ctx, user.Email, jwtToken, time.Duration(24)*time.Hour)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RegisterHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := service.RegisterService(db, u.Username, u.Email, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwtToken, err := auth.CreateToken(u.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message":  "Login Successful",
		"username": user.Username,
		"email":    user.Email,
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func CheckIsAuth(ctx *context.Context, rdb *redis.Client, w http.ResponseWriter, r *http.Request) {
	var email map[string]string
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	key := "auth:token" + email["email"]

	val, err := rdb.Get(*ctx, key).Result()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if val == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	response := map[string]bool{
		"authenticated": true,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
