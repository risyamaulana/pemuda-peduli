package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/program_incidental/domain/entity"
	"pemuda-peduli/src/program_incidental/domain/interfaces"
)

func FindProgramIncidental(ctx context.Context, repo interfaces.IProgramIncidentalRepository, data *entity.ProgramIncidentalQueryEntity) (response []entity.ProgramIncidentalEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetProgramIncidental(ctx context.Context, repo interfaces.IProgramIncidentalRepository, id string) (response entity.ProgramIncidentalEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
