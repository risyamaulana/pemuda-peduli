package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/program_donasi/common/constants"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/domain/interfaces"
	"pemuda-peduli/src/program_donasi/infrastructure/repository"
	"strings"
	"time"

	penggalangDanaDom "pemuda-peduli/src/penggalang_dana/domain"
)

func UpdateProgramDonasi(ctx context.Context, db *db.ConnectTo, data entity.ProgramDonasiEntity, dataDetail entity.ProgramDonasiDetailEntity, id string) (response entity.ProgramDonasiEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRepository(db)

	currentDate := time.Now().UTC()

	// Check available data
	checkData, err := GetProgramDonasi(ctx, db, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	// Check SEO URL
	if data.SEOURL == "" {
		data.SEOURL = strings.ToLower(strings.ReplaceAll(data.Title, " ", "-"))
	}

	// Check Penggalang dana
	if data.IDPPCPPenggalangDana != "" {
		checkPenggalangDana, errCheckPenggalangDana := penggalangDanaDom.GetPenggalangDana(ctx, db, data.IDPPCPPenggalangDana)
		if errCheckPenggalangDana != nil {
			err = errors.New("Failed, penggalang dana not found")
		}
		data.PenggalangDana = checkPenggalangDana
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
	checkData.Nominal = data.Nominal
	checkData.SEOURL = data.SEOURL
	checkData.Target = data.Target
	checkData.IDPPCPPenggalangDana = data.IDPPCPPenggalangDana

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

	log.Println("DATA DETAIL UPDATE : ", utility.PrettyPrint(checkData.Detail))
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

	response, _ = GetProgramDonasi(ctx, db, checkData.IDPPCPProgramDonasi)

	return
}

func UpdateProgramDonasiNews(ctx context.Context, db *db.ConnectTo, data entity.ProgramDonasiNewsEntity, id int64) (response entity.ProgramDonasiNewsEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRepository(db)

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
	data.IDPPCPProgramDonasi = checkData.IDPPCPProgramDonasi

	data.UpdatedAt = &currentDate

	response, err = repo.UpdateNews(ctx, checkData, id)
	if err != nil {
		return
	}

	return
}

func PublishProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string) (response entity.ProgramDonasiEntity, err error) {
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

func HideProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string) (response entity.ProgramDonasiEntity, err error) {
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

func UpdateDonationCollect(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string, amount float64) (response entity.ProgramDonasiEntity, err error) {
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

func DeleteProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, id string) (response entity.ProgramDonasiEntity, err error) {
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

func DeleteProgramDonasiNews(ctx context.Context, repo interfaces.IProgramDonasiRepository, id int64) (response entity.ProgramDonasiNewsEntity, err error) {
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
