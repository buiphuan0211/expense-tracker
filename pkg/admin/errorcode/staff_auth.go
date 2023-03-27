package errorcode

import "expense-tracker/internal/response"

const (
	StaffAuthRequiredEmail           = "staff_auth_required_email"
	StaffAuthRequiredPassword        = "staff_auth_required_password"
	StaffAuthRequiredPasswordConfirm = "staff_auth_required_password_confirm"
	StaffAuthGenerateToken           = "staff_auth_generate_token"
	StaffAuthInvalidEmailOrPassword  = "staff_auth_invalid_email_or_password"
)

var staffAuth = []response.Code{
	{
		Key:     StaffAuthRequiredEmail,
		Message: "email không được trống",
		Code:    400,
	},
	{
		Key:     StaffAuthRequiredPassword,
		Message: "mật khẩu không được trống",
		Code:    401,
	},
	{
		Key:     StaffAuthRequiredPasswordConfirm,
		Message: "mật khẩu xác nhận không được trống",
		Code:    402,
	},
	{
		Key:     StaffAuthGenerateToken,
		Message: "lỗi generate token",
		Code:    403,
	},
	{
		Key:     StaffAuthInvalidEmailOrPassword,
		Message: "email hoặc mật khẩu không đúng",
		Code:    404,
	},
}
