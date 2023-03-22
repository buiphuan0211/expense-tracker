package initialize

import (
	"expense-tracker/internal/config/database"
)

func mongodb() {
	database.ConnectMongoDBExpenseTracker()
}
