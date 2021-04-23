package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/achievement/domain/entity"
	"pemuda-peduli/src/achievement/domain/interfaces"
)

func FindAchievement(ctx context.Context, repo interfaces.IAchievementRepository, data *entity.AchievementQueryEntity) (response []entity.AchievementEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetAchievement(ctx context.Context, repo interfaces.IAchievementRepository, id string) (response entity.AchievementEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
