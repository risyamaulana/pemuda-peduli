package interfaces

import (
	"context"
	"pemuda-peduli/src/program_incidental/domain/entity"
)

type IProgramIncidentalRepository interface {
	Insert(ctx context.Context, data *entity.ProgramIncidentalEntity) (err error)
	Update(ctx context.Context, data entity.ProgramIncidentalEntity, id string) (response entity.ProgramIncidentalEntity, err error)
	Find(ctx context.Context, data *entity.ProgramIncidentalQueryEntity) (response []entity.ProgramIncidentalEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramIncidentalEntity, err error)
}
