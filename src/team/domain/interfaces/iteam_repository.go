package interfaces

import (
	"context"
	"pemuda-peduli/src/team/domain/entity"
)

type ITeamRepository interface {
	Insert(ctx context.Context, data *entity.TeamEntity) (err error)

	Update(ctx context.Context, data entity.TeamEntity, id string) (response entity.TeamEntity, err error)

	Find(ctx context.Context, data *entity.TeamQueryEntity) (company []entity.TeamEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.TeamEntity, err error)
}
