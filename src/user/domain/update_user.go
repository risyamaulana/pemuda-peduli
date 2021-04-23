package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/user/domain/entity"
	"pemuda-peduli/src/user/domain/interfaces"
)

func UpdateUser(ctx context.Context, repo interfaces.IUserRepository, data entity.UserEntity) (response entity.UserEntity, err error) {
	// Check data
	checkData, err := repo.Get(ctx, data.IDUser)
	if err != nil {
		err = errors.New("Failed: user not found")
		return
	}
	data.ID = checkData.ID
	data.Username = checkData.Username
	data.Salt = checkData.Salt
	data.Email = checkData.Email
	data.PhoneNumber = checkData.PhoneNumber
	data.Password = checkData.Password
	data.IsDeleted = checkData.IsDeleted
	data.CreatedAt = checkData.CreatedAt

	// Update data
	response, err = repo.Update(ctx, data)
	return
}

func ChangePassword(ctx context.Context, repo interfaces.IUserRepository, data entity.UserEntity) (response entity.UserEntity, err error) {
	// Check data
	checkData, err := repo.Get(ctx, data.IDUser)
	if err != nil {
		err = errors.New("Failed: user not found")
		return
	}

	checkData.Salt = data.Salt
	checkData.Password = data.Password

	checkData.UpdatedAt = data.UpdatedAt
	// Update data
	response, err = repo.Update(ctx, checkData)
	return
}
