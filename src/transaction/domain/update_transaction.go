package domain

import (
	"context"
	"errors"

	"pemuda-peduli/src/common/infrastructure/db"

	"pemuda-peduli/src/transaction/common/constants"
	"pemuda-peduli/src/transaction/domain/entity"
	"pemuda-peduli/src/transaction/domain/interfaces"
	"pemuda-peduli/src/transaction/infrastructure/repository"

	adminDom "pemuda-peduli/src/admin_user/domain"

	donasiRutinDom "pemuda-peduli/src/program_donasi_rutin/domain"
	donasiRutinRep "pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"

	donasiDom "pemuda-peduli/src/program_donasi/domain"
	donasiRep "pemuda-peduli/src/program_donasi/infrastructure/repository"

	"time"
)

func UploadPaymentReceipt(ctx context.Context, db *db.ConnectTo, transactionID string, imagePaymentURL string) (response entity.TransactionEntity, err error) {
	currentDate := time.Now()
	// Repo
	repo := repository.NewTransactionRepository(db)

	// Get data
	data, err := GetTransaction(ctx, &repo, transactionID)
	if err != nil {
		err = errors.New("Failed, transaction data not found")
	}

	// Check user transaction data
	if data.UserID != ctx.Value("user_id").(string) {
		err = errors.New("Failed, user data is unauthorized for this transaction")
		return
	}

	data.ImagePaymentURL = imagePaymentURL
	data.PaidAt = &currentDate
	data.Status = constants.StatusNeedApproval
	data.UpdatedAt = &currentDate

	response, err = updateTransaction(ctx, &repo, data)
	return
}

func AppliedPayment(ctx context.Context, db *db.ConnectTo, transactionID string) (response entity.TransactionEntity, err error) {
	currentDate := time.Now()
	// Repo
	repo := repository.NewTransactionRepository(db)
	donasiRepo := donasiRep.NewProgramDonasiRepository(db)
	donasiRutinRepo := donasiRutinRep.NewProgramDonasiRutinRepository(db)

	// Get data
	data, err := GetTransaction(ctx, &repo, transactionID)
	if err != nil {
		err = errors.New("Failed, transaction data not found")
	}

	// Check data Approve admin
	adminData, errAdminData := adminDom.GetAdminUser(ctx, db, ctx.Value("user_id").(string))
	if errAdminData != nil {
		err = errors.New("Failed, this user is unauthorized for this action")
		return
	}

	// Collect cash
	if data.IsRutin {
		idDonasiRutin := data.IDPPCPProgramDonasiRutin
		_, errCollect := donasiRutinDom.UpdateDonationCollect(ctx, &donasiRutinRepo, idDonasiRutin, data.Amount)
		if errCollect != nil {
			err = errors.New("Failed Applied payment, donation id not found")
			return
		}
	} else {
		idDonasi := data.IDPPCPProgramDonasi
		_, errCollect := donasiDom.UpdateDonationCollect(ctx, &donasiRepo, idDonasi, data.Amount)
		if errCollect != nil {
			err = errors.New("Failed Applied payment, donation id not found")
			return
		}
	}

	data.ApprovedBy = adminData.Email
	data.ApprovedAt = &currentDate

	data.Status = constants.StatusPaid
	data.UpdatedAt = &currentDate

	response, err = updateTransaction(ctx, &repo, data)
	return
}

func DeclinePayment(ctx context.Context, db *db.ConnectTo, transactionID string) (response entity.TransactionEntity, err error) {
	currentDate := time.Now()
	// Repo
	repo := repository.NewTransactionRepository(db)

	// Get data
	data, err := GetTransaction(ctx, &repo, transactionID)
	if err != nil {
		err = errors.New("Failed, transaction data not found")
	}

	// Check data Approve admin
	adminData, errAdminData := adminDom.GetAdminUser(ctx, db, ctx.Value("user_id").(string))
	if errAdminData != nil {
		err = errors.New("Failed, this user is unauthorized for this action")
		return
	}
	data.ApprovedBy = adminData.Email
	data.ApprovedAt = &currentDate

	data.Status = constants.StatusDecline
	data.UpdatedAt = &currentDate

	response, err = updateTransaction(ctx, &repo, data)
	return
}

func CancelPayment(ctx context.Context, db *db.ConnectTo, transactionID string) (response entity.TransactionEntity, err error) {
	currentDate := time.Now()
	// Repo
	repo := repository.NewTransactionRepository(db)

	// Get data
	data, err := GetTransaction(ctx, &repo, transactionID)
	if err != nil {
		err = errors.New("Failed, transaction data not found")
	}

	// Check user transaction data
	if data.UserID != ctx.Value("user_id").(string) {
		err = errors.New("Failed, user data is unauthorized for this transaction")
		return
	}

	data.Status = constants.StatusCancel
	data.UpdatedAt = &currentDate

	response, err = updateTransaction(ctx, &repo, data)
	return
}

func updateTransaction(ctx context.Context, repo interfaces.ITransactionRepository, data entity.TransactionEntity) (response entity.TransactionEntity, err error) {
	response, err = repo.Update(ctx, data, data.IDPPTransaction)
	return
}
