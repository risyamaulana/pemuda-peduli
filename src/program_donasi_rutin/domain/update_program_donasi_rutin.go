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
	"strings"
	"time"
	// kategoriDom "pemuda-peduli/src/program_donasi_kategori/domain"
	// kategoriRep "pemuda-peduli/src/program_donasi_kategori/infrastructure/repository"
)

func EditProgramDonasiRutin(ctx context.Context, db *db.ConnectTo, data entity.ProgramDonasiRutinEntity, id string) (response entity.ProgramDonasiRutinEntity, err error) {
	repo := repository.NewProgramDonasiRutinRepository(db)
	// kategoriRepo := kategoriRep.NewProgramDonasiKategoriRepository(db)

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
	// if checkData.IDPPCPProgramDonasiKategori != data.IDPPCPProgramDonasiKategori {
	// 	kategoriData, errKategoriData := kategoriDom.GetProgramDonasiKategori(ctx, &kategoriRepo, data.IDPPCPProgramDonasiKategori)
	// 	if errKategoriData != nil {
	// 		err = errors.New("Failed, kategori not found")
	// 		return
	// 	}

	// 	data.KategoriName = kategoriData.KategoriName
	// }

	// Check SEO URL
	if data.SEOURL == "" {
		data.SEOURL = strings.ToLower(strings.ReplaceAll(data.Title, " ", "-"))
	}

	log.Println("IS SHOW DATA DOMAIN : ", data.IsShow)

	response, err = updateProgramDonasiRutin(ctx, &repo, data, id)

	return
}

func updateProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data entity.ProgramDonasiRutinEntity, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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
	checkData.SubTitle = data.SubTitle
	checkData.Tag = data.Tag
	checkData.Content = data.Content
	checkData.ThumbnailImageURL = data.ThumbnailImageURL
	checkData.Description = data.Description
	checkData.SEOURL = data.SEOURL

	// checkData.IDPPCPProgramDonasiKategori = data.IDPPCPProgramDonasiKategori
	// checkData.KategoriName = data.KategoriName

	checkData.IDPPCPMasterQris = data.IDPPCPMasterQris
	checkData.QrisImageURL = data.QrisImageURL

	checkData.UpdatedAt = &currentDate
	checkData.IsDeleted = false
	checkData.IsShow = data.IsShow

	_, err = repo.Update(ctx, checkData, id)
	if err != nil {
		return
	}

	response, _ = GetProgramDonasiRutin(ctx, repo, checkData.IDPPCPProgramDonasiRutin)

	return
}

func UpdateProgramDonasiNews(ctx context.Context, db *db.ConnectTo, data entity.ProgramDonasiRutinNewsEntity, id int64) (response entity.ProgramDonasiRutinNewsEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	currentDate := time.Now().UTC()

	// Check available data
	checkData, err := GetProgramDonasiNews(ctx, db, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	if checkData.IsDeleted {
		err = errors.New("Can't update this data")
		return
	}

	data.ID = checkData.ID
	data.IDPPCPProgramDonasiRutin = checkData.IDPPCPProgramDonasiRutin

	data.UpdatedAt = &currentDate

	response, err = repo.UpdateNews(ctx, checkData, id)
	if err != nil {
		return
	}

	return
}

func PublishProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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

func HideProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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

func UpdateDonationCollect(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string, amount float64) (response entity.ProgramDonasiRutinEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	checkData, err := repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	checkData.DonationCollect = checkData.DonationCollect + amount
	checkData.UpdatedAt = &currentDate

	response, err = repo.Update(ctx, checkData, id)
	return
}

func DeleteProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (response entity.ProgramDonasiRutinEntity, err error) {
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

func DeleteProgramDonasiNews(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id int64) (response entity.ProgramDonasiRutinNewsEntity, err error) {
	currentDate := time.Now().UTC()
	// Check available daata
	data, err := repo.GetNews(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	data.IsDeleted = true
	data.UpdatedAt = &currentDate

	response, err = repo.UpdateNews(ctx, data, id)
	return
}
