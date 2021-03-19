package domain

import (
	"context"
	"pemuda-peduli/src/tentang_kami/domain/entity"
	"pemuda-peduli/src/tentang_kami/domain/interfaces"
)

func CreateTentangKami(ctx context.Context, repo interfaces.ITentangKamiRepository, data *entity.TentangKamiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
