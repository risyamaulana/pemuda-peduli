package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/kontak_kami/common/constants"
	"pemuda-peduli/src/kontak_kami/domain/entity"
	"pemuda-peduli/src/kontak_kami/domain/interfaces"
	"time"
)

func UpdateKontakKami(ctx context.Context, repo interfaces.IKontakKamiRepository, data entity.KontakKamiEntity, id string) (response entity.KontakKamiEntity, err error) {
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
	checkData.SKLegalitas = data.SKLegalitas
	checkData.Address = data.Address
	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishKontakKami(ctx context.Context, repo interfaces.IKontakKamiRepository, id string) (response entity.KontakKamiEntity, err error) {
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

func HideKontakKami(ctx context.Context, repo interfaces.IKontakKamiRepository, id string) (response entity.KontakKamiEntity, err error) {
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

func DeleteKontakKami(ctx context.Context, repo interfaces.IKontakKamiRepository, id string) (response entity.KontakKamiEntity, err error) {
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
