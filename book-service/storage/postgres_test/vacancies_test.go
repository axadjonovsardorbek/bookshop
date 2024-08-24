package postgres_test

import (
	bp "book/genproto/books"
	"book/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateVacancy(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewVacanciesRepo(db)

	req := &bp.VacanciesCreateReq{
		PublisherId:  uuid.New().String(),
		Title:        "Software Engineer",
		Description:  "Develop and maintain software applications.",
		SalaryFrom:   60000,
		SalaryTo:     80000,
		WorkingStyles: "Remote",
		WorkingTypes: "Full-time",
		PhoneNumber:  "1234567890",
		Location:     "New York",
	}

	mock.ExpectExec("INSERT INTO vacancies").
		WithArgs(
			sqlmock.AnyArg(), 
			req.PublisherId,
			req.Title,
			req.Description,
			req.SalaryFrom,
			req.SalaryTo,
			req.WorkingStyles,
			req.WorkingTypes,
			req.PhoneNumber,
			req.Location,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetVacancyById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewVacanciesRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	rows := sqlmock.NewRows([]string{
		"id", "title", "description", "status", "salary_from", "salary_to", "working_styles",
		"working_types", "phone_number", "location", "view_count", "publisher_id", "publisher_name",
		"publisher_email", "publisher_phone_number", "publisher_address",
	}).
		AddRow(req.Id, "Software Engineer", "Develop software.", "active", 60000, 80000, "Remote", "Full-time",
			"1234567890", "New York", 0, uuid.New().String(), "Publisher Name", "publisher@example.com",
			"9876543210", "Publisher Address")

	mock.ExpectQuery("SELECT v.id, v.title, v.description, v.status, v.salary_from, v.salary_to, v.working_styles, v.working_types, v.phone_number, v.location, v.view_count, p.id, p.name, p.email, p.phone_number, p.address FROM vacancy v JOIN publishers p ON v.publisher_id = p.id AND p.deleted_at = 0 WHERE v.id = \\$1 AND v.deleted_at = 0").
		WithArgs(req.Id).
		WillReturnRows(rows)

	result, err := repo.GetById(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Id, result.Id)
	assert.Equal(t, "Software Engineer", result.Title)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllVacancies(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewVacanciesRepo(db)

	req := &bp.VacanciesGetAllReq{
		SalaryFrom: 60000,
		SalaryTo:   80000,
		Status:     "active",
		Page: &bp.Filter{
			Page: 1,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "title", "description", "status", "salary_from", "salary_to", "working_styles",
		"working_types", "phone_number", "location", "view_count", "publisher_id", "publisher_name",
		"publisher_email", "publisher_phone_number", "publisher_address",
	}).
		AddRow(uuid.New().String(), "Software Engineer", "Develop software.", "active", 60000, 80000, "Remote", "Full-time",
			"1234567890", "New York", 0, uuid.New().String(), "Publisher Name", "publisher@example.com",
			"9876543210", "Publisher Address").
		AddRow(uuid.New().String(), "Product Manager", "Manage product development.", "active", 70000, 90000, "Hybrid", "Full-time",
			"0987654321", "San Francisco", 5, uuid.New().String(), "Publisher Name 2", "publisher2@example.com",
			"1234567890", "Publisher Address 2")

	mock.ExpectQuery("SELECT v.id, v.title, v.description, v.status, v.salary_from, v.salary_to, v.working_styles, v.working_types, v.phone_number, v.location, v.view_count, p.id, p.name, p.email, p.phone_number, p.address FROM vacancy v JOIN publishers p ON v.publisher_id = p.id AND p.deleted_at = 0 WHERE v.deleted_at = 0 AND v.salary_from >= \\$1 AND v.salary_to <= \\$2 AND v.status = \\$3 LIMIT \\$4 OFFSET \\$5").
		WithArgs(req.SalaryFrom, req.SalaryTo, req.Status, 10, 0).
		WillReturnRows(rows)

	result, err := repo.GetAll(req)
	assert.NoError(t, err)
	assert.Len(t, result.Vacancies, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateVacancy(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewVacanciesRepo(db)

	req := &bp.VacanciesUpdateReq{
		Id: uuid.New().String(),
		Vacancy: &bp.VacanciesCreateReq{
			Title:        "Senior Software Engineer",
			Description:  "Lead software development.",
			Location:     "Los Angeles",
			PhoneNumber:  "1122334455",
			WorkingStyles: "Hybrid",
			WorkingTypes: "Full-time",
			SalaryFrom:   80000,
			SalaryTo:     100000,
			ViewCount:    1,
		},
	}

	mock.ExpectExec("UPDATE vacancies SET title = \\$1, description = \\$2, location = \\$3, phone_number = \\$4, working_styles = \\$5, working_types = \\$6, salary_from = \\$7, salary_to = \\$8, view_count = view_count \\+ \\$9, updated_at = now\\(\\) WHERE id = \\$10 AND deleted_at = 0").
		WithArgs(
			req.Vacancy.Title,
			req.Vacancy.Description,
			req.Vacancy.Location,
			req.Vacancy.PhoneNumber,
			req.Vacancy.WorkingStyles,
			req.Vacancy.WorkingTypes,
			req.Vacancy.SalaryFrom,
			req.Vacancy.SalaryTo,
			req.Vacancy.ViewCount,
			req.Id,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Update(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteVacancy(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewVacanciesRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	mock.ExpectExec("UPDATE vacancies SET status = \\'non-active' WHERE id = \\$1 AND deleted_at = 0 AND status <> \\'active'").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
