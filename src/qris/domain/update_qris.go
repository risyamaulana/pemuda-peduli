package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/qris/common/constants"
	"pemuda-peduli/src/qris/domain/entity"
	"pemuda-peduli/src/qris/domain/interfaces"
	"time"
)

func UpdateQris(ctx context.Context, repo interfaces.IQrisRepository, data entity.QrisEntity, id string) (response entity.QrisEntity, err error) {
	currentDate := time.Now()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	if checkData.Status == constants.StatusDeleted {
		err = errors.New("Can't update this data")
		return
	}
	checkData.Title = data.Title
	checkData.Description = data.Description
	checkData.ThumbnailImageURL = data.ThumbnailImageURL

	checkData.UpdatedAt = &currentDate

	response, err = repo.Update(ctx, checkData, id)
	return
}

func DeleteQris(ctx context.Context, repo interfaces.IQrisRepository, id string) (response entity.QrisEntity, err error) {
	currentDate := time.Now()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	checkData.Status = constants.StatusDeleted
	checkData.UpdatedAt = &currentDate
	response, err = repo.Update(ctx, checkData, id)
	return
}
