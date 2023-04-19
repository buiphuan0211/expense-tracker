package auth

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/config/plogger"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kr/pretty"
)

type User struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// GetCurrentUserByToken ...
func GetCurrentUserByToken(tokenString string) *User {
	if tokenString == "" {
		return nil
	}

	data, err := getDataClaims(tokenString)
	if err != nil {
		return nil
	}

	var user = new(User)

	id := data["_id"].(string)
	if id != "" {
		user.ID = id
	}

	name := data["name"].(string)
	if name != "" {
		user.Name = name
	}

	return user
}

// getDataClaims ...
func getDataClaims(tokenString string) (data map[string]interface{}, err error) {
	var envVars = config.GetENV()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(envVars.Auth.SecretKey), nil
	})
	if err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "internal.auth - GetUserByToken - jwt.Parse",
			Message: err.Error(),
			Data: plogger.Map{
				"tokenString": tokenString,
			},
		})
	}

	claims := token.Claims.(jwt.MapClaims)

	data = claims["data"].(map[string]interface{})
	return
}

// GetCurrentUserByTokenError ...
func GetCurrentUserByTokenError(token interface{}) *User {
	if token == nil {
		return nil
	}

	pretty.Println(token)
	data, ok := token.(*jwt.Token)
	if !ok {
		fmt.Println("error 1")
		return nil
	}

	fmt.Println("data: ", data)

	m, ok := data.Claims.(jwt.MapClaims)
	if !ok || !data.Valid || m == nil {
		fmt.Println("error 2")

		return nil
	}
	var (
		user = new(User)
	)

	fmt.Println("m: ", m)

	fmt.Println("m[\"_id\"]: ", m["_id"])
	fmt.Println("m[\"_name\"]: ", m["name"])

	if m["_id"] != "" {
		fmt.Println("error 3")
		user.ID = m["_id"].(string)
	}
	if m["name"] != "" {
		fmt.Println("error 4")
		name, ok := m["name"]
		if ok {
			user.Name = name.(string)
		}
		fmt.Println("error 5")
	}

	return user
}
