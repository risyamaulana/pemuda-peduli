package domain

import (
	"context"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/entity"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/interfaces"
	"pemuda-peduli/src/program_donasi_fundraiser/infrastructure/repository"
)

func FindProgramDonasiFundraiser(ctx context.Context, db *db.ConnectTo, data entity.ProgramDonasiFundraiserQueryEntity) (responses []entity.ProgramDonasiFundraiserEntity, count int, err error) {
	// Repo
	repo := repository.NewProgramDonasiFundraiserRepository(db)

	responses, count, err = findProgramDonasiFundraiser(ctx, &repo, data)
	return
}

func GetProgramDonasiFundraiser(ctx context.Context, db *db.ConnectTo, id string) (response entity.ProgramDonasiFundraiserEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiFundraiserRepository(db)

	response, err = getProgramDonasiFundraiser(ctx, &repo, id)
	return
}

func GetProgramDonasiFundraiserSeo(ctx context.Context, db *db.ConnectTo, seoURL string) (response entity.ProgramDonasiFundraiserEntity, err error) {
	// Repo
	repo := repository.NewProgramDonasiFundraiserRepository(db)

	response, err = getProgramDonasiFundraiserSeo(ctx, &repo, seoURL)
	return
}

func findProgramDonasiFundraiser(ctx context.Context, repo interfaces.IProgramDonasiFundraiserRepository, data entity.ProgramDonasiFundraiserQueryEntity) (responses []entity.ProgramDonasiFundraiserEntity, count int, err error) {
	responses, count, err = repo.Find(ctx, data)
	return
}

func getProgramDonasiFundraiser(ctx context.Context, repo interfaces.IProgramDonasiFundraiserRepository, id string) (response entity.ProgramDonasiFundraiserEntity, err error) {
	response, err = repo.Get(ctx, id)
	return
}

func getProgramDonasiFundraiserSeo(ctx context.Context, repo interfaces.IProgramDonasiFundraiserRepository, seoURL string) (response entity.ProgramDonasiFundraiserEntity, err error) {
	response, err = repo.GetSeo(ctx, seoURL)
	return
}
