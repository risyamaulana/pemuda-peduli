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

func FindAdminUser(ctx context.Context, db *db.ConnectTo, data *entity.AdminUserQueryEntity) (response []entity.AdminUserEntity, count int, err error) {
	repo := repository.NewAdminUserRepository(db)
	response, count, err = repo.Find(ctx, data)

	for i, data := range response {
		// Get Role Data
		roleRepository := roleRepo.NewRoleRepository(db)
		roleData, errRole := roleDom.GetRole(ctx, &roleRepository, data.Role)
		if errRole != nil {
			err = errors.New("Role data not found")
			return
		}
		response[i].RoleData = roleData
	}
	return
}

func GetAdminUser(ctx context.Context, db *db.ConnectTo, id string) (response entity.AdminUserEntity, err error) {
	repo := repository.NewAdminUserRepository(db)
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	// Get Role Data
	roleRepository := roleRepo.NewRoleRepository(db)
	roleData, err := roleDom.GetRole(ctx, &roleRepository, response.Role)
	if err != nil {
		err = errors.New("Role data not found")
		return
	}

	response.RoleData = roleData
	return
}

func GetAdminUserByUsername(ctx context.Context, db *db.ConnectTo, username string) (response entity.AdminUserEntity, err error) {
	repo := repository.NewAdminUserRepository(db)
	response, err = repo.GetByUsername(ctx, username)
	if err != nil {
		err = errors.New("User not found")
		return
	}

	// Get Role Data
	roleRepository := roleRepo.NewRoleRepository(db)
	roleData, err := roleDom.GetRole(ctx, &roleRepository, response.Role)
	if err != nil {
		err = errors.New("Role data not found")
		return
	}

	response.RoleData = roleData
	return
}
