package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cetRide/api-rideyu/api/handlers"
	r "github.com/cetRide/api-rideyu/api/routes"
	"github.com/cetRide/api-rideyu/infrastructure/repository"
	service "github.com/cetRide/api-rideyu/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	port := os.Getenv("PORT")
	dbUrI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, username, dbName, password)

	dba := repository.ConnectDB(dbUrI)

	conn := repository.NewRepo(dba)

	repo := service.NewRepoHandler(conn)

	h := handlers.NewUseCaseHandler(repo)

	router := r.NewRouter(h)
	
	if port == "" {
		port = "4747"
	}
	log.Fatal(router.Run(":" + port))

}
