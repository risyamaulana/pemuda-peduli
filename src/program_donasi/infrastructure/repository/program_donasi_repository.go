package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/program_donasi/common/constants"
	"pemuda-peduli/src/program_donasi/domain/entity"
	"pemuda-peduli/src/program_donasi/domain/interfaces"
	"strconv"
	"strings"
)

// ProgramDonasiRepository
type ProgramDonasiRepository struct {
	db *db.ConnectTo
	interfaces.IProgramDonasiRepository
}

// NewProgramDonasiRepository
func NewProgramDonasiRepository(db *db.ConnectTo) ProgramDonasiRepository {
	return ProgramDonasiRepository{db: db}
}

// Create data token
func (c *ProgramDonasiRepository) Insert(ctx context.Context, data *entity.ProgramDonasiEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPProgramDonasi = utility.GetUUID()

	// Set status created
	data.DonasiType = constants.DonasiTypeOneTIme
	data.Status = constants.StatusCreated

	sql := `INSERT INTO pp_cp_program_donasi `
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
		log.Println("Error insert pp_cp_program_donasi:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

func (c *ProgramDonasiRepository) InsertDetail(ctx context.Context, data *entity.ProgramDonasiDetailEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPProgramDonasiDetail = utility.GetUUID()

	sql := `INSERT INTO pp_cp_program_donasi_detail `
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
		log.Println("Error insert pp_cp_program_donasi_detail:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

func (c *ProgramDonasiRepository) InsertNews(ctx context.Context, data *entity.ProgramDonasiNewsEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	sql := `INSERT INTO pp_cp_program_donasi_news `
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
		log.Println("Error insert pp_cp_program_donasi_news:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *ProgramDonasiRepository) Update(ctx context.Context, data entity.ProgramDonasiEntity, id string) (response entity.ProgramDonasiEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_program_donasi SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_program_donasi" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_program_donasi = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_program_donasi:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

func (c *ProgramDonasiRepository) UpdateDetail(ctx context.Context, data entity.ProgramDonasiDetailEntity, id string) (response entity.ProgramDonasiDetailEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_program_donasi_detail SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_program_donasi_detail" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_program_donasi_detail = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error update pp_cp_program_donasi_detail:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

func (c *ProgramDonasiRepository) UpdateNews(ctx context.Context, data entity.ProgramDonasiNewsEntity, id int64) (response entity.ProgramDonasiNewsEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_program_donasi_news SET `
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
	log.Print("QUERY pp_cp_program_donasi_news: ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error update pp_cp_program_donasi_news:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// READ
func (c *ProgramDonasiRepository) Find(ctx context.Context, data *entity.ProgramDonasiQueryEntity) (response []entity.ProgramDonasiEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_program_donasi WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_cp_program_donasi"
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
				str.WriteString("LOWER(" + field + ") LIKE LOWER('%" + fil.Keyword + "%') AND ")
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

func (c *ProgramDonasiRepository) FindNews(ctx context.Context, data *entity.ProgramDonasiQueryEntity) (response []entity.ProgramDonasiNewsEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_program_donasi_news WHERE 1=1 `

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
				str.WriteString("LOWER(" + field + ") LIKE LOWER('%" + fil.Keyword + "%') AND ")
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

func (c *ProgramDonasiRepository) Get(ctx context.Context, id string) (response entity.ProgramDonasiEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi WHERE id_pp_cp_program_donasi = $1", id); err != nil {
		return
	}
	return
}

func (c *ProgramDonasiRepository) GetBySeo(ctx context.Context, seo string) (response entity.ProgramDonasiEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi WHERE seo_url = $1", seo); err != nil {
		return
	}
	return
}

func (c *ProgramDonasiRepository) GetDetail(ctx context.Context, id string) (response entity.ProgramDonasiDetailEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi_detail WHERE id_pp_cp_program_donasi = $1", id); err != nil {
		return
	}
	return
}

func (c *ProgramDonasiRepository) GetNews(ctx context.Context, id int64) (response entity.ProgramDonasiNewsEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi_news WHERE id = $1", id); err != nil {
		return
	}
	return
}
