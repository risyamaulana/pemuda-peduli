package interfaces

import (
	"context"
	"pemuda-peduli/src/kontak_kami/domain/entity"
)

type IKontakKamiRepository interface {
	Insert(ctx context.Context, data *entity.KontakKamiEntity) (err error)

	Update(ctx context.Context, data entity.KontakKamiEntity, id string) (response entity.KontakKamiEntity, err error)

	Find(ctx context.Context, data *entity.KontakKamiQueryEntity) (company []entity.KontakKamiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.KontakKamiEntity, err error)
}
