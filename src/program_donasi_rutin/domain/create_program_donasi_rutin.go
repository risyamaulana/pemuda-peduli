package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"

	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"

	kategoriDom "pemuda-peduli/src/program_donasi_kategori/domain"
	kategoriRep "pemuda-peduli/src/program_donasi_kategori/infrastructure/repository"
)

func CreateProgramDonasiRutin(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiRutinEntity, dataDetail *entity.ProgramDonasiRutinDetailEntity) (response entity.ProgramDonasiRutinEntity, err error) {
	repo := repository.NewProgramDonasiRutinRepository(db)
	kategoriRepo := kategoriRep.NewProgramDonasiKategoriRepository(db)

	// Check Kategori
	kategoriData, err := kategoriDom.GetProgramDonasiKategori(ctx, &kategoriRepo, data.IDPPCPProgramDonasiKategori)
	if err != nil {
		err = errors.New("Failed, kategori not found")
		return
	}

	data.KategoriName = kategoriData.KategoriName

	insertProgramDonasiRutin(ctx, &repo, data, dataDetail)

	return
}

func insertProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data *entity.ProgramDonasiRutinEntity, dataDetail *entity.ProgramDonasiRutinDetailEntity) (response entity.ProgramDonasiRutinEntity, err error) {
	err = repo.Insert(ctx, data)

	// Insert Detail
	dataDetail.IDPPCPProgramDonasiRutin = data.IDPPCPProgramDonasiRutin
	dataDetail.Tag = data.Tag
	if errDetail := repo.InsertDetail(ctx, dataDetail); errDetail != nil {
		log.Println("ERR Insert Detail: ", errDetail)
	}

	response, _ = GetProgramDonasiRutin(ctx, repo, data.IDPPCPProgramDonasiRutin)

	return
}
