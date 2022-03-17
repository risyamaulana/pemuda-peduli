package interfaces

import (
	"context"
	"pemuda-peduli/src/kategori_program_donasi/domain/entity"
)

type IKategoriProgramDonasiRepository interface {
	Insert(ctx context.Context, data *entity.KategoriProgramDonasiEntity) (err error)
	Update(ctx context.Context, data entity.KategoriProgramDonasiEntity, id string) (response entity.KategoriProgramDonasiEntity, err error)
	Find(ctx context.Context, data *entity.KategoriProgramDonasiQueryEntity) (response []entity.KategoriProgramDonasiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.KategoriProgramDonasiEntity, err error)
}
