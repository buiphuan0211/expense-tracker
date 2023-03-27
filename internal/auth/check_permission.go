package auth

import (
	"github.com/thoas/go-funk"
)

// CheckPermission ...
func CheckPermission(permissions []string, tokenString string) (result bool) {
	if tokenString == "" {
		return
	}

	data, err := getDataClaims(tokenString)
	if err != nil {
		return
	}
	
	if isRoot := data["isRoot"].(bool); isRoot {
		return true
	}

	permissionsInToken := data["permissions"].([]interface{})
	if len(permissions) < 0 {
		return
	}

	for _, p := range permissions {
		if !funk.Contains(permissionsInToken, p) {
			return
		}
	}

	return true
}
