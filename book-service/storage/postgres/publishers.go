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

type PublishersRepo struct {
	db *sql.DB
}

func NewPublishersRepo(db *sql.DB) *PublishersRepo {
	return &PublishersRepo{db: db}
}

func (r *PublishersRepo) Create(req *bp.PublishersRes) (*bp.Void, error) {

	req.Id = uuid.New().String()

	query := `
	INSERT INTO publishers (
		id,
		name,
		email,
		phone_number,
		address
	) VALUES($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(query, req.Id, req.Name, req.Email, req.PhoneNumber, req.Address)
	if err != nil {
		log.Println("Error while creating publishers: ", err)
		return nil, err
	}

	log.Println("Successfully created publishers")

	return nil, nil
}
func (r *PublishersRepo) GetById(req *bp.ById) (*bp.PublishersRes, error) {

	var publisher bp.PublishersRes

	query := `
	SELECT
		id,
		name,
		email,
		phone_number,
		address
	FROM 
		publishers
	WHERE 
		deleted_at = 0
	AND 
		id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&publisher.Id,
		&publisher.Name,
		&publisher.Email,
		&publisher.PhoneNumber,
		&publisher.Address,
	)

	if err != nil {
		log.Println("Error while getting publisher by id: ", err)
		return nil, err
	}

	log.Println("Successfully got publisher")

	return &publisher, nil
}
func (r *PublishersRepo) GetAll(req *bp.PublishersGetAllReq) (*bp.PublishersGetAllRes, error) {

	publishers := bp.PublishersGetAllRes{}

	query := `
	SELECT
		id,
		name,
		email,
		phone_number,
		address
	FROM 
		publishers
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		conditions = append(conditions, " LOWER(name) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.Name+"%")
	}
	if req.Address != "" && req.Address != "string" {
		conditions = append(conditions, " LOWER(address) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.Address+"%")
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
		log.Println("publishers not found")
		return nil, errors.New("publishers list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving publishers: ", err)
		return nil, err
	}

	for rows.Next() {
		publisher := bp.PublishersRes{}

		err := rows.Scan(
			&publisher.Id,
			&publisher.Name,
			&publisher.Email,
			&publisher.PhoneNumber,
			&publisher.Address,
		)

		if err != nil {
			log.Println("Error while scanning all publishers: ", err)
			return nil, err
		}

		publishers.Publishers = append(publishers.Publishers, &publisher)
	}

	log.Println("Successfully fetched all publishers")

	return &publishers, nil
}
func (r *PublishersRepo) Update(req *bp.PublishersUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		publishers
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Publisher.Name != "" && req.Publisher.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Publisher.Name)
	}
	if req.Publisher.Address != "" && req.Publisher.Address != "string" {
		conditions = append(conditions, " address = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Publisher.Address)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Publisher.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating publisher: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("publisher with id %s couldn't update", req.Publisher.Id)
	}

	log.Println("Successfully updated publisher")

	return nil, nil
}
func (r *PublishersRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		publishers
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting publisher: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("publisher with id %s not found", req.Id)
	}

	log.Println("Successfully deleted publisher")

	return nil, nil
}
