package interfaces

import (
	"context"
	"pemuda-peduli/src/program_donasi/domain/entity"
)

type IProgramDonasiRepository interface {
	Insert(ctx context.Context, data *entity.ProgramDonasiEntity) (err error)
	InsertDetail(ctx context.Context, data *entity.ProgramDonasiDetailEntity) (err error)

	Update(ctx context.Context, data entity.ProgramDonasiEntity, id string) (response entity.ProgramDonasiEntity, err error)
	UpdateDetail(ctx context.Context, data entity.ProgramDonasiDetailEntity, id string) (response entity.ProgramDonasiDetailEntity, err error)

	Find(ctx context.Context, data *entity.ProgramDonasiQueryEntity) (company []entity.ProgramDonasiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramDonasiEntity, err error)

	GetBySeo(ctx context.Context, seo string) (response entity.ProgramDonasiEntity, err error)
	GetDetail(ctx context.Context, id string) (response entity.ProgramDonasiDetailEntity, err error)
}
