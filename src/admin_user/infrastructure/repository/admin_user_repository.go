package repository

import (
	"context"
	"errors"
	"log"
	"pemuda-peduli/src/admin_user/domain/entity"
	"pemuda-peduli/src/admin_user/domain/interfaces"
	"pemuda-peduli/src/common/infrastructure/db"
	"pemuda-peduli/src/common/utility"
	"strconv"
	"strings"
)

// AdminUserRepository
type AdminUserRepository struct {
	db *db.ConnectTo
	interfaces.IAdminUserRepository
}

// NewAdminUserRepository
func NewAdminUserRepository(db *db.ConnectTo) AdminUserRepository {
	return AdminUserRepository{db: db}
}

// Create data token
func (c *AdminUserRepository) Insert(ctx context.Context, data *entity.AdminUserEntity) (err error) {

	tx := c.db.DBExec.MustBegin()

	// Generate UUID
	data.IDPPCPAdminUser = utility.GetUUID()

	sql := `INSERT INTO pp_bo_user_admin `
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
		log.Println("Error insert pp_bo_user_admin:", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	data.ID, _ = resp.LastInsertId()
	return
}

// Update
func (c *AdminUserRepository) Update(ctx context.Context, data entity.AdminUserEntity, id string) (response entity.AdminUserEntity, err error) {
	tx := c.db.DBExec.MustBegin()

	// Update Data delivery order
	sql := `Update pp_bo_user_admin SET `
	var str strings.Builder
	fields := utility.GetNamedStruct(data)
	for _, field := range fields {
		if field == "id" || field == "id_user_admin" || field == "created_at" {
			continue
		}
		str.WriteString(field + "=:" + field + ", ")
	}
	queryCondition := strings.TrimSuffix(str.String(), ", ")

	sql += queryCondition + " WHERE id_user_admin = '" + id + "'"
	log.Print("QUERY : ", sql)
	_, err = tx.NamedExec(sql, data)
	if err != nil {
		log.Println("Error insert pp_bo_user_admin:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	response = data

	return
}

// READ
func (c *AdminUserRepository) Find(ctx context.Context, data *entity.AdminUserQueryEntity) (response []entity.AdminUserEntity, count int, err error) {

	sql := `SELECT * FROM pp_bo_user_admin WHERE 1=1 `

	var str strings.Builder
	if len(data.Filter) != 0 {
		for _, fil := range data.Filter {
			field := fil.Field
			if fil.Field == "id" {
				field = "id_user_admin"
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

	// Filter by role level
	sql += "AND role_level > " + strconv.Itoa(ctx.Value("user_role_level").(int)) + " "

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

func (c *AdminUserRepository) Get(ctx context.Context, id string) (response entity.AdminUserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_bo_user_admin WHERE id_user_admin = $1", id); err != nil {
		return
	}
	return
}

func (c *AdminUserRepository) GetByUsername(ctx context.Context, username string) (response entity.AdminUserEntity, err error) {
	if err = c.db.DBRead.Get(&response, "SELECT * FROM pp_bo_user_admin WHERE username = $1", username); err != nil {
		return
	}
	if response.IsDeleted {
		err = errors.New("Failed: user not found")
		return
	}
	return
}
