package domain

import (
	"context"
	"pemuda-peduli/src/banner/domain/entity"
	"pemuda-peduli/src/banner/domain/interfaces"
)

func CreateBanner(ctx context.Context, repo interfaces.IBannerRepository, data *entity.BannerEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
