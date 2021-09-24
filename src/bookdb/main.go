package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)
var server = "localhost"
var port = 1401
var user = "sa"
var password = "saP@ssw0rd"
var database = "bookdb"

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)

	log.Printf("connString =[" + connString + "]")

	db_sqlserver, err1 := sql.Open("mssql", connString)

	if err1 != nil {
		log.Fatal("Open connection failed:", err1.Error())

		panic(err1)
	}
	log.Printf("Sql Server Connected..!\n")

	checkVersion(db_sqlserver)

	defer db_sqlserver.Close()

	//fmt.Sprintf([username[:password]@][protocol[(address)]]/dbname)
	fmt.Println("------------")

	// db_mariadb, err2 := sql.Open("mysql", "root:root@/") // default --> tcp(localhost:3306)

	db_mariadb, err2 := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err2 != nil {
		os.Exit(1)
		panic(err2)
	}
	// See "Important settings" section.
	db_mariadb.SetConnMaxLifetime(time.Minute * 3)
	db_mariadb.SetMaxOpenConns(10)
	db_mariadb.SetMaxIdleConns(10)

	var version string
	db_mariadb.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	checkVersion(db_mariadb)
}

func checkVersion(db *sql.DB) {
	ctx := context.Background()

	//err := db.PingContext(ctx)
	var err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Ping database failed:", err.Error())
	}

	var version string
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&version)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	log.Printf("Version=%s\n", version)
}
