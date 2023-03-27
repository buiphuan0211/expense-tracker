package errorcode

import "expense-tracker/internal/response"

const (
	StaffNotFound           = "staff_not_found"
	StaffExistedEmail       = "staff_existed_email"
	StaffInvalidFormatEmail = "staff_invalid_format_email"
	StaffInvalidSStatus     = "staff_invalid_status"
)

var staff = []response.Code{
	{
		Key:     StaffNotFound,
		Message: "không tìm thấy nhân viên",
		Code:    300,
	},
	{
		Key:     StaffExistedEmail,
		Message: "email đã tồn tại",
		Code:    301,
	},
	{
		Key:     StaffInvalidFormatEmail,
		Message: "định dạng email không đúng",
		Code:    302,
	},
	{
		Key:     StaffInvalidFormatEmail,
		Message: "định dạng email không đúng",
		Code:    303,
	},
	{
		Key:     StaffInvalidSStatus,
		Message: "trạng thái không hợp lệ",
		Code:    304,
	},
}
