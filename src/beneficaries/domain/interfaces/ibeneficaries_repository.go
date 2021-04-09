package interfaces

import (
	"context"
	"pemuda-peduli/src/beneficaries/domain/entity"
)

type IBeneficariesRepository interface {
	Insert(ctx context.Context, data *entity.BeneficariesEntity) (err error)

	Update(ctx context.Context, data entity.BeneficariesEntity, id string) (response entity.BeneficariesEntity, err error)

	Find(ctx context.Context, data *entity.BeneficariesQueryEntity) (company []entity.BeneficariesEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.BeneficariesEntity, err error)
}
