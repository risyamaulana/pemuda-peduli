package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/user/domain/entity"
	"pemuda-peduli/src/user/domain/interfaces"
)

// Register / Create User
func RegisterUser(ctx context.Context, repo interfaces.IUserRepository, data *entity.UserEntity) (err error) {
	// Check available user (duplicate)
	if checkDuplicate, errCheckDuplicate := repo.GetDuplicateCheck(ctx, data.Username, data.PhoneNumber, data.Email); errCheckDuplicate == nil {
		message := "Failed register data already exists for user : " + checkDuplicate.NamaLengkap + ""
		err = errors.New(message)
		return
	}

	err = repo.Insert(ctx, data)
	return
}
