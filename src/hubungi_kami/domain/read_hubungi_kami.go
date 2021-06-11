package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/hubungi_kami/domain/entity"
	"pemuda-peduli/src/hubungi_kami/domain/interfaces"
)

func FindHubungiKami(ctx context.Context, repo interfaces.IHubungiKamiRepository, data *entity.HubungiKamiQueryEntity) (response []entity.HubungiKamiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetHubungiKami(ctx context.Context, repo interfaces.IHubungiKamiRepository, id string) (response entity.HubungiKamiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
