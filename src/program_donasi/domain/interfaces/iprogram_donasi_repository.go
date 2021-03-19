package interfaces

import (
	"context"
	"pemuda-peduli/src/program_donasi/domain/entity"
)

type IProgramDonasiRepository interface {
	Insert(ctx context.Context, data *entity.ProgramDonasiEntity) (err error)

	Update(ctx context.Context, data entity.ProgramDonasiEntity, id string) (response entity.ProgramDonasiEntity, err error)

	Find(ctx context.Context, data *entity.ProgramDonasiQueryEntity) (company []entity.ProgramDonasiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramDonasiEntity, err error)
}
