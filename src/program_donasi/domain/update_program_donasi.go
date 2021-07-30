package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/program_donasi/common/constants"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/domain/interfaces"
	"time"
)

func UpdateProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, data entity.ProgramDonasiEntity, dataDetail entity.ProgramDonasiDetailEntity, id string) (response entity.ProgramDonasiEntity, err error) {
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
	// checkData.DonasiType = data.DonasiType
	checkData.Tag = data.Tag
	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.Description = data.Description
	checkData.ValidFrom = data.ValidFrom
	checkData.ValidTo = data.ValidTo
	checkData.Target = data.Target

	checkData.KitaBisaLink = data.KitaBisaLink
	checkData.AyoBantuLink = data.AyoBantuLink

	checkData.IDPPCPMasterQris = data.IDPPCPMasterQris
	checkData.QrisImageURL = data.QrisImageURL

	checkData.Description = data.Description

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false
	checkData.IsShow = data.IsShow

	_, err = repo.Update(ctx, checkData, id)
	if err != nil {
		return
	}

	if checkData.Detail.IDPPCPProgramDonasiDetail != "" {
		// Update Data Detail
		checkData.Detail.Content = dataDetail.Content
		checkData.Detail.Tag = data.Tag
		if _, errUpdateDetail := repo.UpdateDetail(ctx, checkData.Detail, checkData.Detail.IDPPCPProgramDonasiDetail); errUpdateDetail != nil {
			log.Println("Failed update berita detail: ", errUpdateDetail)
		}
	} else {
		// Insert Data Detail
		// Insert Detail
		dataDetail.IDPPCPProgramDonasi = checkData.IDPPCPProgramDonasi
		dataDetail.Tag = data.Tag
		if errDetail := repo.InsertDetail(ctx, &dataDetail); errDetail != nil {
			log.Println("ERR Insert Detail: ", errDetail)
		}
	}

	response, _ = GetProgramDonasi(ctx, repo, checkData.IDPPCPProgramDonasi)

	return
}

func PublishProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string) (response entity.ProgramDonasiEntity, err error) {
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

func HideProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string) (response entity.ProgramDonasiEntity, err error) {
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

func DeleteProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string) (response entity.ProgramDonasiEntity, err error) {
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
