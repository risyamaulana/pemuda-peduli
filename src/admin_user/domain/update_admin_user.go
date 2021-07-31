package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/admin_user/domain/entity"
	"pemuda-peduli/src/admin_user/domain/interfaces"
	"pemuda-peduli/src/admin_user/infrastructure/repository"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	roleDom "pemuda-peduli/src/role/domain"
	roleRepo "pemuda-peduli/src/role/infrastructure/repository"
	"time"
)

func UpdateAdminUser(ctx context.Context, repo interfaces.IAdminUserRepository, data entity.AdminUserEntity, id string) (response entity.AdminUserEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available data
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	if checkData.IsDeleted {
		err = errors.New("Can't update this data")
		return
	}

	checkData.Username = data.Username
	checkData.Email = data.Email
	checkData.NamaLengkap = data.NamaLengkap
	checkData.Alamat = data.Alamat

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func UpdatePassword(ctx context.Context, repo interfaces.IAdminUserRepository, data entity.AdminUserEntity, id string) (err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	if checkData.IsDeleted {
		err = errors.New("Can't update this data")
		return
	}

	checkData.Salt = data.Salt
	checkData.Password = data.Password

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	_, err = repo.Update(ctx, checkData, id)
	return
}

func ResetPassword(ctx context.Context, db *db.ConnectTo, id string) (newPassword string, err error) {
	currentDate := time.Now().UTC()

	repo := repository.NewAdminUserRepository(db)
	// Check available data
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	if checkData.IsDeleted {
		err = errors.New("Can't update this data")
		return
	}

	// Check Role
	roleLevel := ctx.Value("user_role_level").(int)

	roleRepository := roleRepo.NewRoleRepository(db)
	roleData, err := roleDom.GetRole(ctx, &roleRepository, checkData.Role)
	if err != nil {
		err = errors.New("Role not found")
		return
	}

	if roleLevel >= roleData.RoleLevel {
		err = errors.New("Can't create user: role does not qualify to create user")
		return
	}

	// Applied reset password
	newPassword = utility.GenerateSalt(10)
	salt := utility.GenerateSalt(4)
	password := utility.GeneratePass(salt, newPassword)

	checkData.Salt = salt
	checkData.Password = password

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false
	_, err = repo.Update(ctx, checkData, id)
	return
}

//TODO: Change role

func ChangeRole(ctx context.Context, db *db.ConnectTo, id string, role string) (response entity.AdminUserEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	repo := repository.NewAdminUserRepository(db)
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	if checkData.IsDeleted {
		err = errors.New("Can't update this data")
		return
	}

	// Check Role
	roleLevel := ctx.Value("user_role_level").(int)

	roleRepository := roleRepo.NewRoleRepository(db)
	roleData, err := roleDom.GetRole(ctx, &roleRepository, role)
	if err != nil {
		err = errors.New("Role not found")
		return
	}

	if roleLevel >= roleData.RoleLevel {
		err = errors.New("Can't create user: role does not qualify to create user")
		return
	}

	checkData.Role = roleData.IDPPCPMasterRole
	checkData.RoleLevel = roleData.RoleLevel

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false
	response, err = repo.Update(ctx, checkData, id)

	return
}

func DeleteAdminUser(ctx context.Context, repo interfaces.IAdminUserRepository, id string) (response entity.AdminUserEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = true

	response, err = repo.Update(ctx, checkData, id)
	return
}
