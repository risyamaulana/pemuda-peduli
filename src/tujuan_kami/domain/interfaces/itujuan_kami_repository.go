package interfaces

import (
	"context"
	"pemuda-peduli/src/tujuan_kami/domain/entity"
)

type ITujuanKamiRepository interface {
	Insert(ctx context.Context, data *entity.TujuanKamiEntity) (err error)

	Update(ctx context.Context, data entity.TujuanKamiEntity, id string) (response entity.TujuanKamiEntity, err error)

	Find(ctx context.Context, data *entity.TujuanKamiQueryEntity) (company []entity.TujuanKamiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.TujuanKamiEntity, err error)
}
