package domain

import (
	"context"
	"pemuda-peduli/src/token/domain/entity"
	"pemuda-peduli/src/token/domain/interfaces"
)

func UpdateToken(ctx context.Context, repo interfaces.ITokenRepository, data *entity.TokenEntity) (err error) {
	// Check Available device data
	checkData, err := repo.CheckTokenDevice(data.Token, data.DeviceID, data.DeviceType)
	if err != nil {
		return
	}
	checkData.IsLogin = data.IsLogin
	checkData.LoginID = data.LoginID
	err = repo.Update(ctx, &checkData)
	return
}
