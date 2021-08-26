package domain

import (
	"context"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/penggalang_dana/domain/entity"
	"pemuda-peduli/src/penggalang_dana/domain/interfaces"
	"pemuda-peduli/src/penggalang_dana/infrastructure/repository"
)

func CreatePenggalangDana(ctx context.Context, db *db.ConnectTo, data *entity.PenggalangDanaEntity) (err error) {
	// Repo
	repo := repository.NewPenggalangDanaRepository(db)

	err = insertPenggalangDana(ctx, &repo, data)
	return
}

func insertPenggalangDana(ctx context.Context, repo interfaces.IPenggalangDanaRepository, data *entity.PenggalangDanaEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
