package domain

import (
	"context"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/penggalang_dana/domain/entity"
	"pemuda-peduli/src/penggalang_dana/domain/interfaces"
	"pemuda-peduli/src/penggalang_dana/infrastructure/repository"
)

func FindPenggalangDana(ctx context.Context, db *db.ConnectTo, data entity.PenggalangDanaQueryEntity) (responses []entity.PenggalangDanaEntity, count int, err error) {
	// Repo
	repo := repository.NewPenggalangDanaRepository(db)

	responses, count, err = findPenggalangDana(ctx, &repo, data)
	return
}

func GetPenggalangDana(ctx context.Context, db *db.ConnectTo, id string) (response entity.PenggalangDanaEntity, err error) {
	// Repo
	repo := repository.NewPenggalangDanaRepository(db)

	response, err = getPenggalangDana(ctx, &repo, id)
	return
}

func findPenggalangDana(ctx context.Context, repo interfaces.IPenggalangDanaRepository, data entity.PenggalangDanaQueryEntity) (responses []entity.PenggalangDanaEntity, count int, err error) {
	responses, count, err = repo.Find(ctx, &data)
	return
}

func getPenggalangDana(ctx context.Context, repo interfaces.IPenggalangDanaRepository, id string) (response entity.PenggalangDanaEntity, err error) {
	response, err = repo.Get(ctx, id)
	return
}
