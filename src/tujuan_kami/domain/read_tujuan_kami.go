package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/tujuan_kami/domain/entity"
	"pemuda-peduli/src/tujuan_kami/domain/interfaces"
)

func FindTujuanKami(ctx context.Context, repo interfaces.ITujuanKamiRepository, data *entity.TujuanKamiQueryEntity) (response []entity.TujuanKamiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetTujuanKami(ctx context.Context, repo interfaces.ITujuanKamiRepository, id string) (response entity.TujuanKamiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
