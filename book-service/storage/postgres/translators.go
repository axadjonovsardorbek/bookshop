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

type TranslatorsRepo struct {
	db *sql.DB
}

func NewTranslatorsRepo(db *sql.DB) *TranslatorsRepo {
	return &TranslatorsRepo{db: db}
}

func (r *TranslatorsRepo) Create(req *bp.TranslatorsRes) (*bp.Void, error) {

	req.Id = uuid.New().String()

	query := `
	INSERT INTO translators (
		id,
		name,
		surname
	) VALUES($1, $2, $3)
	`

	_, err := r.db.Exec(query, req.Id, req.Name, req.Surname)
	if err != nil {
		log.Println("Error while creating translators: ", err)
		return nil, err
	}

	log.Println("Successfully created translators")

	return nil, nil
}
func (r *TranslatorsRepo) GetById(req *bp.ById) (*bp.TranslatorsRes, error) {

	var translator bp.TranslatorsRes

	query := `
	SELECT
		id,
		name,
		surname
	FROM 
		translators
	WHERE 
		deleted_at = 0
	AND 
		id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&translator.Id,
		&translator.Name,
		&translator.Surname,
	)

	if err != nil {
		log.Println("Error while getting translator by id: ", err)
		return nil, err
	}

	log.Println("Successfully got translator")

	return &translator, nil
}
func (r *TranslatorsRepo) GetAll(req *bp.TranslatorsGetAllReq) (*bp.TranslatorsGetAllRes, error) {

	translators := bp.TranslatorsGetAllRes{}

	query := `
	SELECT
		id,
		name,
		surname
	FROM 
		translators
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		conditions = append(conditions, " LOWER(name) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.Name+"%")
	}
	if req.Surname != "" && req.Surname != "string" {
		conditions = append(conditions, " LOWER(surname) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.Surname+"%")
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
		log.Println("translators not found")
		return nil, errors.New("translators list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving translators: ", err)
		return nil, err
	}

	for rows.Next() {
		translator := bp.TranslatorsRes{}

		err := rows.Scan(
			&translator.Id,
			&translator.Name,
			&translator.Surname,
		)

		if err != nil {
			log.Println("Error while scanning all translators: ", err)
			return nil, err
		}

		translators.Translators = append(translators.Translators, &translator)
	}

	log.Println("Successfully fetched all translators")

	return &translators, nil

}
func (r *TranslatorsRepo) Update(req *bp.TranslatorsUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		translators
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Translator.Name != "" && req.Translator.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Translator.Name)
	}
	if req.Translator.Surname != "" && req.Translator.Surname != "string" {
		conditions = append(conditions, " surname = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Translator.Surname)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Translator.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating translator: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("translator with id %s couldn't update", req.Translator.Id)
	}

	log.Println("Successfully updated translator")

	return nil, nil
}
func (r *TranslatorsRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		translators
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting translator: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("translator with id %s not found", req.Id)
	}

	log.Println("Successfully deleted translator")

	return nil, nil
}
