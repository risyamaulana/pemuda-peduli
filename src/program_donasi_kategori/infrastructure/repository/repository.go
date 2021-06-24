package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"

	"pemuda-peduli/src/program_donasi_kategori/domain/entity"
	"pemuda-peduli/src/program_donasi_kategori/domain/interfaces"
	"strconv"
	"strings"
)

// ProgramDonasiKategoriRepository
type ProgramDonasiKategoriRepository struct {
	db *db.ConnectTo
	interfaces.IProgramDonasiKategoriRepository
}

// NewProgramDonasiKategoriRepository
func NewProgramDonasiKategoriRepository(db *db.ConnectTo) ProgramDonasiKategoriRepository {
	return ProgramDonasiKategoriRepository{db: db}
}

// Create data token
func (c *ProgramDonasiKategoriRepository) Insert(ctx context.Context, data *entity.ProgramDonasiKategoriEntity) (err error) {

	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPProgramDonasiKategori = utility.GetUUID()

	sql := `INSERT INTO pp_cp_master_kategori_donasi `
	var strField strings.Builder
	var strValue strings.Builder
	filedItem := utility.GetNamedStruct(*data)
	for _, field := range filedItem {
		if field != "id" {
			strField.WriteString(field + ",")
			strValue.WriteString(":" + field + ",")
		}
	}

	sql += "(" + strings.TrimSuffix(strField.String(), ",") + ")" + " VALUES(" + strings.TrimSuffix(strValue.String(), ",") + ")"
	resp, err := tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_master_kategori_donasi:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *ProgramDonasiKategoriRepository) Update(ctx context.Context, data entity.ProgramDonasiKategoriEntity, id string) (response entity.ProgramDonasiKategoriEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_master_kategori_donasi SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_master_kategori_donasi" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_master_kategori_donasi = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_master_kategori_donasi:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// Delete
func (c *ProgramDonasiKategoriRepository) Delete(ctx context.Context, id string) (err error) {
	if _, err = c.db.DBExec.Exec("DELETE FROM public.pp_cp_master_kategori_donasi WHERE id_pp_cp_master_kategori_donasi = $1", id); err != nil {
		return
	}

	return
}

// READ
func (c *ProgramDonasiKategoriRepository) Find(ctx context.Context, data *entity.ProgramDonasiKategoriQueryEntity) (response []entity.ProgramDonasiKategoriEntity, count int, err error) {

	sql := `SELECT * FROM pp_cp_master_kategori_donasi WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_cp_master_kategori_donasi"
			}

			str.WriteString(field + " LIKE '%" + fil.Keyword + "%' AND ")

		}
		queryCondition := strings.TrimSuffix(str.String(), "AND ")
		sql += "AND (" + queryCondition + ") "
	}

	log.Println(sql)
	// Get count Total data
	sqlCount := strings.ReplaceAll(sql, "*", "count(*)")
	if err = c.db.DBRead.Get(&count, sqlCount); err != nil {
		return
	}

	// Order param
	if data.Order != "" {
		sql += "ORDER BY " + data.Order + " " + data.Sort + " "
	} else {
		sql += "ORDER BY created_at DESC "
	}

	// Limit Offset
	limit, _ := strconv.Atoi(data.Limit)
	offset, _ := strconv.Atoi(data.Offset)
	if limit == 0 {
		limit = 5
	}

	if offset == 0 {
		offset = (1 * limit) - limit
	} else {
		offset = (offset * limit) - limit
	}

	sql += "LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa(offset)

	// Result query
	if err = c.db.DBRead.Select(&response, sql); err != nil {
		return
	}
	return
}

func (c *ProgramDonasiKategoriRepository) Get(ctx context.Context, id string) (response entity.ProgramDonasiKategoriEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_master_kategori_donasi WHERE id_pp_cp_master_kategori_donasi = $1", id); err != nil {
		return
	}
	return
}
