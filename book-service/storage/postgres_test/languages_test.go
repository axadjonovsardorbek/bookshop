package postgres_test

import (
	bp "book/genproto/books"
	"book/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateLanguage(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewLanguagesRepo(db)

	req := &bp.LanguagesRes{
		Name: "English",
	}

	mock.ExpectExec("INSERT INTO languages").
		WithArgs(sqlmock.AnyArg(), req.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestGetLanguageById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewLanguagesRepo(db)

	req := &bp.ById{
		Id: uuid.New().String(),
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(req.Id, "English")

	mock.ExpectQuery("SELECT id, name FROM languages WHERE deleted_at = 0 AND id = \\$1").
		WithArgs(req.Id).
		WillReturnRows(rows)

	res, err := repo.GetById(req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, req.Id, res.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllLanguages(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewLanguagesRepo(db)

	req := &bp.LanguagesGetAllReq{
		Name: "English",
		Page: &bp.Filter{Page: 1},
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(uuid.New().String(), "English").
		AddRow(uuid.New().String(), "Spanish")

	mock.ExpectQuery("SELECT id, name FROM languages WHERE deleted_at = 0 AND LOWER\\(name\\) LIKE LOWER\\(\\$1\\) LIMIT \\$2 OFFSET \\$3").
		WithArgs("%English%", 10, 0).
		WillReturnRows(rows)

	res, err := repo.GetAll(req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Languages, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateLanguage(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewLanguagesRepo(db)

	req := &bp.LanguagesUpdateReq{
		Language: &bp.LanguagesRes{
			Id:   uuid.New().String(),
			Name: "Updated Language",
		},
	}

	mock.ExpectExec("UPDATE languages SET name = \\$1, updated_at = now\\(\\) WHERE id = \\$2 AND deleted_at = 0").
		WithArgs(req.Language.Name, req.Language.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Update(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteLanguage(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewLanguagesRepo(db)

	req := &bp.ById{
		Id: uuid.New().String(),
	}

	mock.ExpectExec("UPDATE languages SET deleted_at = EXTRACT\\(EPOCH FROM NOW\\(\\)\\) WHERE id = \\$1 AND deleted_at = 0").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
