package interfaces

import (
	"context"
	"pemuda-peduli/src/team_flag/domain/entity"
)

type ITeamFlagRepository interface {
	Find(ctx context.Context) (response []entity.TeamFlagEntity, err error)
	Get(ctx context.Context, id int64) (response entity.TeamFlagEntity, err error)
}
