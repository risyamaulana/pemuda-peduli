package repository

import (
	"context"
	"log"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"pemuda-peduli/src/transaction/domain/entity"
	"pemuda-peduli/src/transaction/domain/interfaces"
	"strconv"
	"strings"
)

// TransactionRepository
type TransactionRepository struct {
	db *db.ConnectTo
	interfaces.ITransactionRepository
}

// NewTransactionRepository
func NewTransactionRepository(db *db.ConnectTo) TransactionRepository {
	return TransactionRepository{db: db}
}

// Create data token
func (c *TransactionRepository) Insert(ctx context.Context, data *entity.TransactionEntity) (err error) {

	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPTransaction = utility.GetUUID()

	sql := `INSERT INTO pp_transaction `
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
		log.Println("Error insert pp_transaction:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *TransactionRepository) Update(ctx context.Context, data entity.TransactionEntity, id string) (response entity.TransactionEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_transaction SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_pp_transaction" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_pp_transaction = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error update pp_transaction:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// READ
func (c *TransactionRepository) Find(ctx context.Context, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error) {

	sql := `SELECT * FROM pp_transaction WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_transaction"
			}
			switch field {
			case "status":
				str.WriteString(field + " = '" + fil.Keyword + "' AND ")
			case "is_rutin":
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

	// Created at
	if data.PaidAtFrom != "" {
		if data.PaidAtTo != "" {
			sql += "AND to_char(paid_at, 'YYYY-MM-DD') >= '" + data.PaidAtFrom + "' AND to_char(paid_at, 'YYYY-MM-DD') <= '" + data.PaidAtTo + "' "
		} else {
			sql += "AND to_char(paid_at, 'YYYY-MM-DD') = '" + data.PaidAtFrom + "' "
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

func (c *TransactionRepository) FindMyTransaction(ctx context.Context, data *entity.TransactionQueryEntity) (response []entity.TransactionEntity, count int, err error) {

	sql := `SELECT * FROM pp_transaction WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_pp_transaction"
			}
			switch field {
			case "status":
				str.WriteString(field + " = '" + fil.Keyword + "' AND ")
			case "is_rutin":
				str.WriteString(field + " = " + fil.Keyword + " AND ")
			default:
				str.WriteString(field + " LIKE '%" + fil.Keyword + "%' AND ")
			}

		}
		queryCondition := strings.TrimSuffix(str.String(), "AND ")
		sql += "AND (" + queryCondition + ") "
	}

	sql += "AND user_id = '" + ctx.Value("user_id").(string) + "' "

	// Created at
	if data.CreatedAtFrom != "" {
		if data.CreatedAtTo != "" {
			sql += "AND to_char(created_at, 'YYYY-MM-DD') >= '" + data.CreatedAtFrom + "' AND to_char(created_at, 'YYYY-MM-DD') <= '" + data.CreatedAtTo + "' "
		} else {
			sql += "AND to_char(created_at, 'YYYY-MM-DD') = '" + data.CreatedAtFrom + "' "
		}
	}

	// Created at
	if data.PaidAtFrom != "" {
		if data.PaidAtTo != "" {
			sql += "AND to_char(paid_at, 'YYYY-MM-DD') >= '" + data.PaidAtFrom + "' AND to_char(paid_at, 'YYYY-MM-DD') <= '" + data.PaidAtTo + "' "
		} else {
			sql += "AND to_char(paid_at, 'YYYY-MM-DD') = '" + data.PaidAtFrom + "' "
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

func (c *TransactionRepository) FindRutinTransaction(ctx context.Context) (response []entity.TransactionEntity, err error) {
	if err = c.db.DBRead.Select(&response, "SELECT * FROM pp_transaction where is_rutin = true and status  = 'Paid'"); err != nil {
		return
	}

	return
}

func (c *TransactionRepository) Get(ctx context.Context, id string) (response entity.TransactionEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_transaction WHERE id_pp_transaction = $1", id); err != nil {
		return
	}
	return
}
