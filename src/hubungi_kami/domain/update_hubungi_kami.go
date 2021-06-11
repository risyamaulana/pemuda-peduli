package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/hubungi_kami/common/constants"
	"pemuda-peduli/src/hubungi_kami/domain/entity"
	"pemuda-peduli/src/hubungi_kami/domain/interfaces"
	"time"
)

func UpdateHubungiKami(ctx context.Context, repo interfaces.IHubungiKamiRepository, data entity.HubungiKamiEntity, id string) (response entity.HubungiKamiEntity, err error) {
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
	checkData.Icon = data.Icon
	checkData.Link = data.Link
	checkData.Title = data.Title

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishHubungiKami(ctx context.Context, repo interfaces.IHubungiKamiRepository, id string) (response entity.HubungiKamiEntity, err error) {
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

func HideHubungiKami(ctx context.Context, repo interfaces.IHubungiKamiRepository, id string) (response entity.HubungiKamiEntity, err error) {
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

func DeleteHubungiKami(ctx context.Context, repo interfaces.IHubungiKamiRepository, id string) (response entity.HubungiKamiEntity, err error) {
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
