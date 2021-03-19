package domain

import (
	"context"
	"pemuda-peduli/src/partner_kami/domain/entity"
	"pemuda-peduli/src/partner_kami/domain/interfaces"
)

func CreatePartnerKami(ctx context.Context, repo interfaces.IPartnerKamiRepository, data *entity.PartnerKamiEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
