package interfaces

import (
	"context"
	"pemuda-peduli/src/qris/domain/entity"
)

type IQrisRepository interface {
	Insert(ctx context.Context, data *entity.QrisEntity) (err error)
	Update(ctx context.Context, data entity.QrisEntity, id string) (response entity.QrisEntity, err error)
	Find(ctx context.Context, data *entity.QrisQueryEntity) (response []entity.QrisEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.QrisEntity, err error)
}
