package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/transaction/domain/entity"
	"pemuda-peduli/src/transaction/domain/interfaces"
)

func FindTransaction(ctx context.Context, repo interfaces.ITransactionRepository, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func FindMyTransaction(ctx context.Context, repo interfaces.ITransactionRepository, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error) {
	response, count, err = repo.FindMyTransaction(ctx, data)
	return
}

func GetTransaction(ctx context.Context, repo interfaces.ITransactionRepository, id string) (response entity.TransactionEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
