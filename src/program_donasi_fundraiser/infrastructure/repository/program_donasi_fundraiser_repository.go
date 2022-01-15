package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/program_donasi/common/constants"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/entity"
	"pemuda-peduli/src/program_donasi_fundraiser/domain/interfaces"
	"strconv"
	"strings"
)

// ProgramDonasiFundraiserRepository
type ProgramDonasiFundraiserRepository struct {
	db *db.ConnectTo
	interfaces.IProgramDonasiFundraiserRepository
}

// NewProgramDonasiFundraiserRepository
func NewProgramDonasiFundraiserRepository(db *db.ConnectTo) ProgramDonasiFundraiserRepository {
	return ProgramDonasiFundraiserRepository{db: db}
}

func (c *ProgramDonasiFundraiserRepository) Insert(ctx context.Context, data *entity.ProgramDonasiFundraiserEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPProgramDonasiFundraiser = utility.GetUUID()

	// Set status published
	data.DonasiType = constants.DonasiTypeOneTIme
	data.Status = constants.StatusPublished

	sql := `INSERT INTO pp_cp_program_donasi_fundraiser `
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
		log.Println("Error insert pp_cp_program_donasi_fundraiser:", err)
		tx.Rollback()
		return
	}
	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

func (c *ProgramDonasiFundraiserRepository) Update(ctx context.Context, data *entity.ProgramDonasiFundraiserEntity) (err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_cp_program_donasi_fundraiser SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(*data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_cp_program_donasi_fundraiser" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_cp_program_donasi_fundraiser = '" + data.IDPPCPProgramDonasiFundraiser + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_cp_program_donasi_fundraiser:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return
}

func (c *ProgramDonasiFundraiserRepository) Find(ctx context.Context, data entity.ProgramDonasiFundraiserQueryEntity) (response []entity.ProgramDonasiFundraiserEntity, count int, err error) {
	sql := `SELECT * FROM pp_cp_program_donasi_fundraiser WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_cp_program_donasi_fundraiser"
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

func (c *ProgramDonasiFundraiserRepository) Get(ctx context.Context, id string) (response entity.ProgramDonasiFundraiserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi_fundraiser WHERE id_pp_cp_program_donasi_fundraiser = $1", id); err != nil {
		return
	}
	return
}

func (c *ProgramDonasiFundraiserRepository) GetSeo(ctx context.Context, seoURL string) (response entity.ProgramDonasiFundraiserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_cp_program_donasi_fundraiser WHERE seo_url = $1", seoURL); err != nil {
		return
	}
	return
}
