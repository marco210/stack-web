package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "password"
    dbName := "book_example"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        // panic(err.Error())
		fmt.Println("Err", err.Error())
    }
	fmt.Println("Successful Connection to Database!")
    return db	
}