package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/banner/domain/entity"
	"pemuda-peduli/src/banner/domain/interfaces"
)

func FindBanner(ctx context.Context, repo interfaces.IBannerRepository, data *entity.BannerQueryEntity) (response []entity.BannerEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetBanner(ctx context.Context, repo interfaces.IBannerRepository, id string) (response entity.BannerEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
