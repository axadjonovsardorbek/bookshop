package postgres_test

import (
	bp "book/genproto/books"
	"book/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatePublisher(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewPublishersRepo(db)

	req := &bp.PublishersRes{
		Name:        "Test Publisher",
		Email:       "test@example.com",
		PhoneNumber: "1234567890",
		Address:     "123 Test St",
	}

	mock.ExpectExec("INSERT INTO publishers").
		WithArgs(sqlmock.AnyArg(), req.Name, req.Email, req.PhoneNumber, req.Address).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetPublisherById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewPublishersRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "address"}).
		AddRow(req.Id, "Test Publisher", "test@example.com", "1234567890", "123 Test St")

	mock.ExpectQuery("SELECT id, name, email, phone_number, address FROM publishers WHERE deleted_at = 0 AND id = \\$1").
		WithArgs(req.Id).
		WillReturnRows(rows)

	result, err := repo.GetById(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Id, result.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllPublishers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewPublishersRepo(db)

	req := &bp.PublishersGetAllReq{
		Page: &bp.Filter{
			Page: 1,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "address"}).
		AddRow(uuid.New().String(), "Test Publisher 1", "test1@example.com", "1234567890", "123 Test St").
		AddRow(uuid.New().String(), "Test Publisher 2", "test2@example.com", "0987654321", "456 Test Ave")

	mock.ExpectQuery("SELECT id, name, email, phone_number, address FROM publishers WHERE deleted_at = 0").
		WillReturnRows(rows)

	result, err := repo.GetAll(req)
	assert.NoError(t, err)
	assert.Len(t, result.Publishers, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePublisher(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewPublishersRepo(db)

	req := &bp.PublishersUpdateReq{
		Publisher: &bp.PublishersRes{
			Id:          uuid.New().String(),
			Name:        "Updated Publisher",
			Address:     "456 Updated St",
		},
	}


	mock.ExpectExec("UPDATE publishers SET name = \\$1, address = \\$2, updated_at = now\\(\\) WHERE id = \\$3 AND deleted_at = 0").
		WithArgs(req.Publisher.Name, req.Publisher.Address, req.Publisher.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Update(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePublisher(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewPublishersRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	mock.ExpectExec("UPDATE publishers SET deleted_at = EXTRACT\\(EPOCH FROM NOW\\(\\)\\) WHERE id = \\$1 AND deleted_at = 0").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
