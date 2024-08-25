package schema

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserSchema struct {
	Fullname    string      `json:"fullname" validate:"required"`
	Email       string      `json:"email" validate:"required"`
	Password    string      `json:"password"`
	Phone       null.String `json:"phone"`
	TokenVerify null.String `json:"token_verify"`
	Address     null.String `json:"address"`
	IsActive    bool        `json:"is_active"`
	IsBlocked   bool        `json:"is_blocked"`
	RoleID      uuid.UUID   `json:"role_id" validate:"required"`
}
