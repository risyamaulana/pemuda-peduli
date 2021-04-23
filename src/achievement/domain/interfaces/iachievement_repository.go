package interfaces

import (
	"context"
	"pemuda-peduli/src/achievement/domain/entity"
)

type IAchievementRepository interface {
	Insert(ctx context.Context, data *entity.AchievementEntity) (err error)

	Update(ctx context.Context, data entity.AchievementEntity, id string) (response entity.AchievementEntity, err error)

	Find(ctx context.Context, data *entity.AchievementQueryEntity) (company []entity.AchievementEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.AchievementEntity, err error)
}
