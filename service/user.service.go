package service

import (
	"context"
	"fmt"
	"gofiber-sqlc/database"
	"gofiber-sqlc/database/sqlc"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (service *UserService) FindAll(ctx context.Context, c *fiber.Ctx) ([]sqlc.GetUsersRow, int64, error) {
	var data []sqlc.GetUsersRow
	var total int64
	var err error

	const limit int32 = 1000

	queryPage, _ := strconv.Atoi(c.Query("page", "1"))
	queryPageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	queryOffset := (queryPage - 1) * queryPageSize
	queryLimit := int32(queryPageSize)

	if queryPage < 1 {
		queryOffset = 0
	}

	if queryLimit > limit {
		queryLimit = limit
	}

	// get data
	data, err = sqlc.New(database.DB).GetUsers(ctx, sqlc.GetUsersParams{
		Offset: int32(queryOffset),
		Limit:  queryLimit,
	})

	// get total data
	total, _ = sqlc.New(database.DB).CountUser(ctx)

	if err != nil {
		return data, total, err
	}

	return data, total, err
}

func (service *UserService) FindOne(ctx context.Context, id uuid.UUID) (*sqlc.GetUsersRow, error) {
	// get data
	data, err := sqlc.New(database.DB).GetUserWithRelation(ctx, id)

	// dto user
	user := &sqlc.GetUsersRow{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Fullname:  data.Fullname,
		Email:     data.Email,
		Phone:     data.Phone,
		Address:   data.Address,
		IsActive:  data.IsActive,
		IsBlocked: data.IsBlocked,
		RoleID:    data.RoleID,
		RoleName:  data.RoleName,
	}

	if err != nil {
		return user, err
	}

	return user, err
}

func (service *UserService) Create(ctx context.Context, input sqlc.NewUserParams) (*sqlc.User, error) {
	// create user data
	data, err := sqlc.New(database.DB).NewUser(ctx, input)

	user := &sqlc.User{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Fullname:  data.Fullname,
		Email:     data.Email,
		Phone:     data.Phone,
		Address:   data.Address,
		IsActive:  data.IsActive,
		IsBlocked: data.IsBlocked,
		RoleID:    data.RoleID,
	}

	if err != nil {
		return user, err
	}

	return user, err
}

func (service *UserService) Update(ctx context.Context, input sqlc.UpdateUserParams) (uuid.UUID, error) {
	user, err := sqlc.New(database.DB).GetUser(ctx, input.ID)

	if err != nil {
		return user.ID, err
	}

	fmt.Println(user, "TEST", "\n\n", err, "ERR", "\n\n", input)

	// update user data
	data, err := sqlc.New(database.DB).UpdateUser(ctx, input)

	fmt.Println(input, "\n\n", data, "\n\n", err)

	if err != nil {
		return data, err
	}

	return data, err
}
