package interfaces

import (
	"context"
	"pemuda-peduli/src/program_kami/domain/entity"
)

type IProgramKamiRepository interface {
	Insert(ctx context.Context, data *entity.ProgramKamiEntity) (err error)

	Update(ctx context.Context, data entity.ProgramKamiEntity, id string) (response entity.ProgramKamiEntity, err error)

	Find(ctx context.Context, data *entity.ProgramKamiQueryEntity) (company []entity.ProgramKamiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramKamiEntity, err error)
}
