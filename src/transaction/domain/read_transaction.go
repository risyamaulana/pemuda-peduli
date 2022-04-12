package domain

import (
	"context"
	"errors"
	"fmt"
	"pemuda-peduli/src/transaction/domain/entity"
	"pemuda-peduli/src/transaction/domain/interfaces"
	"time"
)

func FindTransaction(ctx context.Context, repo interfaces.ITransactionRepository, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func FindMyTransaction(ctx context.Context, repo interfaces.ITransactionRepository, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error) {
	response, count, err = repo.FindMyTransaction(ctx, data)
	return
}

func FindRutinTransaction(ctx context.Context, repo interfaces.ITransactionRepository) (response []entity.TransactionEntity, err error) {
	responseData, err := repo.FindRutinTransaction(ctx)
	if err != nil {
		return
	}

	for _, transactionEntity := range responseData {
		hourDate := int(time.Since(transactionEntity.CreatedAt).Hours())

		if hourDate > 0 {
			fmt.Println(hourDate)
			dayDate := int(hourDate / 24)
			fmt.Println(dayDate)
			fmt.Println("================")
			if dayDate%30 == 0 {
				response = append(response, transactionEntity)
			}
		}

	}
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
