package domain

import (
	"context"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
	"strings"
	// kategoriDom "pemuda-peduli/src/program_donasi_kategori/domain"
	// kategoriRep "pemuda-peduli/src/program_donasi_kategori/infrastructure/repository"
)

func CreateProgramDonasiRutin(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiRutinEntity) (response entity.ProgramDonasiRutinEntity, err error) {
	repo := repository.NewProgramDonasiRutinRepository(db)
	// kategoriRepo := kategoriRep.NewProgramDonasiKategoriRepository(db)

	// Check Kategori
	// kategoriData, err := kategoriDom.GetProgramDonasiKategori(ctx, &kategoriRepo, data.IDPPCPProgramDonasiKategori)
	// if err != nil {
	// 	err = errors.New("Failed, kategori not found")
	// 	return
	// }
	// data.KategoriName = kategoriData.KategoriName

	// Check SEO URL
	if data.SEOURL == "" {
		data.SEOURL = strings.ToLower(strings.ReplaceAll(data.Title, " ", "-"))
	}

	response, err = insertProgramDonasiRutin(ctx, &repo, data)

	return
}

func insertProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data *entity.ProgramDonasiRutinEntity) (response entity.ProgramDonasiRutinEntity, err error) {
	err = repo.Insert(ctx, data)
	if err != nil {
		return
	}

	response, _ = GetProgramDonasiRutin(ctx, repo, data.IDPPCPProgramDonasiRutin)

	return
}

func CreateProgramDonasiNews(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiRutinNewsEntity) (err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	err = repo.InsertNews(ctx, data)
	return
}
