package interfaces

import (
	"context"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
)

type IProgramDonasiRutinRepository interface {
	Insert(ctx context.Context, data *entity.ProgramDonasiRutinEntity) (err error)
	Update(ctx context.Context, data entity.ProgramDonasiRutinEntity, id string) (response entity.ProgramDonasiRutinEntity, err error)
	Find(ctx context.Context, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramDonasiRutinEntity, err error)

	InsertPaket(ctx context.Context, data *entity.ProgramDonasiRutinPaketEntity) (err error)
	UpdatePaket(ctx context.Context, data entity.ProgramDonasiRutinPaketEntity, id string) (response entity.ProgramDonasiRutinPaketEntity, err error)
	FindPaket(ctx context.Context, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinPaketEntity, count int, err error)
	GetPaket(ctx context.Context, id string) (response entity.ProgramDonasiRutinPaketEntity, err error)
}
