package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/transaction/common/constants"
	"pemuda-peduli/src/transaction/domain/entity"
	"pemuda-peduli/src/transaction/domain/interfaces"
	"pemuda-peduli/src/transaction/infrastructure/repository"

	donasiRutinDom "pemuda-peduli/src/program_donasi_rutin/domain"
	donasiRutinRep "pemuda-peduli/src/program_donasi_rutin/infrastructure/repository"

	donasiDom "pemuda-peduli/src/program_donasi/domain"

	userDom "pemuda-peduli/src/user/domain"
	userRep "pemuda-peduli/src/user/infrastructure/repository"
)

func CreateTransaction(ctx context.Context, db *db.ConnectTo, data *entity.TransactionEntity) (err error) {
	// Repo
	repo := repository.NewTransactionRepository(db)
	userRepo := userRep.NewUserRepository(db)
	donasiRutinRepo := donasiRutinRep.NewProgramDonasiRutinRepository(db)

	// Check user
	userData, err := userDom.ReadUser(ctx, &userRepo, ctx.Value("user_id").(string))
	if err != nil {
		err = errors.New("Failed user unauthorized for this transaction")
		return
	}
	// Applied User Data
	data.UserID = userData.IDUser
	data.Username = userData.Username
	data.Email = userData.Email
	data.NamaLengkap = userData.NamaLengkap
	data.NamaPanggilan = userData.NamaPanggilan

	// Check Data Donasi
	if data.IsRutin {
		// Check data donasi rutin
		donasiRutinData, errDonasiRutin := donasiRutinDom.GetProgramDonasiRutin(ctx, &donasiRutinRepo, data.IDPPCPProgramDonasiRutin)
		if errDonasiRutin != nil {
			err = errors.New("Failed, donasi not found")
			return
		}
		data.DonasiTitle = donasiRutinData.Title

		// Check QR data
		if data.PaymentMethod == constants.PaymentQris {
			if donasiRutinData.QrisImageURL == nil || *donasiRutinData.QrisImageURL == "" {
				err = errors.New("Failed, transaction can't use QRIS")
				return
			}
			data.QRPaymentURL = donasiRutinData.QrisImageURL
		}
	} else {
		// Check data donasi one time
		log.Println("ID : ", data.IDPPCPProgramDonasi)
		donasiData, errDonasi := donasiDom.GetProgramDonasi(ctx, db, data.IDPPCPProgramDonasi)
		if errDonasi != nil {
			err = errors.New("Failed, donasi not found")
			return
		}
		data.DonasiTitle = donasiData.Title
		// Check QR data
		if data.PaymentMethod == constants.PaymentQris {
			if donasiData.QrisImageURL == nil || *donasiData.QrisImageURL == "" {
				err = errors.New("Failed, transaction can't use QRIS")
				return
			}
			data.QRPaymentURL = donasiData.QrisImageURL
		}
	}

	err = insertTransaction(ctx, &repo, data)
	return
}

func insertTransaction(ctx context.Context, repo interfaces.ITransactionRepository, data *entity.TransactionEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
