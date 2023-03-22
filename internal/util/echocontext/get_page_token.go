package echocontext

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

// PageToken ...
type PageToken struct {
	Page      int64
	Timestamp time.Time
	Order     int
}

func getDefaultPageToken() (response PageToken) {
	response.Page = 0
	response.Timestamp = time.Now()
	return response
}

// PageTokenDecode decode page token from query
func PageTokenDecode(pk string) PageToken {
	encodedString := pk
	if encodedString == "" {
		return getDefaultPageToken()
	}

	// Decode string
	decoded, err := base64.StdEncoding.DecodeString(encodedString)

	if err != nil {
		return getDefaultPageToken()
	}

	// Parse token
	var pageToken PageToken
	err = json.Unmarshal(decoded, &pageToken)
	if err != nil {
		return getDefaultPageToken()
	}

	return pageToken
}
