package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/banner/common/constants"
	"pemuda-peduli/src/banner/domain/entity"
	"pemuda-peduli/src/banner/domain/interfaces"
	"time"
)

func UpdateBanner(ctx context.Context, repo interfaces.IBannerRepository, data entity.BannerEntity, id string) (response entity.BannerEntity, err error) {
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
	checkData.Title = data.Title
	checkData.SubTitle = data.SubTitle
	checkData.TitleContent = data.TitleContent
	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.TitleButtonRight = data.TitleButtonRight
	checkData.DeeplinkRight = data.DeeplinkRight
	checkData.TitleButtonLeft = data.TitleButtonLeft
	checkData.DeeplinkLeft = data.DeeplinkLeft
	checkData.Description = data.Description

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishBanner(ctx context.Context, repo interfaces.IBannerRepository, id string) (response entity.BannerEntity, err error) {
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

func HideBanner(ctx context.Context, repo interfaces.IBannerRepository, id string) (response entity.BannerEntity, err error) {
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

func DeleteBanner(ctx context.Context, repo interfaces.IBannerRepository, id string) (response entity.BannerEntity, err error) {
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
