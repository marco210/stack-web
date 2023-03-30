package database

import (
	"log"
	"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    // "github.com/joho/godotenv"
)

func DBConn() (db *sql.DB) {
    // dbDriver := "mysql"
    // dbUser := "root"
    // dbPass := "password"
    // dbName := "book_example"
    // dbHost := "0.0.0.0"
    // dbPort := "3306"
    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatalf("Error loading .env file: %s", err)
    // }
    dbDriver:=os.Getenv("DB_DRIVER")
    dbUser:=os.Getenv("DB_USER")
    dbPass:=os.Getenv("DB_PASS")
    dbName:=os.Getenv("DB_NAME")
    dbHost:=os.Getenv("DB_HOST")
    dbPort:=os.Getenv("DB_PORT")

    log.Println(dbUser,dbPass,dbHost,dbPort,dbName)

    url:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser,dbPass,dbHost,dbPort,dbName)
    fmt.Println(url)
    // db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
    if err != nil {
        panic(err.Error())
		// fmt.Println("Err", err.Error())
    }
	fmt.Println("Successful Connection to Database!")
    return db	
}