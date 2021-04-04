package models

import (
	"log"
	"net/http"
	"loginjwt/config"

	_ "github.com/go-sql-driver/mysql"

)

type User struct {
    ID			int			`db:"id"		json:"id"`
    Name		string		`db:"name"		json:"name"		validate:"required"`
    Password	string		`db:"password"	json:"password"	validate:"required"`
    Email		string		`db:"email"		json:"email"`
}

type Token struct{
	Token		string		`json:"token"`
}

type Profile struct{
	Name		string		`json:"name"`
    Email		string		`json:"email"`
}

// models GetUser
func GetUser(c *User, id string) int {
	query := `SELECT id, name, password, email FROM user WHERE name = ?`

	err := config.Db.Get(c, query, id)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound
	}
	return http.StatusOK
}