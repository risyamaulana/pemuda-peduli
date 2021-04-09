package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/beneficaries/domain/entity"
	"pemuda-peduli/src/beneficaries/domain/interfaces"
)

func FindBeneficaries(ctx context.Context, repo interfaces.IBeneficariesRepository, data *entity.BeneficariesQueryEntity) (response []entity.BeneficariesEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetBeneficaries(ctx context.Context, repo interfaces.IBeneficariesRepository, id string) (response entity.BeneficariesEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
