package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/user/domain/entity"
	"pemuda-peduli/src/user/domain/interfaces"
	"strings"
	"time"
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

func ForgotPassword(ctx context.Context, repo interfaces.IUserRepository, email string) (response entity.UserEntity, err error) {
	// Get data by email
	data, err := ReadUserByEmail(ctx, repo, email)
	if err != nil {
		err = errors.New("Failed: user not found")
		return
	}
	tokenValid := time.Now().UTC().Add(5 * time.Hour)

	data.IsReset = true
	data.TokenReset = utility.GenerateSalt(15)

	data.TokenValid = &tokenValid

	response, err = repo.Update(ctx, data)

	// Send mail
	sendMailForgotPassword(data)
	return
}

func sendMailForgotPassword(data entity.UserEntity) (err error) {
	to := []string{data.Email}
	var msgStr strings.Builder
	msgStr.WriteString("Berikut token untuk reset password : " + data.TokenReset + "\n")
	msgStr.WriteString("Berikut url token untuk reset : " + "http://ayokitapeduli.com/reset?token=" + data.TokenReset)
	err = infrastructure.SendMail(to, "Reset password", msgStr.String())
	return
}

func ResetPassword(ctx context.Context, repo interfaces.IUserRepository, token, newPassword string) (response entity.UserEntity, err error) {
	// Get data by token
	data, err := repo.GetByToken(ctx, token)
	if err != nil {
		err = errors.New("Failed: token not found / unauthorized")
		return
	}

	if !data.IsReset {
		err = errors.New("Failed, data unavailable to reset")
		return
	}

	if data.TokenValid != nil {
		if time.Now().UTC().After(*data.TokenValid) {
			err = errors.New("failed, token is expired")
			return
		}
	} else {
		err = errors.New("failed, token is expired")
		return
	}

	salt := utility.GenerateSalt(4)
	data.Salt = salt
	data.Password = utility.GeneratePass(salt, newPassword)

	data.IsReset = false
	data.TokenReset = "-"
	data.TokenValid = nil

	response, err = repo.Update(ctx, data)
	return
}
