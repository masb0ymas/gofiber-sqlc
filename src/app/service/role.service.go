package service

import (
	"context"
	"gofiber-sqlc/src/database"
	"gofiber-sqlc/src/database/sqlc"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (service *RoleService) FindAll(ctx context.Context, c *fiber.Ctx) ([]sqlc.Role, int64, error) {
	var data []sqlc.Role
	var total int64
	var err error

	const limit int32 = 1000

	queryPage, _ := strconv.Atoi(c.Query("page", "1"))
	queryPageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	queryOffset := int32(queryPage) - 1
	queryLimit := int32(queryPageSize)

	if queryLimit > limit {
		queryLimit = limit
	}

	// get data
	data, err = sqlc.New(database.DB).
		GetRoles(ctx, sqlc.GetRolesParams{Offset: queryOffset, Limit: queryLimit})

	// get total data
	total, _ = sqlc.New(database.DB).CountRole(ctx)

	if err != nil {
		return data, total, err
	}

	return data, total, err
}

func (service *RoleService) FindOne(ctx context.Context, id uuid.UUID) (*sqlc.Role, error) {
	// get data
	data, err := sqlc.New(database.DB).GetRole(ctx, id)

	// dto role
	role := &sqlc.Role{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Name:      data.Name,
	}

	if err != nil {
		return role, err
	}

	return role, err
}

func (service *RoleService) Create(ctx context.Context, name string) (*sqlc.Role, error) {
	// create data
	data, err := sqlc.New(database.DB).NewRole(ctx, name)

	// dto role
	role := &sqlc.Role{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Name:      data.Name,
	}

	if err != nil {
		return role, err
	}

	return role, err
}

func (service *RoleService) Update(ctx context.Context, id uuid.UUID, name string) (*sqlc.Role, error) {
	// update data
	data, err := sqlc.New(database.DB).
		UpdateRole(ctx, sqlc.UpdateRoleParams{ID: id, Name: name})

	// dto role
	role := &sqlc.Role{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Name:      data.Name,
	}

	if err != nil {
		return role, err
	}

	return role, err
}

func (service *RoleService) Restore(ctx context.Context, id uuid.UUID) error {
	// restore data
	err := sqlc.New(database.DB).RestoreRole(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (service *RoleService) SoftDelete(ctx context.Context, id uuid.UUID) error {
	// soft delete data
	err := sqlc.New(database.DB).SoftDeleteRole(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (service *RoleService) ForceDelete(ctx context.Context, id uuid.UUID) error {
	// delete data
	err := sqlc.New(database.DB).ForceDeleteRole(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
