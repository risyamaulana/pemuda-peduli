package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
)

func FindProgramDonasiRutin(ctx context.Context, repo interfaces.IProgramDonasiRutinRepository, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
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
