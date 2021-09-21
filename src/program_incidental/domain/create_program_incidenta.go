package domain

import (
	"context"
	"pemuda-peduli/src/program_incidental/domain/entity"
	"pemuda-peduli/src/program_incidental/domain/interfaces"
)

func CreateProgramIncidental(ctx context.Context, repo interfaces.IProgramIncidentalRepository, data *entity.ProgramIncidentalEntity) (response entity.ProgramIncidentalEntity, err error) {
	err = repo.Insert(ctx, data)

	response, _ = GetProgramIncidental(ctx, repo, data.IDPPCPProgramIncidental)

	return
}
