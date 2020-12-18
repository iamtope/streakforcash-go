package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

// initialise the databse 
func ReturnInstance() *sql.DB{
	e := godotenv.Load(".env") // load .env file, ideally this should be done in the func main
	if e != nil {
		fmt.Print(e);
	}

	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	sslmode := os.Getenv("SSLMODE")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s  sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", dbURL)
	
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
	  panic(err)
	}
  
	fmt.Println("Successfully connected!")
	fmt.Println("You are Successfully connected!")
	return db

}