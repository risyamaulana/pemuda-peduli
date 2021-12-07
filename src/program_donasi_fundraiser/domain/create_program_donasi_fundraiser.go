package domain

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/entity"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/interfaces"
	"pemuda-peduli/src/program_donasi_fundraiser/infrastructure/repository"

	donasiDom "pemuda-peduli/src/program_donasi/domain"
	userDom "pemuda-peduli/src/user/domain"
	userRep "pemuda-peduli/src/user/infrastructure/repository"
)

func CreateProgramDonasiFundraiser(ctx context.Context, db *db.ConnectTo, data *entity.ProgramDonasiFundraiserEntity) (err error) {
	// Repo
	repo := repository.NewProgramDonasiFundraiserRepository(db)
	userRepo := userRep.NewUserRepository(db)

	// Get Donasi Data
	donasiData, err := donasiDom.GetProgramDonasi(ctx, db, data.IDPPCPProgramDonasi)
	if err != nil {
		log.Println("Failed get donasi data: ", err)
		err = errors.New("Failed get donasi data")
		return
	}

	// Get user data
	userData, err := userDom.ReadUser(ctx, &userRepo, ctx.Value("user_id").(string))
	if err != nil {
		log.Println("Failed get user data: ", err)
		err = errors.New("Failed get user data (unauthorization)")
		return
	}

	// Applied user data
	data.IDUser = userData.IDUser
	data.Username = userData.Username
	data.Email = userData.Email
	data.PhoneNumber = userData.PhoneNumber
	data.NamaLengkap = userData.NamaLengkap
	data.NamaPanggilan = userData.NamaPanggilan
	data.Alamat = userData.Alamat

	// Applied donasi data
	data.IDPPCPPenggalangDana = donasiData.PenggalangDana.IDPPCPPenggalangDana
	data.IDPPCPProgramDonasi = donasiData.IDPPCPProgramDonasi
	data.DonasiType = donasiData.DonasiType
	data.Tag = donasiData.Tag
	data.Content = donasiData.Detail.Content
	data.ThumbnailImageURL = donasiData.ThumbnailImageURL
	data.Status = donasiData.Status
	data.ValidFrom = donasiData.ValidFrom
	data.ValidTo = donasiData.ValidTo

	data.CreatedBy = &userData.Username

	data.KitaBisaLink = donasiData.KitaBisaLink
	data.AyoBantuLink = donasiData.AyoBantuLink
	data.IDPPCPMasterQris = donasiData.IDPPCPMasterQris
	data.QrisImageURL = donasiData.QrisImageURL

	err = insertProgramDonasiFundraiser(ctx, &repo, data)
	if err != nil {
		log.Println("Failed insert program donasi fundraiser: ", err)
		return
	}
	return
}

func insertProgramDonasiFundraiser(ctx context.Context, repo interfaces.IProgramDonasiFundraiserRepository, data *entity.ProgramDonasiFundraiserEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
