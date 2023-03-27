package requestmodel

import (
	"expense-tracker/internal/constant"
	mgmodel "expense-tracker/internal/model/mg"
	pgenerate "expense-tracker/internal/util/generate"
	"expense-tracker/internal/util/ptime"
	"expense-tracker/pkg/admin/errorcode"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

// TokenPayload ...
type TokenPayload struct {
	ID          string   `json:"_id"`
	Name        string   `json:"name"`
	Phone       string   `json:"phone,omitempty"`
	Email       string   `json:"email,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	IsRoot      bool     `json:"isRoot"`
}

// RegisterPayload ...
type RegisterPayload struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (m RegisterPayload) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required.Error(errorcode.StaffAuthRequiredEmail),
			validation.Match(regexp.MustCompile(constant.RegexEmail)).Error(errorcode.StaffInvalidFormatEmail)),
		validation.Field(&m.Name, validation.Required.Error(errorcode.CategoryRequiredName)),
		validation.Field(&m.Password, validation.Required.Error(errorcode.StaffAuthRequiredPassword)),
		validation.Field(&m.PasswordConfirm, validation.Required.Error(errorcode.StaffAuthRequiredPasswordConfirm)),
	)
}

func (m RegisterPayload) ConvertToBSON() mgmodel.Staff {
	return mgmodel.Staff{
		ID:        primitive.NewObjectID(),
		Email:     m.Email,
		Name:      m.Name,
		Password:  pgenerate.HashPassword(m.Password),
		IsRoot:    false,
		CreatedAt: ptime.Now(),
		UpdatedAt: ptime.Now(),
	}
}

// LoginPayload ...
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m LoginPayload) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required.Error(errorcode.StaffAuthRequiredEmail)),
		validation.Field(&m.Password, validation.Required.Error(errorcode.StaffAuthRequiredPassword)),
	)
}
