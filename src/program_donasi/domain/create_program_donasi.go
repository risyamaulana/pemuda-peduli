package domain

import (
	"context"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/domain/interfaces"
)

func CreateProgramDonasi(ctx context.Context, repo interfaces.IProgramDonasiRepository, data *entity.ProgramDonasiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
