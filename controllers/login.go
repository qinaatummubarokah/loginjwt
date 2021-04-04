package controllers

import(
	"loginjwt/models"
	"loginjwt/config"
	"github.com/labstack/echo"
	"log"
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

// GetToken function
func GetToken(c echo.Context) error {
	jsonmap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonmap)
	if err != nil {
		return err
	}

	username := jsonmap["name"]
	password := jsonmap["password"]

	var user models.User
	status := models.GetUser(&user, username.(string))
	if status == http.StatusOK {
		log.Println("USER: ",user)
		if user.Password == password.(string){
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id"   		: user.ID,
				"name" 		: user.Name,
				"email" 	: user.Email,
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte(config.JWTSecret))
			if err != nil {
				log.Println(err)
				return echo.NewHTTPError(status, "error SignedString")
			}
			return  c.JSON(status, tokenString)
		}else{
			return echo.NewHTTPError(status, "Password tidak sesuai")
		}
	}else{
		return echo.NewHTTPError(status, "Account not found")
	}
}

// GetProfile function
func GetProfile(c echo.Context) error {
	request := c.Request()

	authorizationHeader := request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return echo.NewHTTPError(http.StatusBadRequest, "err")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		log.Println("err: ",err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}

	claims , ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}
	
	var profile models.Profile
	profile.Name = claims["name"].(string)
	profile.Email = claims["email"].(string)

	return echo.NewHTTPError(http.StatusBadRequest, profile)
}