package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/partner_kami/common/constants"
	"pemuda-peduli/src/partner_kami/domain/entity"
	"pemuda-peduli/src/partner_kami/domain/interfaces"
	"time"
)

func UpdatePartnerKami(ctx context.Context, repo interfaces.IPartnerKamiRepository, data entity.PartnerKamiEntity, id string) (response entity.PartnerKamiEntity, err error) {
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
	checkData.Name = data.Name
	checkData.ThumbnailImageURL = data.ThumbnailImageURL

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)
	return
}

func PublishPartnerKami(ctx context.Context, repo interfaces.IPartnerKamiRepository, id string) (response entity.PartnerKamiEntity, err error) {
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

func HidePartnerKami(ctx context.Context, repo interfaces.IPartnerKamiRepository, id string) (response entity.PartnerKamiEntity, err error) {
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

func DeletePartnerKami(ctx context.Context, repo interfaces.IPartnerKamiRepository, id string) (response entity.PartnerKamiEntity, err error) {
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
