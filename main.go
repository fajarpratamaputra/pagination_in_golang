package main

import (
	"fmt"
	"general/config"
	"general/routes"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	// env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error loading db")
		panic(err)
	}
	config.DB = conn

	// Get generic database object sql.DB to use its functions
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatal("Error loading db")
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(25)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}

func main() {
	fmt.Println("-------- Run --------")
	r := routes.SetupRouter()
	//running
	r.Run(os.Getenv("PORT"))
}
