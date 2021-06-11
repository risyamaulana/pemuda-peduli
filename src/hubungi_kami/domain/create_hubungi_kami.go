package domain

import (
	"context"
	"pemuda-peduli/src/hubungi_kami/domain/entity"
	"pemuda-peduli/src/hubungi_kami/domain/interfaces"
)

func CreateHubungiKami(ctx context.Context, repo interfaces.IHubungiKamiRepository, data *entity.HubungiKamiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
