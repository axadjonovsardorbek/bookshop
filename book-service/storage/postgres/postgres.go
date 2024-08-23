package postgres

import (
	"book/config"
	"book/storage"
	"database/sql"
	"fmt"

	// "github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db           *sql.DB
	AuthorsS     storage.AuthorsI
	BooksS       storage.BooksI
	CategoriesS  storage.CategoriesI
	LanguagesS   storage.LanguagesI
	VacanciesS   storage.VacanciesI
	PublishersS  storage.PublishersI
	TranslatorsS storage.TranslatorsI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// rAddr := fmt.Sprintf("%s%s", config.REDIS_HOST, config.REDIS_PORT)

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: rAddr,
	// })

	author := NewAuthorsRepo(db)
	book := NewBooksRepo(db)
	publisher := NewPublishersRepo(db)
	translator := NewTranslatorsRepo(db)
	category := NewCategoriesRepo(db)
	language := NewLanguagesRepo(db)
	vacancy := NewVacanciesRepo(db)

	return &Storage{
		Db:           db,
		AuthorsS:     author,
		BooksS:       book,
		PublishersS:  publisher,
		TranslatorsS: translator,
		CategoriesS:  category,
		LanguagesS:   language,
		VacanciesS:   vacancy,
	}, nil
}
