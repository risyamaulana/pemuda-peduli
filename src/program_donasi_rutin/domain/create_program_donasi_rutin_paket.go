package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
	"strings"
)

func CreateProgramDonasiRutinPaket(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiRutinPaketEntity) (err error) {
	// repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	// Check data Donasi Rutin
	_, errCheckData := GetProgramDonasiRutin(ctx, &repo, data.IDPPCPProgramDonasiRutin)
	if errCheckData != nil {
		err = errors.New("Failed, data donasi not found")
		return
	}
	// Check SEO URL
	if data.SeoURL == "" {
		data.SeoURL = strings.ToLower(strings.ReplaceAll(data.PaketName, " ", "-"))
	}

	err = insertProgramDonasiRutinPaket(ctx, &repo, data)
	return
}

func insertProgramDonasiRutinPaket(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data *entity.ProgramDonasiRutinPaketEntity) (err error) {
	err = repo.InsertPaket(ctx, data)
	return
}
