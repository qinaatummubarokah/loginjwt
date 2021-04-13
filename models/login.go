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

func CreateUser(params map[string]string) int {
	query := "INSERT INTO user("
	// Get params
	var fields, values string
	i := 0
	for key, value := range params {
		fields += "`" + key + "`"
		values += "'" + value + "'"
		if (len(params) - 1) > i {
			fields += ", "
			values += ", "
		}
		i++
	}
	// Combile params to build query
	query += fields + ") VALUES(" + values + ")"
	log.Println(query)

	tx, err := config.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusBadGateway
	}
	_, err = tx.Exec(query)
	tx.Commit()
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func UpdateUser(params map[string]string) int {
	query := "UPDATE user SET "
	// Get params
	i := 0
	for key, value := range params {
		if key != "id" {
			query += "`" + key + "`" + " = '" + value + "'"
			if (len(params) - 2) > i {
				query += ", "
			}
			i++
		}
	}
	query += " WHERE id = '" + params["id"] + "'"
	log.Println(query)

	tx, err := config.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusBadGateway
	}
	_, err = tx.Exec(query)
	tx.Commit()
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	return http.StatusOK
}