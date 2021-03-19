package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/team/domain/entity"
	"pemuda-peduli/src/team/domain/interfaces"
)

func FindTeam(ctx context.Context, repo interfaces.ITeamRepository, data *entity.TeamQueryEntity) (response []entity.TeamEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetTeam(ctx context.Context, repo interfaces.ITeamRepository, id string) (response entity.TeamEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
