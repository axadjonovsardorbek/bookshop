package postgres_test

import (
	bp "book/genproto/books"
	"book/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCategoriesRepo(db)

	req := &bp.CategoriesRes{
		Id:          uuid.New().String(),
		Name:        "Test Category",
		Description: "Test Description",
	}

	mock.ExpectExec("INSERT INTO categories").
		WithArgs(sqlmock.AnyArg(), req.Name, req.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestGetCategoryById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCategoriesRepo(db)

	id := uuid.New().String()
	expectedCategory := &bp.CategoriesRes{
		Id:          id,
		Name:        "Test Category",
		Description: "Test Description",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(expectedCategory.Id, expectedCategory.Name, expectedCategory.Description)

	mock.ExpectQuery("SELECT id, name, description FROM categories WHERE deleted_at = 0 AND id = \\$1").
		WithArgs(id).
		WillReturnRows(rows)

	category, err := repo.GetById(&bp.ById{Id: id})
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllCategories(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCategoriesRepo(db)

	req := &bp.CategoriesGetAllReq{
		Page: &bp.Filter{Page: 1},
		Name: "Test",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(uuid.New().String(), "Category 1", "Description 1").
		AddRow(uuid.New().String(), "Category 2", "Description 2")

	mock.ExpectQuery("SELECT id, name, description FROM categories WHERE deleted_at = 0 AND LOWER\\(name\\) LIKE LOWER\\(\\$1\\) LIMIT \\$2 OFFSET \\$3").
		WithArgs("%"+req.Name+"%", int32(10), int32(0)).
		WillReturnRows(rows)

	categories, err := repo.GetAll(req)
	assert.NoError(t, err)
	assert.Len(t, categories.Categories, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCategoriesRepo(db)

	req := &bp.CategoriesUpdateReq{
		Category: &bp.CategoriesRes{
			Id:          uuid.New().String(),
			Name:        "Updated Category",
			Description: "Updated Description",
		},
	}

	mock.ExpectExec("UPDATE categories SET name = \\$1, description = \\$2, updated_at = now\\(\\) WHERE id = \\$3 AND deleted_at = 0").
		WithArgs(req.Category.Name, req.Category.Description, req.Category.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Update(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCategoriesRepo(db)

	id := uuid.New().String()

	mock.ExpectExec("UPDATE categories SET deleted_at = EXTRACT\\(EPOCH FROM NOW\\(\\)\\) WHERE id = \\$1 AND deleted_at = 0").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(&bp.ById{Id: id})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
