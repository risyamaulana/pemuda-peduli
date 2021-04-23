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

	InsertDetail(ctx context.Context, data *entity.ProgramKamiDetailEntity) (err error)
	UpdateDetail(ctx context.Context, data entity.ProgramKamiDetailEntity, id string) (response entity.ProgramKamiDetailEntity, err error)

	GetDetail(ctx context.Context, id string) (response entity.ProgramKamiDetailEntity, err error)
}
