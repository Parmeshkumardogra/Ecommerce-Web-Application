package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMySQLDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:Ril@12345@tcp(localhost:3306)/")
	if err != nil {
		log.Println("error during creating sql open connection: ", err)
		return err
	}
	DB.SetConnMaxLifetime(time.Minute * 5)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = DB.PingContext(ctx)
	if err != nil {
		log.Println("error while DB ping", err)
		return err
	}
	_, err = DB.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS bankDB")
	if err != nil {
		log.Println("error creating database:", err)
		return err
	}
	// log.Println("Database created or already exists.")

	// Now you can connect to the specific database
	_, err = DB.ExecContext(ctx, "USE bankDB")
	if err != nil {
		log.Println("error selecting database:", err)
		return err
	}
	log.Println("Connected to bankDB in MySQL!")
	return nil

}

func DisconnectMySQLDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error disconnecting from MySQL: %v", err)
		} else {
			log.Println("Disconnected from MySQL!")
		}
	} else {
		log.Println("MySQL DB is nil, cannot disconnect")
	}
}
