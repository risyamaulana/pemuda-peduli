package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/kontak_kami/domain/entity"
	"pemuda-peduli/src/kontak_kami/domain/interfaces"
)

func FindKontakKami(ctx context.Context, repo interfaces.IKontakKamiRepository, data *entity.KontakKamiQueryEntity) (response []entity.KontakKamiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetKontakKami(ctx context.Context, repo interfaces.IKontakKamiRepository, id string) (response entity.KontakKamiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
