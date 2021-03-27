package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/admin_user/domain/entity"
	"pemuda-peduli/src/admin_user/infrastructure/repository"
	"pemuda-peduli/src/common/infrastructure/db"

	roleDom "pemuda-peduli/src/role/domain"
	roleRepo "pemuda-peduli/src/role/infrastructure/repository"
)

func CreateAdminUser(ctx context.Context, db *db.ConnectTo, data *entity.AdminUserEntity) (err error) {
	// Check Role
	roleLevel := ctx.Value("user_role_level").(int)

	roleRepository := roleRepo.NewRoleRepository(db)
	roleData, err := roleDom.GetRole(ctx, &roleRepository, data.Role)
	if err != nil {
		err = errors.New("Role not found")
		return
	}

	if roleLevel >= roleData.RoleLevel {
		err = errors.New("Can't create user: role does not qualify to create user")
		return
	}
	repo := repository.NewAdminUserRepository(db)

	// Applied role level
	data.RoleLevel = roleData.RoleLevel

	err = repo.Insert(ctx, data)
	return
}
