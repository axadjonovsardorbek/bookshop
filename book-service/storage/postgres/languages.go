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

type LanguagesRepo struct {
	db *sql.DB
}

func NewLanguagesRepo(db *sql.DB) *LanguagesRepo {
	return &LanguagesRepo{db: db}
}

func (r *LanguagesRepo) Create(req *bp.LanguagesRes) (*bp.Void, error) {

	req.Id = uuid.New().String()

	query := `
	INSERT INTO languages (
		id,
		name
	) VALUES($1, $2)
	`

	_, err := r.db.Exec(query, req.Id, req.Name)
	if err != nil {
		log.Println("Error while creating language: ", err)
		return nil, err
	}

	log.Println("Successfully created language")

	return nil, nil
}
func (r *LanguagesRepo) GetById(req *bp.ById) (*bp.LanguagesRes, error) {

	var language bp.LanguagesRes

	query := `
	SELECT
		id,
		name
	FROM 
		languages
	WHERE 
		deleted_at = 0
	AND 
		id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&language.Id,
		&language.Name,
	)

	if err != nil {
		log.Println("Error while getting language by id: ", err)
		return nil, err
	}

	log.Println("Successfully got language")

	return &language, nil
}
func (r *LanguagesRepo) GetAll(req *bp.LanguagesGetAllReq) (*bp.LanguagesGetAllRes, error) {

	languages := bp.LanguagesGetAllRes{}

	query := `
	SELECT
		id,
		name
	FROM 
		languages
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
		log.Println("languages not found")
		return nil, errors.New("languages list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving languages: ", err)
		return nil, err
	}

	for rows.Next() {
		language := bp.LanguagesRes{}

		err := rows.Scan(
			&language.Id,
			&language.Name,
		)

		if err != nil {
			log.Println("Error while scanning all languages: ", err)
			return nil, err
		}

		languages.Languages = append(languages.Languages, &language)
	}

	log.Println("Successfully fetched all languages")

	return &languages, nil
}
func (r *LanguagesRepo) Update(req *bp.LanguagesUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		languages
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Language.Name != "" && req.Language.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Language.Name)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Language.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating language: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("language with id %s couldn't update", req.Language.Id)
	}

	log.Println("Successfully updated language")

	return nil, nil
}
func (r *LanguagesRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		languages
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting language: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("language with id %s not found", req.Id)
	}

	log.Println("Successfully deleted language")

	return nil, nil
}
