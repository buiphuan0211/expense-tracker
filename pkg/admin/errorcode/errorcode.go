package errorcode

import "expense-tracker/internal/response"

// Init list error codes
// Common code internal: 1-99
// Common: 100-199
// Category: 200-299
func Init() {
	response.Init()
	response.AddListCodes(common)
	response.AddListCodes(category)
}
