package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/qris/domain/entity"
	"pemuda-peduli/src/qris/domain/interfaces"
)

func FindQris(ctx context.Context, repo interfaces.IQrisRepository, data *entity.QrisQueryEntity) (response []entity.QrisEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetQris(ctx context.Context, repo interfaces.IQrisRepository, id string) (response entity.QrisEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
