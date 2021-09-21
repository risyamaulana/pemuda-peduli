package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/program_kami/common/constants"
	"pemuda-peduli/src/program_kami/domain/entity"
	"pemuda-peduli/src/program_kami/domain/interfaces"
	"time"
)

func UpdateProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, data entity.ProgramKamiEntity, dataDetail entity.ProgramKamiDetailEntity, id string) (response entity.ProgramKamiEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := GetProgramKami(ctx, repo, id)
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
	checkData.Tag = data.Tag
	checkData.BeneficariesImageURL = data.BeneficariesImageURL
	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.Description = data.Description

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	response, err = repo.Update(ctx, checkData, id)

	_, err = repo.Update(ctx, checkData, id)
	if err != nil {
		return
	}

	if checkData.Detail.IDPPCPProgramKamiDetail != "" {
		// Update Data Detail
		checkData.Detail.Content = dataDetail.Content
		checkData.Detail.Tag = data.Tag
		if _, errUpdateDetail := repo.UpdateDetail(ctx, checkData.Detail, checkData.Detail.IDPPCPProgramKamiDetail); errUpdateDetail != nil {
			log.Println("Failed update berita detail: ", errUpdateDetail)
		}
	} else {
		// Insert Data Detail
		// Insert Detail
		dataDetail.IDPPCPProgramKami = checkData.IDPPCPProgramKami
		dataDetail.Tag = data.Tag
		if errDetail := repo.InsertDetail(ctx, &dataDetail); errDetail != nil {
			log.Println("ERR Insert Detail: ", errDetail)
		}
	}

	response, _ = GetProgramKami(ctx, repo, checkData.IDPPCPProgramKami)

	return
}

func PublishProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, id string) (response entity.ProgramKamiEntity, err error) {
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

func HideProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, id string) (response entity.ProgramKamiEntity, err error) {
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

func DeleteProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, id string) (response entity.ProgramKamiEntity, err error) {
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
