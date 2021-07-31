package interfaces

import (
	"context"
	"pemuda-peduli/src/transaction/domain/entity"
)

type ITransactionRepository interface {
	Insert(ctx context.Context, data *entity.TransactionEntity) (err error)
	Update(ctx context.Context, data entity.TransactionEntity, id string) (response entity.TransactionEntity, err error)
	Find(ctx context.Context, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error)
	FindMyTransaction(ctx context.Context, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.TransactionEntity, err error)
}
