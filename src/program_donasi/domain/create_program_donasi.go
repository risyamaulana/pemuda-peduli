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

	kategoriProgramDonasiDom "pemuda-peduli/src/kategori_program_donasi/domain"
	kategoriProgramDonasiRep "pemuda-peduli/src/kategori_program_donasi/infrastructure/repository"
)

func CreateProgramDonasi(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiEntity, dataDetail *entity.ProgramDonasiDetailEntity) (response entity.ProgramDonasiEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRepository(db)
	kategoriProgramDonasiRepo := kategoriProgramDonasiRep.NewKategoriProgramDonasiRepository(db)

	// check kategori
	if data.KategoriID != "" {
		kategoriData, errGetKategoriData := kategoriProgramDonasiDom.GetKategoriProgramDonasi(ctx, &kategoriProgramDonasiRepo, data.KategoriID)
		if errGetKategoriData != nil {
			err = errors.New("failed, kategori not found")
			return
		}
		data.KategoriID = kategoriData.IDPPCPKategoriProgramDonasi
		data.KategoriName = kategoriData.Name
	}

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

func CreateProgramDonasiNews(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiNewsEntity) (err error) {
	// Repo
	repo := repository.NewProgramDonasiRepository(db)

	err = repo.InsertNews(ctx, data)
	return
}
