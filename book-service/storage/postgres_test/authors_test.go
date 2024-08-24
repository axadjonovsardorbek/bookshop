package postgres_test

import (
	bp "book/genproto/books"
	"book/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewAuthorsRepo(db)

	req := &bp.AuthorsRes{
		Id:        uuid.New().String(),
		Name:      "Test Author",
		Biography: "Test Biography",
	}

	mock.ExpectExec("INSERT INTO authors").
		WithArgs(sqlmock.AnyArg(), req.Name, req.Biography).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAuthorById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewAuthorsRepo(db)

	req := &bp.ById{Id: uuid.New().String()}
	expectedAuthor := &bp.AuthorsRes{
		Id:        req.Id,
		Name:      "Test Author",
		Biography: "Test Biography",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).
		AddRow(expectedAuthor.Id, expectedAuthor.Name, expectedAuthor.Biography)

	mock.ExpectQuery("SELECT (.+) FROM authors WHERE deleted_at = 0 AND id = \\$1").
		WithArgs(req.Id).
		WillReturnRows(rows)

	author, err := repo.GetById(req)
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthor, author)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllAuthors(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewAuthorsRepo(db)

	req := &bp.AuthorsGetAllReq{
		Name: "Test",
		Page: &bp.Filter{Page: 1},
	}

	expectedAuthor := &bp.AuthorsRes{
		Id:        uuid.New().String(),
		Name:      "Test Author",
		Biography: "Test Biography",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).
		AddRow(expectedAuthor.Id, expectedAuthor.Name, expectedAuthor.Biography)

	mock.ExpectQuery("SELECT (.+) FROM authors WHERE deleted_at = 0 AND LOWER\\(name\\) LIKE LOWER\\(\\$1\\) LIMIT \\$2 OFFSET \\$3").
		WithArgs("%"+req.Name+"%", 10, 0).
		WillReturnRows(rows)

	authors, err := repo.GetAll(req)
	assert.NoError(t, err)
	assert.Len(t, authors.Authors, 1)
	assert.Equal(t, expectedAuthor.Id, authors.Authors[0].Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewAuthorsRepo(db)

	req := &bp.AuthorsUpdateReq{
		Author: &bp.AuthorsRes{
			Id:        uuid.New().String(),
			Name:      "Updated Author",
			Biography: "Updated Biography",
		},
	}

	mock.ExpectExec("UPDATE authors SET name = \\$1, biography = \\$2, updated_at = now\\(\\) WHERE id = \\$3 AND deleted_at = 0").
	WithArgs(req.Author.Name, req.Author.Biography, req.Author.Id).
	WillReturnResult(sqlmock.NewResult(1, 1))


	_, err = repo.Update(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewAuthorsRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	mock.ExpectExec("UPDATE authors SET deleted_at = EXTRACT\\(EPOCH FROM NOW\\(\\)\\) WHERE id = \\$1 AND deleted_at = 0").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
