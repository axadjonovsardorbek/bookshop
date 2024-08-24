package postgres_test

import (
	bp "book/genproto/books"
	"book/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewBooksRepo(db)

	req := &bp.BooksCreateReq{
		ProviderId:  uuid.New().String(),
		PublisherId: uuid.New().String(),
		CategoryId:  uuid.New().String(),
		TranslatorId: uuid.New().String(),
		AuthorId: uuid.New().String(),
		LanguageId: uuid.New().String(),
		Title:       "Book Title",
		Description: "Book Description",
		PublishedYear: "2024",
		TotalPages:  200,
		Price:       29.99,
		Stock:       10,
		ImageUrl:    "http://example.com/image.jpg",
		WritingType: "Fiction",
	}

	mock.ExpectExec(`INSERT INTO vacancies `).
		WithArgs(
			sqlmock.AnyArg(),
			req.ProviderId,
			req.PublisherId,
			req.CategoryId,
			req.TranslatorId,
			req.AuthorId,
			req.LanguageId,
			req.Title,
			req.Description,
			req.PublishedYear,
			req.TotalPages,
			req.Price,
			req.Stock,
			req.ImageUrl,
			req.WritingType,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestUpdateBook(t *testing.T) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    repo := postgres.NewBooksRepo(db)

    req := &bp.BooksUpdateReq{
        Id: uuid.New().String(),
        Book: &bp.BooksCreateReq{
            Title:       "Updated Title",
            Description: "Updated Description",
            Price:       35.99,
            Stock:       15,
            ViewCount:   1,
        },
    }

	mock.ExpectExec(`UPDATE books\s+SET\s+title\s*=\s*\$1,\s*description\s*=\s*\$2,\s*stock\s*=\s*\$3,\s*price\s*=\s*\$4,\s*view_count\s*=\s*view_count\s*\+\s*\$5,\s*updated_at\s*=\s*now\(\)\s+WHERE\s+id\s*=\s*\$6\s+AND\s+deleted_at\s*=\s*0`).
    WithArgs(
        req.Book.Title,
        req.Book.Description,
        req.Book.Stock,
        req.Book.Price,
        req.Book.ViewCount,
        req.Id,
    ).
    WillReturnResult(sqlmock.NewResult(1, 1))




    _, err = repo.Update(req)
    if err != nil {
        t.Fatalf("Error during update: %v", err)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("Unmet expectations: %v", err)
    }
}


func TestDeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewBooksRepo(db)

	req := &bp.ById{Id: uuid.New().String()}

	mock.ExpectExec("UPDATE books SET deleted_at = EXTRACT\\(EPOCH FROM NOW\\(\\)\\) WHERE id = \\$1 AND deleted_at = 0").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
