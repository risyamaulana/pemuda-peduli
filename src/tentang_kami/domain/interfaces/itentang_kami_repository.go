package interfaces

import (
	"context"
	"pemuda-peduli/src/tentang_kami/domain/entity"
)

type ITentangKamiRepository interface {
	Insert(ctx context.Context, data *entity.TentangKamiEntity) (err error)

	Update(ctx context.Context, data entity.TentangKamiEntity, id string) (response entity.TentangKamiEntity, err error)

	Find(ctx context.Context, data *entity.TentangKamiQueryEntity) (company []entity.TentangKamiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.TentangKamiEntity, err error)
}
