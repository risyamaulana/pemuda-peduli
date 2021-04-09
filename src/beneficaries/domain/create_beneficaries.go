package domain

import (
	"context"
	"pemuda-peduli/src/beneficaries/domain/entity"
	"pemuda-peduli/src/beneficaries/domain/interfaces"
)

func CreateBeneficaries(ctx context.Context, repo interfaces.IBeneficariesRepository, data *entity.BeneficariesEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
