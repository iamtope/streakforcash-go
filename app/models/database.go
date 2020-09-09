package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var db *sql.DB

// initialise the databse 
func init() {
	e := godotenv.Load() // load .env file, ideally this should be done in the func main
	if e != nil {
		fmt.Print(e);
	}


DBHOST := os.Getenv("DBHOST")
	DBUSER := os.Getenv("DBUSER")
	DBPASSWORD := os.Getenv("DBPASSWORD")
	DBNAME := os.Getenv("DBNAME")
	DBPORT := os.Getenv("DBPORT")
	sslmode := os.Getenv("SSLMODE")

	dbURL := fmt.Sprintf("DBHOST=%s DBUSER=%s DBPASSWORD=%s DBNAME=%s DBPORT=%s sslmode=%s", DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, sslmode)
	fmt.Println(dbURL)

	db, err := sql.Open("postgres", dbURL)
	
	fmt.Println(err)
	if err != nil {
	panic(err)
	// defer db.Close()
	} 
	defer db.Close()
	// else {
	// 	fmt.Println("We are connected to the postgresql database", dbURL)
	// }

}