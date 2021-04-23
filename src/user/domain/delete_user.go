package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/user/domain/entity"
	"pemuda-peduli/src/user/domain/interfaces"
)

func RemoveDeleteUser(ctx context.Context, repo interfaces.IUserRepository, id string) (response entity.UserEntity, err error) {
	// Check data
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Failed: user not found")
		return
	}

	checkData.IsDeleted = true

	//Update User
	response, err = repo.Update(ctx, checkData)
	return
}
