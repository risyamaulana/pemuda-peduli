package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/testimoni/domain/entity"
	"pemuda-peduli/src/testimoni/domain/interfaces"
)

func FindTestimoni(ctx context.Context, repo interfaces.ITestimoniRepository, data *entity.TestimoniQueryEntity) (response []entity.TestimoniEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetTestimoni(ctx context.Context, repo interfaces.ITestimoniRepository, id string) (response entity.TestimoniEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
