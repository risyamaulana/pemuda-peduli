package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/partner_kami/domain/entity"
	"pemuda-peduli/src/partner_kami/domain/interfaces"
)

func FindPartnerKami(ctx context.Context, repo interfaces.IPartnerKamiRepository, data *entity.PartnerKamiQueryEntity) (response []entity.PartnerKamiEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetPartnerKami(ctx context.Context, repo interfaces.IPartnerKamiRepository, id string) (response entity.PartnerKamiEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
