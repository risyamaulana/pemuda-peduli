package domain

import (
	"context"
	"pemuda-peduli/src/tujuan_kami/domain/entity"
	"pemuda-peduli/src/tujuan_kami/domain/interfaces"
)

func CreateTujuanKami(ctx context.Context, repo interfaces.ITujuanKamiRepository, data *entity.TujuanKamiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
