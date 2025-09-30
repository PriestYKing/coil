package main

import (
	"coil/auth"
	"coil/db"
	"coil/handlers"
	"coil/rediscache"
	"context"
	"log"
	"net/http"
)

func main() {

	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Print("Connected to DB")

	ctx := context.Background()
	rdb := rediscache.InitRedis()
	defer rdb.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(db, &ctx, rdb, w, r)
	})
	mux.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(db, w, r)
	})

	mux.HandleFunc("POST /isAuth", func(w http.ResponseWriter, r *http.Request) {
		handlers.CheckIsAuth(&ctx, rdb, w, r)
	})

	println("Server is running on port 8080")
	http.ListenAndServe(":8080", auth.GlobalMiddleware(mux))

}
