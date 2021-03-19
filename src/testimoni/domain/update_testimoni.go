package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/testimoni/common/constants"
	"pemuda-peduli/src/testimoni/domain/entity"
	"pemuda-peduli/src/testimoni/domain/interfaces"
	"time"
)

func UpdateTestimoni(ctx context.Context, repo interfaces.ITestimoniRepository, data entity.TestimoniEntity, id string) (response entity.TestimoniEntity, err error) {
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
	checkData.Name = data.Name
	checkData.Role = data.Role
	checkData.ThumbnailPhotoURL = data.ThumbnailPhotoURL
	checkData.Message = data.Message

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishTestimoni(ctx context.Context, repo interfaces.ITestimoniRepository, id string) (response entity.TestimoniEntity, err error) {
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
	checkData.Status = constants.StatusPublished
	checkData.UpdatedAt = &currentDate
	checkData.PublishedAt = &currentDate
	checkData.IsDeleted = false
	response, err = repo.Update(ctx, checkData, id)
	return
}

func HideTestimoni(ctx context.Context, repo interfaces.ITestimoniRepository, id string) (response entity.TestimoniEntity, err error) {
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

	checkData.Status = constants.StatusHidden
	checkData.UpdatedAt = &currentDate
	checkData.PublishedAt = nil
	checkData.IsDeleted = false
	response, err = repo.Update(ctx, checkData, id)
	return
}

func DeleteTestimoni(ctx context.Context, repo interfaces.ITestimoniRepository, id string) (response entity.TestimoniEntity, err error) {
	currentDate := time.Now()
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
