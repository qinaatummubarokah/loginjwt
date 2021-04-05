package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

)

// JWTSecret
var JWTSecret string = "SL6ANV4cMfu2cBI240iV0xYLgv6RxUIh"

var Db *sqlx.DB

func Connect() *sqlx.DB {
	db := sqlx.MustConnect("mysql", "root:@tcp(127.0.0.1:3306)/login")
	log.Println(db)
    return db
}