package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/penggalang_dana/domain/entity"
	"pemuda-peduli/src/penggalang_dana/domain/interfaces"
	"pemuda-peduli/src/penggalang_dana/infrastructure/repository"
	"time"
)

func EditPenggalangDana(ctx context.Context, db *db.ConnectTo, data entity.PenggalangDanaEntity) (response entity.PenggalangDanaEntity, err error) {
	// Repo
	repo := repository.NewPenggalangDanaRepository(db)

	// Check available data
	checkData, errCheckData := getPenggalangDana(ctx, &repo, data.IDPPCPPenggalangDana)
	if errCheckData != nil {
		err = errors.New("Failed, data not found")
		return
	}

	data.ID = checkData.ID
	data.CreatedAt = checkData.CreatedAt
	data.IsVerified = checkData.IsVerified
	data.IsDeleted = checkData.IsDeleted

	response, err = updatePenggalangDana(ctx, &repo, data)
	return
}

func DeletePenggalangDana(ctx context.Context, db *db.ConnectTo, id string) (response entity.PenggalangDanaEntity, err error) {
	currentTime := time.Now().UTC()
	// Repo
	repo := repository.NewPenggalangDanaRepository(db)

	// Check available data
	checkData, errCheckData := getPenggalangDana(ctx, &repo, id)
	if errCheckData != nil {
		err = errors.New("Failed, data not found")
		return
	}

	checkData.IsDeleted = true
	checkData.UpdatedAt = &currentTime

	response, err = updatePenggalangDana(ctx, &repo, checkData)
	return
}

func ToogleVerified(ctx context.Context, db *db.ConnectTo, id string) (response entity.PenggalangDanaEntity, err error) {
	currentTime := time.Now().UTC()
	// Repo
	repo := repository.NewPenggalangDanaRepository(db)

	// Check available data
	checkData, errCheckData := getPenggalangDana(ctx, &repo, id)
	if errCheckData != nil {
		err = errors.New("Failed, data not found")
		return
	}

	if checkData.IsVerified {
		checkData.IsVerified = false
	} else {
		checkData.IsVerified = true
	}

	checkData.UpdatedAt = &currentTime

	response, err = updatePenggalangDana(ctx, &repo, checkData)
	return
}

func updatePenggalangDana(ctx context.Context, repo interfaces.IPenggalangDanaRepository, data entity.PenggalangDanaEntity) (response entity.PenggalangDanaEntity, err error) {
	response, err = repo.Update(ctx, data)
	return
}
