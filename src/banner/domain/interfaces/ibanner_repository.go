package interfaces

import (
	"context"
	"pemuda-peduli/src/banner/domain/entity"
)

type IBannerRepository interface {
	Insert(ctx context.Context, data *entity.BannerEntity) (err error)

	Update(ctx context.Context, data entity.BannerEntity, id string) (response entity.BannerEntity, err error)

	Find(ctx context.Context, data *entity.BannerQueryEntity) (company []entity.BannerEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.BannerEntity, err error)
}
