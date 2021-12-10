package interfaces

import (
	"context"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/entity"
)

type IProgramDonasiFundraiserRepository interface {
	Insert(ctx context.Context, data *entity.ProgramDonasiFundraiserEntity) (err error)
	Update(ctx context.Context, data *entity.ProgramDonasiFundraiserEntity) (err error)
	Find(ctx context.Context, data entity.ProgramDonasiFundraiserQueryEntity) (response []entity.ProgramDonasiFundraiserEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.ProgramDonasiFundraiserEntity, err error)
	GetSeo(ctx context.Context, seoURL string) (response entity.ProgramDonasiFundraiserEntity, err error)
}
