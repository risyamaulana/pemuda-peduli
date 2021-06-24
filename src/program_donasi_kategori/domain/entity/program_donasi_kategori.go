package entity

type ProgramDonasiKategoriEntity struct {
	ID                          int64  `db:"id"`
	IDPPCPProgramDonasiKategori string `db:"id_pp_cp_master_kategori_donasi"`
	KategoriName                string `db:"kategori_name"`
}

type ProgramDonasiKategoriQueryEntity struct {
	Limit         string `db:"limit"`
	Offset        string `db:"offset"`
	Filter        []ProgramDonasiKategoriFilterQueryEntity
	Order         string `db:"order"`
	Sort          string `db:"sort"`
	CreatedAtFrom string `db:"created_at_from"`
	CreatedAtTo   string `db:"created_at_to"`
	PublishAtFrom string `db:"publish_at_from"`
	PublishAtTo   string `db:"publish_at_to"`
}

type ProgramDonasiKategoriFilterQueryEntity struct {
	Field   string
	Keyword string
}
