package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/program_kami/domain/entity"
	"pemuda-peduli/src/program_kami/domain/interfaces"
)

func FindProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, data *entity.ProgramKamiQueryEntity) (response []entity.ProgramKamiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, id string) (response entity.ProgramKamiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
