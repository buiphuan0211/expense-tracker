package pzk

import (
	"expense-tracker/internal/config"
	"fmt"
	"os"
)

// Zookeeper ...
func Zookeeper() {
	uri := os.Getenv("ZOOKEEPER_URI")
	if err := Connect(uri); err != nil {
		panic(err)
	}

	envVars := config.GetENV()
	server(envVars)
	commonValues(envVars)
}

// server ...
func server(envVars *config.ENV) {
	// Admin
	adminPrefix := getAdminPrefix("")
	envVars.Admin.Server = GetStringValue(fmt.Sprintf("%s/server_port/server", adminPrefix))
	envVars.Admin.Port = GetStringValue(fmt.Sprintf("%s/server_port/port", adminPrefix))

	// Mongodb
	commonPrefix := getCommonPrefix("")
	envVars.MongoDB.URI = GetStringValue(fmt.Sprintf("%s/mongodb/uri", commonPrefix))
	envVars.MongoDB.DBName = GetStringValue(fmt.Sprintf("%s/mongodb/db_name", commonPrefix))

}

// commonValues ...
func commonValues(envVars *config.ENV) {}

// getAdminPrefix ...
func getAdminPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixExpenseTrackerAdmin, group)
}

// getCommonPrefix ...
func getCommonPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixExpenseTrackerCommon, group)
}
