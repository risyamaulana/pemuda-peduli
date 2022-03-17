package entity

import "time"

type KategoriProgramDonasiEntity struct {
	ID                          int64      `db:"id"`
	IDPPCPKategoriProgramDonasi string     `db:"id_pp_cp_kategori_program_donasi"`
	Name                        string     `db:"name"`
	CreatedAt                   time.Time  `db:"created_at"`
	CreatedBy                   *string    `db:"created_by"`
	UpdatedAt                   *time.Time `db:"updated_at"`
	UpdatedBy                   *string    `db:"updated_by"`
	IsDeleted                   bool       `db:"is_deleted"`
}

type KategoriProgramDonasiQueryEntity struct {
	Limit  string `db:"limit"`
	Offset string `db:"offset"`
	Filter []KategoriProgramDonasiFilterQueryEntity
	Order  string `db:"order"`
	Sort   string `db:"sort"`
}

type KategoriProgramDonasiFilterQueryEntity struct {
	Field   string
	Keyword string
}
