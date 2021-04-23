package interfaces

import (
	"context"
	"pemuda-peduli/src/berita/domain/entity"
)

type IBeritaRepository interface {
	Insert(ctx context.Context, data *entity.BeritaEntity) (err error)

	Update(ctx context.Context, data entity.BeritaEntity, id string) (response entity.BeritaEntity, err error)

	Find(ctx context.Context, data *entity.BeritaQueryEntity) (company []entity.BeritaEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.BeritaEntity, err error)

	InsertDetail(ctx context.Context, data *entity.BeritaDetailEntity) (err error)
	UpdateDetail(ctx context.Context, data entity.BeritaDetailEntity, id string) (response entity.BeritaDetailEntity, err error)
	GetDetail(ctx context.Context, id string) (response entity.BeritaDetailEntity, err error)
}
