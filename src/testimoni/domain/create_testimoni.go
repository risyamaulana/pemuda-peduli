package domain

import (
	"context"
	"pemuda-peduli/src/testimoni/domain/entity"
	"pemuda-peduli/src/testimoni/domain/interfaces"
)

func CreateTestimoni(ctx context.Context, repo interfaces.ITestimoniRepository, data *entity.TestimoniEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
