package domain

import (
	"context"
	"pemuda-peduli/src/team/domain/entity"
	"pemuda-peduli/src/team/domain/interfaces"
)

func CreateTeam(ctx context.Context, repo interfaces.ITeamRepository, data *entity.TeamEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
