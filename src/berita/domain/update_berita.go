package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/berita/common/constants"
	"pemuda-peduli/src/berita/domain/entity"
	"pemuda-peduli/src/berita/domain/interfaces"
	"time"
)

func UpdateBerita(ctx context.Context, repo interfaces.IBeritaRepository, data entity.BeritaEntity, dataDetail entity.BeritaDetailEntity, id string) (response entity.BeritaEntity, err error) {
	currentDate := time.Now()
	// Check available daata
	checkData, err := GetBerita(ctx, repo, id)
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
	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.Description = data.Description

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false

	_, err = repo.Update(ctx, checkData, id)
	if err != nil {
		return
	}

	if checkData.Detail.IDPPCPBeritaDetail != "" {
		// Update Data Detail
		checkData.Detail.Content = dataDetail.Content
		checkData.Detail.Tag = data.Tag
		if _, errUpdateDetail := repo.UpdateDetail(ctx, checkData.Detail, checkData.Detail.IDPPCPBeritaDetail); errUpdateDetail != nil {
			log.Println("Failed update berita detail: ", errUpdateDetail)
		}
	} else {
		// Insert Data Detail
		// Insert Detail
		dataDetail.IDPPCPBerita = checkData.IDPPCPBerita
		dataDetail.Tag = data.Tag
		if errDetail := repo.InsertDetail(ctx, &dataDetail); errDetail != nil {
			log.Println("ERR Insert Detail: ", errDetail)
		}
	}

	response, _ = GetBerita(ctx, repo, checkData.IDPPCPBerita)

	return
}

func PublishBerita(ctx context.Context, repo interfaces.IBeritaRepository, id string) (response entity.BeritaEntity, err error) {
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

func HideBerita(ctx context.Context, repo interfaces.IBeritaRepository, id string) (response entity.BeritaEntity, err error) {
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

func DeleteBerita(ctx context.Context, repo interfaces.IBeritaRepository, id string) (response entity.BeritaEntity, err error) {
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
