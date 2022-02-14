package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/berita/common/constants"
	"pemuda-peduli/src/berita/domain/entity"
	"pemuda-peduli/src/berita/domain/interfaces"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"strconv"
	"strings"
)

// BeritaRepository
type BeritaRepository struct {
	db *db.ConnectTo
	interfaces.IBeritaRepository
}

// NewBeritaRepository
func NewBeritaRepository(db *db.ConnectTo) BeritaRepository {
	return BeritaRepository{db: db}
}

// Create data token
func (c *BeritaRepository) Insert(ctx context.Context, data *entity.BeritaEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPBerita = utility.GetUUID()

	// Set status created
	data.Status = constants.StatusCreated

	sql := `INSERT INTO pp_cp_berita `
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
		log.Println("Error insert pp_cp_berita:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

func (c *BeritaRepository) InsertDetail(ctx context.Context, data *entity.BeritaDetailEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPBeritaDetail = utility.GetUUID()

	sql := `INSERT INTO pp_cp_berita_detail `
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
	log.Println("SQL DETAIL (BERITA) : ", sql)
	resp, err := tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_berita_detail:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *BeritaRepository) Update(ctx context.Context, data entity.BeritaEntity, id string) (response entity.BeritaEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_berita SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_berita" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_berita = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_berita:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

func (c *BeritaRepository) UpdateDetail(ctx context.Context, data entity.BeritaDetailEntity, id string) (response entity.BeritaDetailEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_berita_detail SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_berita_detail" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_berita_detail = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_berita_detail:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// READ
func (c *BeritaRepository) Find(ctx context.Context, data *entity.BeritaQueryEntity) (response []entity.BeritaEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_berita WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_cp_berita"
			}
			switch field {
			case "is_deleted":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
			case "is_headline":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
			case "status":
				str.WriteString(field + " = '" + fil.Keyword + "' AND ")
			default:
				str.WriteString(field + " LIKE '%" + fil.Keyword + "%' AND ")
			}

		}
		queryCondition := strings.TrimSuffix(str.String(), "AND ")
		sql += "AND (" + queryCondition + ") "
	}

	// Created at
	if data.CreatedAtFrom != "" {
		if data.CreatedAtTo != "" {
			sql += "AND to_char(created_at, 'YYYY-MM-DD') >= '" + data.CreatedAtFrom + "' AND to_char(created_at, 'YYYY-MM-DD') <= '" + data.CreatedAtTo + "' "
		} else {
			sql += "AND to_char(created_at, 'YYYY-MM-DD') = '" + data.CreatedAtFrom + "' "
		}
	}

	// Created at
	if data.PublishAtFrom != "" {
		if data.PublishAtTo != "" {
			sql += "AND to_char(published_at, 'YYYY-MM-DD') >= '" + data.PublishAtFrom + "' AND to_char(published_at, 'YYYY-MM-DD') <= '" + data.PublishAtTo + "' "
		} else {
			sql += "AND to_char(published_at, 'YYYY-MM-DD') = '" + data.PublishAtFrom + "' "
		}
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

func (c *BeritaRepository) Get(ctx context.Context, id string) (response entity.BeritaEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_berita WHERE id_pp_cp_berita = $1", id); err != nil {
		return
	}
	return
}

func (c *BeritaRepository) GetDetail(ctx context.Context, id string) (response entity.BeritaDetailEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_berita_detail WHERE id_pp_cp_berita = $1", id); err != nil {
		return
	}
	return
}

func (c *BeritaRepository) GetListTag(ctx context.Context) (response []entity.TagEntity, err error) {
	if err = c.db.DBRead.Select(&response, "SELECT tag FROM public.pp_cp_berita GROUP BY tag"); err != nil {
		return
	}
	return
}

func (c *BeritaRepository) GetCountIsHeadline(ctx context.Context) (response int, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT count(*) FROM public.pp_cp_berita WHERE is_headline = true AND status != 'deleted'"); err != nil {
		return
	}
	return
}
