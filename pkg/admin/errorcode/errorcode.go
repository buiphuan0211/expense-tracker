package errorcode

import "expense-tracker/internal/response"

// Init list error codes
// Common code internal: 1-99
// Common: 100-199
// Category: 200-299
// Staff: 300-399
// Staff: 400-499
func Init() {
	response.Init()
	response.AddListCodes(common)
	response.AddListCodes(category)
	response.AddListCodes(staff)
	response.AddListCodes(staffAuth)
}
