package postgres

import (
	bp "book/genproto/books"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type CategoriesRepo struct {
	db *sql.DB
}

func NewCategoriesRepo(db *sql.DB) *CategoriesRepo {
	return &CategoriesRepo{db: db}
}

func (r *CategoriesRepo) Create(req *bp.CategoriesRes) (*bp.Void, error) {

	req.Id = uuid.New().String()

	query := `
	INSERT INTO categories (
		id,
		name,
		description
	) VALUES($1, $2, $3)
	`

	_, err := r.db.Exec(query, req.Id, req.Name, req.Description)
	if err != nil {
		log.Println("Error while creating category: ", err)
		return nil, err
	}

	log.Println("Successfully created category")

	return nil, nil
}
func (r *CategoriesRepo) GetById(req *bp.ById) (*bp.CategoriesRes, error) {

	var category bp.CategoriesRes

	query := `
	SELECT
		id,
		name,
		description
	FROM 
		categories
	WHERE 
		deleted_at = 0
	AND 
		id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.Description,
	)

	if err != nil {
		log.Println("Error while getting category by id: ", err)
		return nil, err
	}

	log.Println("Successfully got category")

	return &category, nil
}
func (r *CategoriesRepo) GetAll(req *bp.CategoriesGetAllReq) (*bp.CategoriesGetAllRes, error) {

	categories := bp.CategoriesGetAllRes{}

	query := `
	SELECT
		id,
		name,
		description
	FROM 
		categories
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		conditions = append(conditions, " LOWER(name) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.Name+"%")
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var limit int32 = 10
	var offset int32 = (req.Page.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	if err == sql.ErrNoRows {
		log.Println("categories not found")
		return nil, errors.New("categories list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving categories: ", err)
		return nil, err
	}

	for rows.Next() {
		category := bp.CategoriesRes{}

		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Description,
		)

		if err != nil {
			log.Println("Error while scanning all categories: ", err)
			return nil, err
		}

		categories.Categories = append(categories.Categories, &category)
	}

	log.Println("Successfully fetched all categories")

	return &categories, nil
}
func (r *CategoriesRepo) Update(req *bp.CategoriesUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		categories
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Category.Name != "" && req.Category.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Category.Name)
	}
	if req.Category.Description != "" && req.Category.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Category.Description)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Category.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating category: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("category with id %s couldn't update", req.Category.Id)
	}

	log.Println("Successfully updated category")

	return nil, nil
}
func (r *CategoriesRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		categories
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting category: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("category with id %s not found", req.Id)
	}

	log.Println("Successfully deleted category")

	return nil, nil
}
