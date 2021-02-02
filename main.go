package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cetRide/api-rideyu/api/controllers"
	r "github.com/cetRide/api-rideyu/api/routes"
	"github.com/cetRide/api-rideyu/infrastructure/repository"
	"github.com/cetRide/api-rideyu/usecase"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rideyu"
	password = "password"
	dbname   = "pexs"
)

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Print("error1")
		fmt.Print(err)
		panic(err)
	}


	err = db.Ping()
	if err != nil {
		fmt.Print("error2")
		fmt.Print(err)
		panic(err)
	}
	return db
}

func main() {
	dba := connectDB()

	conn := repository.NewRepo(dba)

	repo := usecase.NewRepoHandler(conn)

	h := controllers.NewUseCaseHandler(repo)

	router := r.NewRouter(h)

	err := http.ListenAndServe(":5004",
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		panic(err)
	}
}
