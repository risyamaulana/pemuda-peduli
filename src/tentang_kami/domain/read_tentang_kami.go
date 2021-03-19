package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/tentang_kami/domain/entity"
	"pemuda-peduli/src/tentang_kami/domain/interfaces"
)

func FindTentangKami(ctx context.Context, repo interfaces.ITentangKamiRepository, data *entity.TentangKamiQueryEntity) (response []entity.TentangKamiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetTentangKami(ctx context.Context, repo interfaces.ITentangKamiRepository, id string) (response entity.TentangKamiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
