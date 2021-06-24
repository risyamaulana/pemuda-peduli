package interfaces

import (
	"context"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
)

type IProgramDonasiRutinRepository interface {
	Insert(ctx context.Context, data *entity.ProgramDonasiRutinEntity) (err error)
	InsertDetail(ctx context.Context, data *entity.ProgramDonasiRutinDetailEntity) (err error)
	Update(ctx context.Context, data entity.ProgramDonasiRutinEntity, id string) (response entity.ProgramDonasiRutinEntity, err error)
	UpdateDetail(ctx context.Context, data entity.ProgramDonasiRutinDetailEntity, id string) (response entity.ProgramDonasiRutinDetailEntity, err error)
	Find(ctx context.Context, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramDonasiRutinEntity, err error)
	GetDetail(ctx context.Context, id string) (response entity.ProgramDonasiRutinDetailEntity, err error)
}
