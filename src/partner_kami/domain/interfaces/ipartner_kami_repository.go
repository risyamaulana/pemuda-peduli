package interfaces

import (
	"context"
	"pemuda-peduli/src/partner_kami/domain/entity"
)

type IPartnerKamiRepository interface {
	Insert(ctx context.Context, data *entity.PartnerKamiEntity) (err error)

	Update(ctx context.Context, data entity.PartnerKamiEntity, id string) (response entity.PartnerKamiEntity, err error)

	Find(ctx context.Context, data *entity.PartnerKamiQueryEntity) (company []entity.PartnerKamiEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.PartnerKamiEntity, err error)
}
