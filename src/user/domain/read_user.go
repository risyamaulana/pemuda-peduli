package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/user/domain/entity"
	"pemuda-peduli/src/user/domain/interfaces"
)

// Get
func ReadUser(ctx context.Context, repo interfaces.IUserRepository, id string) (response entity.UserEntity, err error) {
	response, err = repo.Get(ctx, id)
	return
}

// Get by email
func ReadUserByEmail(ctx context.Context, repo interfaces.IUserRepository, email string) (response entity.UserEntity, err error) {
	response, err = repo.GetByEmail(ctx, email)
	return
}

// Get Login
func ReadLoginuser(ctx context.Context, repo interfaces.IUserRepository, username string) (response entity.UserEntity, err error) {
	username = utility.UsernameIsPhoneNumber(username)
	response, err = repo.GetForLogin(ctx, username)
	if err != nil {
		err = errors.New("Failed: user not found")
		return
	}
	return
}

// Get duplicateCheck
func ReadDuplicateCheck(ctx context.Context, repo interfaces.IUserRepository, username, phoneNumber, email string) (response entity.UserEntity, err error) {
	response, err = repo.GetDuplicateCheck(ctx, username, phoneNumber, email)
	return
}
