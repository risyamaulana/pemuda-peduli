package domain

import (
	"context"
	"pemuda-peduli/src/token/domain/entity"
	"pemuda-peduli/src/token/domain/interfaces"
	"time"
)

func CreateOrUpdateToken(ctx context.Context, repo interfaces.ITokenRepository, data *entity.TokenEntity) (err error) {
	currentTime := time.Now().UTC()

	// Check Available device data
	checkData, err := repo.CheckDevice(data.DeviceID, data.DeviceType)
	if err != nil {
		err = repo.Insert(ctx, data)
	} else {
		data.ID = checkData.ID
		data.IsLogin = false
		data.LoginID = checkData.LoginID
		data.UpdatedAt = &currentTime

		err = repo.Update(ctx, data)
	}

	return
}
