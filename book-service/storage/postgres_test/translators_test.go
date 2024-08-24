package postgres_test

import (
	bp "book/genproto/books"
	"book/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTranslator(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewTranslatorsRepo(db)

	req := &bp.TranslatorsRes{
		Name:    "John",
		Surname: "Doe",
	}

	mock.ExpectExec("INSERT INTO translators").
		WithArgs(sqlmock.AnyArg(), req.Name, req.Surname).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTranslatorById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewTranslatorsRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	rows := sqlmock.NewRows([]string{"id", "name", "surname"}).
		AddRow(req.Id, "John", "Doe")

	mock.ExpectQuery("SELECT id, name, surname FROM translators WHERE deleted_at = 0 AND id = \\$1").
		WithArgs(req.Id).
		WillReturnRows(rows)

	result, err := repo.GetById(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Id, result.Id)
	assert.Equal(t, "John", result.Name)
	assert.Equal(t, "Doe", result.Surname)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllTranslators(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewTranslatorsRepo(db)

	req := &bp.TranslatorsGetAllReq{
		Name:    "John",
		Surname: "Doe",
		Page: &bp.Filter{
			Page: 1,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "surname"}).
		AddRow(uuid.New().String(), "John", "Doe").
		AddRow(uuid.New().String(), "Jane", "Doe")

	mock.ExpectQuery("SELECT id, name, surname FROM translators WHERE deleted_at = 0 AND LOWER\\(name\\) LIKE LOWER\\(\\$1\\) AND LOWER\\(surname\\) LIKE LOWER\\(\\$2\\) LIMIT \\$3 OFFSET \\$4").
		WithArgs("%"+req.Name+"%", "%"+req.Surname+"%", 10, 0).
		WillReturnRows(rows)

	result, err := repo.GetAll(req)
	assert.NoError(t, err)
	assert.Len(t, result.Translators, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTranslator(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewTranslatorsRepo(db)

	req := &bp.TranslatorsUpdateReq{
		Translator: &bp.TranslatorsRes{
			Id:      uuid.New().String(),
			Name:    "John Updated",
			Surname: "Doe Updated",
		},
	}

	mock.ExpectExec("UPDATE translators SET name = \\$1, surname = \\$2, updated_at = now\\(\\) WHERE id = \\$3 AND deleted_at = 0").
		WithArgs(req.Translator.Name, req.Translator.Surname, req.Translator.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Update(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteTranslator(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewTranslatorsRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	mock.ExpectExec("UPDATE translators SET deleted_at = EXTRACT\\(EPOCH FROM NOW\\(\\)\\) WHERE id = \\$1 AND deleted_at = 0").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
