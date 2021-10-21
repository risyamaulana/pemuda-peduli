package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/program_donasi/common/constants"
	"pemuda-peduli/src/program_donasi_rutin/domain/entity"
	"pemuda-peduli/src/program_donasi_rutin/domain/interfaces"
	"strconv"
	"strings"
)

// ProgramDonasiRutinRepository
type ProgramDonasiRutinRepository struct {
	db *db.ConnectTo
	interfaces.IProgramDonasiRutinRepository
}

// NewProgramDonasiRutinRepository
func NewProgramDonasiRutinRepository(db *db.ConnectTo) ProgramDonasiRutinRepository {
	return ProgramDonasiRutinRepository{db: db}
}

// Create
func (c *ProgramDonasiRutinRepository) Insert(ctx context.Context, data *entity.ProgramDonasiRutinEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPProgramDonasiRutin = utility.GetUUID()

	// Set status created
	data.Status = constants.StatusCreated

	sql := `INSERT INTO pp_cp_program_donasi_rutin `
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
		log.Println("Error insert pp_cp_program_donasi_rutin:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

func (c *ProgramDonasiRutinRepository) InsertPaket(ctx context.Context, data *entity.ProgramDonasiRutinPaketEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPProgramDonasiRutinPaket = utility.GetUUID()

	sql := `INSERT INTO pp_cp_program_donasi_rutin_paket `
	var strField strings.Builder
	var strValue strings.Builder

	fields := utility.GetNamedStruct(*data)
	for _, field := range fields {
		if field != "id" {
			strField.WriteString(field + ",")
			strValue.WriteString(":" + field + ",")
		}
	}

	sql += "(" + strings.TrimSuffix(strField.String(), ",") + ")" + " VALUES(" + strings.TrimSuffix(strValue.String(), ",") + ")"
	resp, err := tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_program_donasi_rutin_paket", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

func (c *ProgramDonasiRutinRepository) InsertNews(ctx context.Context, data *entity.ProgramDonasiRutinNewsEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	sql := `INSERT INTO pp_cp_program_donasi_rutin_news `
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
		log.Println("Error insert pp_cp_program_donasi_rutin_news:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *ProgramDonasiRutinRepository) Update(ctx context.Context, data entity.ProgramDonasiRutinEntity, id string) (response entity.ProgramDonasiRutinEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_program_donasi_rutin SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_program_donasi_rutin" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_program_donasi_rutin = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_program_donasi_rutin:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

func (c *ProgramDonasiRutinRepository) UpdatePaket(ctx context.Context, data entity.ProgramDonasiRutinPaketEntity, id string) (response entity.ProgramDonasiRutinPaketEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data
	sql := `Update pp_cp_program_donasi_rutin_paket SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_program_donasi_rutin_paket " || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")
	sql += queryCondition + " WHERE id_pp_cp_program_donasi_rutin_paket = '" + id + "' "
	log.Println("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_program_donasi_rutin_paket:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

func (c *ProgramDonasiRutinRepository) UpdateNews(ctx context.Context, data entity.ProgramDonasiRutinNewsEntity, id int64) (response entity.ProgramDonasiRutinNewsEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_program_donasi_rutin_news SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id = '" + strconv.FormatInt(id, 10) + "'"
	log.Print("QUERY pp_cp_program_donasi_rutin_news: ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error update pp_cp_program_donasi_rutin_news:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// READ
func (c *ProgramDonasiRutinRepository) Find(ctx context.Context, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_program_donasi_rutin WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_cp_program_donasi_rutin"
			}
			switch field {
			case "is_deleted":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
			case "is_show":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
			case "status":
				str.WriteString(field + " = '" + fil.Keyword + "' AND ")
			case "donasi_type":
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

func (c *ProgramDonasiRutinRepository) FindPaket(ctx context.Context, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinPaketEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_program_donasi_rutin_paket WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_cp_program_donasi_rutin_paket"
			}
			switch field {
			case "is_deleted":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
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

func (c *ProgramDonasiRutinRepository) FindNews(ctx context.Context, data *entity.ProgramDonasiRutinQueryEntity) (response []entity.ProgramDonasiRutinNewsEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_program_donasi_rutin_news WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field

			switch field {
			case "id":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
			case "is_deleted":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
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

func (c *ProgramDonasiRutinRepository) Get(ctx context.Context, id string) (response entity.ProgramDonasiRutinEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi_rutin WHERE id_pp_cp_program_donasi_rutin = $1", id); err != nil {
		return
	}
	return
}

func (c *ProgramDonasiRutinRepository) GetPaket(ctx context.Context, id string) (response entity.ProgramDonasiRutinPaketEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi_rutin_paket WHERE id_pp_cp_program_donasi_rutin_paket = $1", id); err != nil {
		return
	}
	return
}

func (c *ProgramDonasiRutinRepository) GetNews(ctx context.Context, id int64) (response entity.ProgramDonasiRutinNewsEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi_rutin_news WHERE id = $1", id); err != nil {
		return
	}
	return
}
