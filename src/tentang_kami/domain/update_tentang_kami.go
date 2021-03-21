package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/tentang_kami/domain/entity"
	"pemuda-peduli/src/tentang_kami/domain/interfaces"
	"time"
)

func UpdateTentangKami(ctx context.Context, repo interfaces.ITentangKamiRepository, data entity.TentangKamiEntity, id string) (response entity.TentangKamiEntity, err error) {
	currentDate := time.Now()
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

	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.Description = data.Description

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func DeleteTentangKami(ctx context.Context, repo interfaces.ITentangKamiRepository, id string) (response entity.TentangKamiEntity, err error) {
	currentDate := time.Now()
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
