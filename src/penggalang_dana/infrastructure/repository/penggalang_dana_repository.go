package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/penggalang_dana/domain/entity"
	"pemuda-peduli/src/penggalang_dana/domain/interfaces"
	"strconv"
	"strings"
)

// PenggalangDanaRepository
type PenggalangDanaRepository struct {
	db *db.ConnectTo
	interfaces.IPenggalangDanaRepository
}

// NewPenggalangDanaRepository
func NewPenggalangDanaRepository(db *db.ConnectTo) PenggalangDanaRepository {
	return PenggalangDanaRepository{db: db}
}

// Create
func (c *PenggalangDanaRepository) Insert(ctx context.Context, data *entity.PenggalangDanaEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPPenggalangDana = utility.GetUUID()

	sql := `INSERT INTO pp_cp_penggalang_dana `
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
		log.Println("Error insert pp_cp_penggalang_dana:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *PenggalangDanaRepository) Update(ctx context.Context, data entity.PenggalangDanaEntity) (response entity.PenggalangDanaEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_penggalang_dana SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_penggalang_dana" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_penggalang_dana = '" + data.IDPPCPPenggalangDana + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_penggalang_dana:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// READ
func (c *PenggalangDanaRepository) Find(ctx context.Context, data *entity.PenggalangDanaQueryEntity) (response []entity.PenggalangDanaEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_penggalang_dana WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_cp_penggalang_dana"
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

func (c *PenggalangDanaRepository) Get(ctx context.Context, id string) (response entity.PenggalangDanaEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_penggalang_dana WHERE id_pp_cp_penggalang_dana = $1", id); err != nil {
		return
	}
	return
}
