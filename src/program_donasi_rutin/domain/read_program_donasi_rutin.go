package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"
)

func FindProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func FindProgramDonasiNews(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinNewsEntity, count int, err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	response, count, err = repo.FindNews(ctx, data)

	return
}

func GetProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, id string) (response entity.ProgramDonasiRutinEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	return
}

func GetProgramDonasiNews(ctx context.Context, db *db.ConnectTo, id int64) (response entity.ProgramDonasiRutinNewsEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiRutinRepository(db)

	response, err = repo.GetNews(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}

	return
}
