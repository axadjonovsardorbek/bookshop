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

type AuthorsRepo struct {
	db *sql.DB
}

func NewAuthorsRepo(db *sql.DB) *AuthorsRepo {
	return &AuthorsRepo{db: db}
}

func (r *AuthorsRepo) Create(req *bp.AuthorsRes) (*bp.Void, error) {

	req.Id = uuid.New().String()

	query := `
	INSERT INTO authors (
		id,
		name,
		biography
	) VALUES($1, $2, $3)
	`

	_, err := r.db.Exec(query, req.Id, req.Name, req.Biography)
	if err != nil {
		log.Println("Error while creating author: ", err)
		return nil, err
	}

	log.Println("Successfully created author")

	return nil, nil
}
func (r *AuthorsRepo) GetById(req *bp.ById) (*bp.AuthorsRes, error) {

	var author bp.AuthorsRes

	query := `
	SELECT
		id,
		name,
		biography
	FROM 
		authors
	WHERE 
		deleted_at = 0
	AND 
		id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&author.Id,
		&author.Name,
		&author.Biography,
	)

	if err != nil {
		log.Println("Error while getting author by id: ", err)
		return nil, err
	}

	log.Println("Successfully got author")

	return &author, nil
}
func (r *AuthorsRepo) GetAll(req *bp.AuthorsGetAllReq) (*bp.AuthorsGetAllRes, error) {

	authors := bp.AuthorsGetAllRes{}

	query := `
	SELECT
		id,
		name,
		biography
	FROM 
		authors
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
		log.Println("authors not found")
		return nil, errors.New("authors list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving authors: ", err)
		return nil, err
	}

	for rows.Next() {
		author := bp.AuthorsRes{}

		err := rows.Scan(
			&author.Id,
			&author.Name,
			&author.Biography,
		)

		if err != nil {
			log.Println("Error while scanning all authors: ", err)
			return nil, err
		}

		authors.Authors = append(authors.Authors, &author)
	}

	log.Println("Successfully fetched all authors")

	return &authors, nil
}
func (r *AuthorsRepo) Update(req *bp.AuthorsUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		authors
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Author.Name != "" && req.Author.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Author.Name)
	}
	if req.Author.Biography != "" && req.Author.Biography != "string" {
		conditions = append(conditions, " biography = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Author.Biography)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Author.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating author: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("author with id %s couldn't update", req.Author.Id)
	}

	log.Println("Successfully updated author")

	return nil, nil
}
func (r *AuthorsRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		authors
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting author: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("author with id %s not found", req.Id)
	}

	log.Println("Successfully deleted author")

	return nil, nil
}
