package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"

	"pemuda-peduli/src/program_donasi_rutin/common/constants"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
	"time"

	kategoriDom "pemuda-peduli/src/program_donasi_kategori/domain"
	kategoriRep "pemuda-peduli/src/program_donasi_kategori/infrastructure/repository"
)

func EditProgramDonasiRutin(ctx context.Context, db *db.ConnectTo, data entity.ProgramDonasiRutinEntity, dataDetail entity.ProgramDonasiRutinDetailEntity, id string) (response entity.ProgramDonasiRutinEntity, err error) {
	repo := repository.NewProgramDonasiRutinRepository(db)
	kategoriRepo := kategoriRep.NewProgramDonasiKategoriRepository(db)

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

	// Check Kategori
	if checkData.IDPPCPProgramDonasiKategori != data.IDPPCPProgramDonasiKategori {
		kategoriData, errKategoriData := kategoriDom.GetProgramDonasiKategori(ctx, &kategoriRepo, data.IDPPCPProgramDonasiKategori)
		if errKategoriData != nil {
			err = errors.New("Failed, kategori not found")
			return
		}

		data.KategoriName = kategoriData.KategoriName
	}

	log.Println("IS SHOW DATA DOMAIN : ", data.IsShow)

	response, err = updateProgramDonasiRutin(ctx, &repo, data, dataDetail, id)

	return
}

func updateProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data entity.ProgramDonasiRutinEntity, dataDetail entity.ProgramDonasiRutinDetailEntity, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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
	checkData.Tag = data.Tag
	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.Description = data.Description

	checkData.IDPPCPProgramDonasiKategori = data.IDPPCPProgramDonasiKategori
	checkData.KategoriName = data.KategoriName

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false
	checkData.IsShow = data.IsShow

	_, err = repo.Update(ctx, checkData, id)
	if err != nil {
		return
	}

	if checkData.Detail.IDPPCPProgramDonasiRutinDetail != "" {
		// Update Data Detail
		checkData.Detail.Content = dataDetail.Content
		checkData.Detail.Tag = data.Tag
		if _, errUpdateDetail := repo.UpdateDetail(ctx, checkData.Detail, checkData.Detail.IDPPCPProgramDonasiRutinDetail); errUpdateDetail != nil {
			log.Println("Failed update berita detail: ", errUpdateDetail)
		}
	} else {
		// Insert Data Detail
		// Insert Detail
		dataDetail.IDPPCPProgramDonasiRutin = checkData.IDPPCPProgramDonasiRutin
		dataDetail.Tag = data.Tag
		if errDetail := repo.InsertDetail(ctx, &dataDetail); errDetail != nil {
			log.Println("ERR Insert Detail: ", errDetail)
		}
	}

	response, _ = GetProgramDonasiRutin(ctx, repo, checkData.IDPPCPProgramDonasiRutin)

	return
}

func PublishProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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

func HideProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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

func DeleteProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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
