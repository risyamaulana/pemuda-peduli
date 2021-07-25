package domain

import (
	"context"
	"pemuda-peduli/src/qris/domain/entity"
	"pemuda-peduli/src/qris/domain/interfaces"
)

func CreateQris(ctx context.Context, repo interfaces.IQrisRepository, data *entity.QrisEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
