package domain

import (
	"context"
	"pemuda-peduli/src/achievement/domain/entity"
	"pemuda-peduli/src/achievement/domain/interfaces"
)

func CreateAchievement(ctx context.Context, repo interfaces.IAchievementRepository, data *entity.AchievementEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
