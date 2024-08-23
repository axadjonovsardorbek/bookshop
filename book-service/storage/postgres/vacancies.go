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

type VacanciesRepo struct {
	db *sql.DB
}

func NewVacanciesRepo(db *sql.DB) *VacanciesRepo {
	return &VacanciesRepo{db: db}
}

func (r *VacanciesRepo) Create(req *bp.VacanciesCreateReq) (*bp.Void, error) {

	id := uuid.New().String()

	query := `
	INSERT INTO vacancies (
		id,
		publisher_id,
		title,
		description,
		salary_from,
		salary_to,
		working_styles,
		working_types,
		phone_number,
		location
	) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.Exec(
		query,
		id,
		req.PublisherId,
		req.Title,
		req.Description,
		req.SalaryFrom,
		req.SalaryTo,
		req.WorkingStyles,
		req.WorkingTypes,
		req.PhoneNumber,
		req.Location,
	)
	if err != nil {
		log.Println("Error while creating vacancies: ", err)
		return nil, err
	}

	log.Println("Successfully created vacancies")

	return nil, nil
}
func (r *VacanciesRepo) GetById(req *bp.ById) (*bp.VacanciesRes, error) {
	vacancy := bp.VacanciesRes{
		Publisher: &bp.PublishersRes{},
	}  

	query := `
	SELECT 
		v.id,
		v.title,
		v.description,
		v.status,
		v.salary_from,
		v.salary_to,
		v.working_styles,
		v.working_types,
		v.phone_number,
		v.location,
		v.view_count,
		p.id,
		p.name,
		p.email,
		p.phone_number,
		p.address
	FROM 
		vacancy v
	JOIN 
		publishers p
	ON
		v.publisher_id = p.id AND p.deleted_at = 0
	WHERE 
		v.id = $1 AND v.deleted_at = 0
	` 

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&vacancy.Id,
		&vacancy.Title,
		&vacancy.Description,
		&vacancy.Status,
		&vacancy.SalaryFrom,
		&vacancy.SalaryTo,
		&vacancy.WorkingStyles,
		&vacancy.WorkingTypes,
		&vacancy.PhoneNumber,
		&vacancy.Location,
		&vacancy.ViewCount,
		&vacancy.Publisher.Id,
		&vacancy.Publisher.Name,
		&vacancy.Publisher.Email,
		&vacancy.Publisher.PhoneNumber,
		&vacancy.Publisher.Address,
	)

	if err != nil {
		log.Println("Error while getting vacancy by id: ", err)
		return nil, err
	}

	log.Println("Successfully got vacancy")

	return &vacancy, nil
}
func (r *VacanciesRepo) GetAll(req *bp.VacanciesGetAllReq) (*bp.VacanciesGetAllRes, error) {
	vacancies := bp.VacanciesGetAllRes{}

	query := `
	SELECT 
		v.id,
		v.title,
		v.description,
		v.status,
		v.salary_from,
		v.salary_to,
		v.working_styles,
		v.working_types,
		v.phone_number,
		v.location,
		v.view_count,
		p.id,
		p.name,
		p.email,
		p.phone_number,
		p.address
	FROM 
		vacancy v
	JOIN 
		publishers p
	ON
		v.publisher_id = p.id AND p.deleted_at = 0
	WHERE 
		v.deleted_at = 0
	` 

	var args []interface{}
	var conditions []string

	if req.SalaryFrom >= 0{
		conditions = append(conditions, " v.salary_from >= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.SalaryFrom)
	}
	if req.SalaryTo > 0{
		conditions = append(conditions, " v.salary_to <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.SalaryTo)
	}
	if req.Status != "" && req.Status != "string" {
		conditions = append(conditions, " v.status = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Status)
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
		log.Println("vacancies not found")
		return nil, errors.New("vacancies list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving vacancies: ", err)
		return nil, err
	}

	for rows.Next() {
		vacancy := bp.VacanciesRes{
			Publisher: &bp.PublishersRes{},
		}

		err := rows.Scan(
			&vacancy.Id,
			&vacancy.Title,
			&vacancy.Description,
			&vacancy.Status,
			&vacancy.SalaryFrom,
			&vacancy.SalaryTo,
			&vacancy.WorkingStyles,
			&vacancy.WorkingTypes,
			&vacancy.PhoneNumber,
			&vacancy.Location,
			&vacancy.ViewCount,
			&vacancy.Publisher.Id,
			&vacancy.Publisher.Name,
			&vacancy.Publisher.Email,
			&vacancy.Publisher.PhoneNumber,
			&vacancy.Publisher.Address,
		)

		if err != nil {
			log.Println("Error while scanning all vacancies: ", err)
			return nil, err
		}

		vacancies.Vacancies = append(vacancies.Vacancies, &vacancy)
	}

	log.Println("Successfully fetched all vacancies")
	return &vacancies, nil
}
func (r *VacanciesRepo) Update(req *bp.VacanciesUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		vacancies
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Vacancy.Title != "" && req.Vacancy.Title != "string" {
		conditions = append(conditions, " title = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.Title)
	}
	if req.Vacancy.Description != "" && req.Vacancy.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.Description)
	}
	if req.Vacancy.Location != "" && req.Vacancy.Location != "string" {
		conditions = append(conditions, " location = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.Location)
	}
	if req.Vacancy.PhoneNumber != "" && req.Vacancy.PhoneNumber != "string" {
		conditions = append(conditions, " phone_number = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.PhoneNumber)
	}
	if req.Vacancy.WorkingStyles != "" && req.Vacancy.WorkingStyles != "string" {
		conditions = append(conditions, " working_styles = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.WorkingStyles)
	}
	if req.Vacancy.WorkingTypes != "" && req.Vacancy.WorkingTypes != "string" {
		conditions = append(conditions, " working_types = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.WorkingTypes)
	}
	if req.Vacancy.SalaryFrom > 0 {
		conditions = append(conditions, " salary_from = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.SalaryFrom)
	}
	if req.Vacancy.SalaryTo > 0 {
		conditions = append(conditions, " salary_to = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.SalaryTo)
	}
	if req.Vacancy.ViewCount == 1 {
		conditions = append(conditions, " view_count = view_count + $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Vacancy.ViewCount)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating vacancy: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("vacancy with id %s couldn't update", req.Id)
	}

	log.Println("Successfully updated vacancy")

	return nil, nil
}
func (r *VacanciesRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		vacancies
	SET 
		status = 'non-active'
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	AND 
		status <> 'active'
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting vacancies: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("vacancy with id %s not found", req.Id)
	}

	log.Println("Successfully deleted vacansy")

	return nil, nil
}
