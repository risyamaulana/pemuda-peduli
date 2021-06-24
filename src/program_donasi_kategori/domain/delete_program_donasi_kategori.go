package domain

import (
	"context"
	"pemuda-peduli/src/program_donasi_kategori/domain/interfaces"
)

func RemoveProgramDonasiKategori(ctx context.Context, repo interfaces.IProgramDonasiKategoriRepository, id string) (err error) {
	err = repo.Delete(ctx, id)
	return
}
