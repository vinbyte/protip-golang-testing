package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vinbyte/protip-golang-testing/app/helpers"
	userHandler "github.com/vinbyte/protip-golang-testing/user/delivery/http"
	userRepo "github.com/vinbyte/protip-golang-testing/user/repository/postgres"
	userUsecase "github.com/vinbyte/protip-golang-testing/user/usecase"
)

func main() {
	//read .env config
	_ = godotenv.Load()

	//setup timeout
	timeoutStr := os.Getenv("TIMEOUT")
	if timeoutStr == "" {
		timeoutStr = "3"
	}
	timeout, _ := strconv.Atoi(timeoutStr)
	timeoutSecond := time.Duration(timeout) * time.Second

	//db connection
	dbConn := helpers.InitPostgres()

	//setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	//setup layer
	ur := userRepo.NewUserRepository(dbConn)
	uu := userUsecase.NewUserUsecase(ur, timeoutSecond)
	userHandler.NewUserHandler(r, uu)

	//run the app
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.ListenAndServe(":"+port, r)
}
