package interfaces

import (
	"context"
	"pemuda-peduli/src/program_donasi_kategori/domain/entity"
)

type IProgramDonasiKategoriRepository interface {
	Insert(ctx context.Context, data *entity.ProgramDonasiKategoriEntity) (err error)
	Update(ctx context.Context, data entity.ProgramDonasiKategoriEntity, id string) (response entity.ProgramDonasiKategoriEntity, err error)
	Delete(ctx context.Context, id string) (err error)
	Find(ctx context.Context, data *entity.ProgramDonasiKategoriQueryEntity) (response []entity.ProgramDonasiKategoriEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramDonasiKategoriEntity, err error)
}
