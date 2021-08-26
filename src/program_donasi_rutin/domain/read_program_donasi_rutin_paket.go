package domain

import (
	"context"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
)

func FindProgramDonasiRutinPaket(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiRutinQueryEntity) (responses []entity.ProgramDonasiRutinPaketEntity, count int, err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	responses, count, err = findProgramDonasiRutinPaket(ctx, &repo, data)
	return
}

func GetProgramDonasiRutinPaket(ctx context.Context, db *db.ConnectTo, id string) (resp entity.ProgramDonasiRutinPaketEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	resp, err = getProgramDonasiRutinPaket(ctx, &repo, id)
	return
}

func findProgramDonasiRutinPaket(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data *entity.ProgramDonasiRutinQueryEntity) (responses []entity.ProgramDonasiRutinPaketEntity, count int, err error) {
	responses, count, err = repo.FindPaket(ctx, data)
	return
}

func getProgramDonasiRutinPaket(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (resp entity.ProgramDonasiRutinPaketEntity, err error) {
	resp, err = repo.GetPaket(ctx, id)
	return
}
