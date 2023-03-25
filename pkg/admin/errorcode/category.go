package errorcode

import "expense-tracker/internal/response"

const (
	CategoryRequiredName = "category_required_name"
	CategoryNotFound     = "category_not_found"
	CategoryExistedName  = "category_exists_name"
	CategoryInvalid      = "category_invalid"
)

var category = []response.Code{
	{
		Key:     CategoryRequiredName,
		Message: "tên nhóm chi tiêu không được trống",
		Code:    200,
	},
	{
		Key:     CategoryNotFound,
		Message: "nhóm chi tiêu không tìm thấy",
		Code:    201,
	},
	{
		Key:     CategoryExistedName,
		Message: "tên nhóm chi tiêu đã tồn tại",
		Code:    202,
	},
	{
		Key:     CategoryInvalid,
		Message: "nhóm chi tiêu không hợp lệ",
		Code:    203,
	},
}
