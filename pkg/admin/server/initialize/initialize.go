package initialize

import (
	"expense-tracker/internal/config"
	"expense-tracker/pkg/admin/errorcode"
)

// Init ...
func Init() {
	// Config ...
	config.Init()

	// ErrorCode ...
	errorcode.Init()

	// zookeeper ...
	zookeeper()

	// Mongodb ...
	mongodb()
}
