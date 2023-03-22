package pzk

import "fmt"

// GetStringValue ...
func GetStringValue(path string) string {
	value, _, err := zkClient.Get(path)
	if err != nil {
		fmt.Printf("Get value from path %s err:%v\n", path, err)
	}

	return string(value)
}
