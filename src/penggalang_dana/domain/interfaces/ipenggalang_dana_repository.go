package interfaces

import (
	"context"
	"pemuda-peduli/src/penggalang_dana/domain/entity"
)

type IPenggalangDanaRepository interface {
	Insert(ctx context.Context, data *entity.PenggalangDanaEntity) (err error)
	Update(ctx context.Context, data entity.PenggalangDanaEntity) (response entity.PenggalangDanaEntity, err error)
	Find(ctx context.Context, data *entity.PenggalangDanaQueryEntity) (response []entity.PenggalangDanaEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.PenggalangDanaEntity, err error)
}
