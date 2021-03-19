package interfaces

import (
	"context"
	"pemuda-peduli/src/testimoni/domain/entity"
)

type ITestimoniRepository interface {
	Insert(ctx context.Context, data *entity.TestimoniEntity) (err error)

	Update(ctx context.Context, data entity.TestimoniEntity, id string) (response entity.TestimoniEntity, err error)

	Find(ctx context.Context, data *entity.TestimoniQueryEntity) (company []entity.TestimoniEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.TestimoniEntity, err error)
}
