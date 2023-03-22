package errorcode

import "expense-tracker/internal/response"

const (
	CommonCode = "common_codeo"
)

var common = []response.Code{
	{
		Key:     CommonCode,
		Message: "lỗi common",
		Code:    1,
	},
}
