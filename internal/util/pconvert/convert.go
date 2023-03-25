package pconvert

import (
	"encoding/json"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

// ObjectIDsToStrings ...
func ObjectIDsToStrings(ids []primitive.ObjectID) []string {
	return funk.Map(ids, func(item primitive.ObjectID) string {
		return item.Hex()
	}).([]string)
}

// StringsToObjectIDs ...
func StringsToObjectIDs(strValues []string) []primitive.ObjectID {
	return funk.Map(strValues, func(item string) primitive.ObjectID {
		return StringToObjectID(item)
	}).([]primitive.ObjectID)
}

// StringToObjectID ...
func StringToObjectID(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

// ToString ...
func ToString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// StringToInt ...
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// IntToString ...
func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}
