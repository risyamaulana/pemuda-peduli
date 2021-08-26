package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/infrastructure/repository"
	"strings"

	penggalangDanaDom "pemuda-peduli/src/penggalang_dana/domain"
)

func CreateProgramDonasi(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiEntity, dataDetail *entity.ProgramDonasiDetailEntity) (response entity.ProgramDonasiEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRepository(db)

	// Check Penggalang dana
	if data.IDPPCPPenggalangDana != "" {
		checkPenggalangDana, errCheckPenggalangDana := penggalangDanaDom.GetPenggalangDana(ctx, db, data.IDPPCPPenggalangDana)
		if errCheckPenggalangDana != nil {
			err = errors.New("Failed, penggalang dana not found")
		}
		data.PenggalangDana = checkPenggalangDana
	}

	// Check SEO URL
	if data.SEOURL == "" {
		data.SEOURL = strings.ToLower(strings.ReplaceAll(data.Title, " ", "-"))
	}

	err = repo.Insert(ctx, data)

	// Insert Detail
	dataDetail.IDPPCPProgramDonasi = data.IDPPCPProgramDonasi
	dataDetail.Tag = data.Tag
	if errDetail := repo.InsertDetail(ctx, dataDetail); errDetail != nil {
		log.Println("ERR Insert Detail: ", errDetail)
	}

	response, _ = GetProgramDonasi(ctx, db, data.IDPPCPProgramDonasi)

	return
}
