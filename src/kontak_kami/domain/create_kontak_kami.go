package domain

import (
	"context"
	"pemuda-peduli/src/kontak_kami/domain/entity"
	"pemuda-peduli/src/kontak_kami/domain/interfaces"
)

func CreateKontakKami(ctx context.Context, repo interfaces.IKontakKamiRepository, data *entity.KontakKamiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
