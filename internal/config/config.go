package config

import (
	"expense-tracker/internal/constant"
	"os"
)

type ENV struct {
	Env string

	ZookeeperPrefixExpenseTrackerCommon string
	ZookeeperPrefixExpenseTrackerAdmin  string

	Admin struct {
		Server string
		Port   string
	}

	MongoDB struct {
		URI    string
		DBName string

		ReplicaSet          string
		CAPem               string
		CertPem             string
		CertKeyFilePassword string
		ReadPrefMode        string
	}
}

var env ENV

func GetENV() *ENV {
	return &env
}

func IsEnvDevelop() bool {
	return env.Env == constant.EnvDevelop
}
func IsEnvProduction() bool {
	return env.Env == constant.EnvProduction
}

func Init() {
	env = ENV{
		Env:                                 os.Getenv("ENV"),
		ZookeeperPrefixExpenseTrackerCommon: os.Getenv("ZOOKEEPER_PREFIX_EXPENSE_TRACKER_COMMON"),
		ZookeeperPrefixExpenseTrackerAdmin:  os.Getenv("ZOOKEEPER_PREFIX_EXPENSE_TRACKER_ADMIN"),
	}
}
