package dto

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserDto struct {
	ID          uuid.UUID   `db:"id" json:"id"`
	CreatedAt   null.Time   `db:"created_at" json:"created_at"`
	UpdatedAt   null.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt   null.Time   `db:"deleted_at" json:"deleted_at"`
	Fullname    string      `db:"fullname" json:"fullname"`
	Email       string      `db:"email" json:"email"`
	Password    null.String `db:"password" json:"password"`
	Phone       null.String `db:"phone" json:"phone"`
	TokenVerify null.String `db:"token_verify" json:"token_verify"`
	Address     null.String `db:"address" json:"address"`
	IsActive    bool        `db:"is_active" json:"is_active"`
	IsBlocked   bool        `db:"is_blocked" json:"is_blocked"`
	RoleID      uuid.UUID   `db:"role_id" json:"role_id"`
	RoleName    string      `db:"role_name" json:"role_name"`
}
