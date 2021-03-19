package domain

import (
	"context"
	"pemuda-peduli/src/program_kami/domain/entity"
	"pemuda-peduli/src/program_kami/domain/interfaces"
)

func CreateProgramKami(ctx context.Context, repo interfaces.IProgramKamiRepository, data *entity.ProgramKamiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
