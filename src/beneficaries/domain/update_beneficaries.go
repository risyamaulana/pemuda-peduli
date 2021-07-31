package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/beneficaries/common/constants"
	"pemuda-peduli/src/beneficaries/domain/entity"
	"pemuda-peduli/src/beneficaries/domain/interfaces"
	"time"
)

func UpdateBeneficaries(ctx context.Context, repo interfaces.IBeneficariesRepository, data entity.BeneficariesEntity, id string) (response entity.BeneficariesEntity, err error) {
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
	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.DeeplinkRight = data.DeeplinkRight
	checkData.DeeplinkLeft = data.DeeplinkLeft
	checkData.Description = data.Description

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishBeneficaries(ctx context.Context, repo interfaces.IBeneficariesRepository, id string) (response entity.BeneficariesEntity, err error) {
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

func HideBeneficaries(ctx context.Context, repo interfaces.IBeneficariesRepository, id string) (response entity.BeneficariesEntity, err error) {
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

func DeleteBeneficaries(ctx context.Context, repo interfaces.IBeneficariesRepository, id string) (response entity.BeneficariesEntity, err error) {
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
