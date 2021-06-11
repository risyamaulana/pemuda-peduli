package interfaces

import (
	"context"
	"pemuda-peduli/src/hubungi_kami/domain/entity"
)

type IHubungiKamiRepository interface {
	Insert(ctx context.Context, data *entity.HubungiKamiEntity) (err error)
	Update(ctx context.Context, data entity.HubungiKamiEntity, id string) (response entity.HubungiKamiEntity, err error)
	Find(ctx context.Context, data *entity.HubungiKamiQueryEntity) (response []entity.HubungiKamiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.HubungiKamiEntity, err error)
}
