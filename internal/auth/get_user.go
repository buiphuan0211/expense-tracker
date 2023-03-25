package auth

import "github.com/golang-jwt/jwt"

type User struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// GetCurrentUserByToken ...
func GetCurrentUserByToken(token interface{}) *User {
	if token == nil {
		return nil
	}

	data, ok := token.(*jwt.Token)
	if !ok {
		return nil
	}
	m, ok := data.Claims.(jwt.MapClaims)
	if !ok || !data.Valid || m == nil {
		return nil
	}
	var (
		user = new(User)
	)

	if m["_id"] != "" {
		user.ID = m["_id"].(string)
	}
	if m["name"] != "" {
		name, ok := m["name"]
		if ok {
			user.Name = name.(string)

		}
	}

	return user
}
