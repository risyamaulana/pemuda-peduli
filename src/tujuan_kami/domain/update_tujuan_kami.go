package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/tujuan_kami/common/constants"
	"pemuda-peduli/src/tujuan_kami/domain/entity"
	"pemuda-peduli/src/tujuan_kami/domain/interfaces"
	"time"
)

func UpdateTujuanKami(ctx context.Context, repo interfaces.ITujuanKamiRepository, data entity.TujuanKamiEntity, id string) (response entity.TujuanKamiEntity, err error) {
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
	checkData.Title = data.Title
	checkData.Description = data.Description
	checkData.Icon = data.Icon

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishTujuanKami(ctx context.Context, repo interfaces.ITujuanKamiRepository, id string) (response entity.TujuanKamiEntity, err error) {
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
	checkData.Status = constants.StatusPublished
	checkData.UpdatedAt = &currentDate
	checkData.PublishedAt = &currentDate
	checkData.IsDeleted = false
	response, err = repo.Update(ctx, checkData, id)
	return
}

func HideTujuanKami(ctx context.Context, repo interfaces.ITujuanKamiRepository, id string) (response entity.TujuanKamiEntity, err error) {
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

	checkData.Status = constants.StatusHidden
	checkData.UpdatedAt = &currentDate
	checkData.PublishedAt = nil
	checkData.IsDeleted = false
	response, err = repo.Update(ctx, checkData, id)
	return
}

func DeleteTujuanKami(ctx context.Context, repo interfaces.ITujuanKamiRepository, id string) (response entity.TujuanKamiEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	checkData.Status = constants.StatusDeleted
	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = true
	response, err = repo.Update(ctx, checkData, id)
	return
}
