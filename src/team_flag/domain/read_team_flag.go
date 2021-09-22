package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/team_flag/domain/entity"
	"pemuda-peduli/src/team_flag/domain/interfaces"
)

func ListTeamFlag(ctx context.Context, repo interfaces.ITeamFlagRepository) (responses []entity.TeamFlagEntity, err error) {
	responses, err = repo.Find(ctx)
	return
}

func GetTeamFlag(ctx context.Context, repo interfaces.ITeamFlagRepository, id int64) (response entity.TeamFlagEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
