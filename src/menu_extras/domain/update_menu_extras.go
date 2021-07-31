package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/menu_extras/common/constants"
	"pemuda-peduli/src/menu_extras/domain/entity"
	"pemuda-peduli/src/menu_extras/domain/interfaces"
	"time"
)

func UpdateMenuExtras(ctx context.Context, repo interfaces.IMenuExtrasRepository, data entity.MenuExtrasEntity, id string) (response entity.MenuExtrasEntity, err error) {
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
	checkData.Link = data.Link

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishMenuExtras(ctx context.Context, repo interfaces.IMenuExtrasRepository, id string) (response entity.MenuExtrasEntity, err error) {
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

func HideMenuExtras(ctx context.Context, repo interfaces.IMenuExtrasRepository, id string) (response entity.MenuExtrasEntity, err error) {
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

func DeleteMenuExtras(ctx context.Context, repo interfaces.IMenuExtrasRepository, id string) (response entity.MenuExtrasEntity, err error) {
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
