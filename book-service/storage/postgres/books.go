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

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{db: db}
}

func (r *BooksRepo) Create(req *bp.BooksCreateReq) (*bp.Void, error) {

	id := uuid.New().String()

	query := `
	INSERT INTO vacancies (
		id,
		provider_id,
		publisher_id,
		category_id,
		translator_id,
		author_id,
		language_id,
		title,
		description,
		published_year,
		total_pages,
		price,
		stock,
		image_url,
		writing_type
	) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`

	_, err := r.db.Exec(
		query,
		id,
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
	)
	if err != nil {
		log.Println("Error while creating books: ", err)
		return nil, err
	}

	log.Println("Successfully created books")

	return nil, nil
}
func (r *BooksRepo) GetById(req *bp.ById) (*bp.BooksRes, error) {
	book := bp.BooksRes{
		Provider: &bp.UsersRes{},
		Publisher: &bp.PublishersRes{},
		Category: &bp.CategoriesRes{},
		Translator: &bp.TranslatorsRes{},
		Author: &bp.AuthorsRes{},
		Language: &bp.LanguagesRes{},
	}  

	query := `
	SELECT 
		b.id,
		b.title,
		b.description,
		b.published_year,
		b.total_pages,
		b.price,
		b.stock,
		b.image_url,
		b.writing_type,
		b.view_count,
		u.id,
		u.first_name,
		u.last_name,
		u.email,
		u.phone_number
		p.id,
		p.name,
		p.email,
		p.phone_number,
		p.address,
		c.id,
		c.name,
		c.description,
		t.id,
		t.name,
		t.surname,
		a.id,
		a.name,
		a.biography,
		l.id,
		l.name
	FROM 
		books b
	JOIN 
		users u
	ON 
		b.provider_id = u.id AND u.deleted_at = 0
	JOIN 
		publishers p
	ON
		b.publisher_id = p.id AND p.deleted_at = 0
	JOIN 
		categories c
	ON 
		b.category_id = c.id AND c.deleted_at = 0
	JOIN 
		translators t
	ON 	
		b.translator_id = t.id AND t.deleted_at = 0
	JOIN 
		authors a
	ON
		b.author_id = a.id AND a.deleted_at = 0
	JOIN 
		languages l
	ON
		b.language_id = l.id AND l.deleted_at = 0
	WHERE 
		b.id = $1 AND b.deleted_at = 0
	` 

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&book.Id,
		&book.Title,
		&book.Description,
		&book.PublishedYear,
		&book.TotalPages,
		&book.Price,
		&book.Stock,
		&book.ImageUrl,
		&book.WritingType,
		&book.ViewCount,
		&book.Provider.Id,
		&book.Provider.FirstName,
		&book.Provider.LastName,
		&book.Provider.Email,
		&book.Provider.PhoneNumber,
		&book.Publisher.Id,
		&book.Publisher.Name,
		&book.Publisher.Email,
		&book.Publisher.PhoneNumber,
		&book.Publisher.Address,
		&book.Category.Id,
		&book.Category.Name,
		&book.Category.Description,
		&book.Translator.Id,
		&book.Translator.Name,
		&book.Translator.Surname,
		&book.Author.Id,
		&book.Author.Name,
		&book.Author.Biography,
		&book.Language.Id,
		&book.Language.Name,
	)

	if err != nil {
		log.Println("Error while getting book by id: ", err)
		return nil, err
	}

	log.Println("Successfully got book")

	return &book, nil
}
func (r *BooksRepo) GetAll(req *bp.BooksGetAllReq) (*bp.BooksGetAllRes, error) {
	books := bp.BooksGetAllRes{}

	query := `
	SELECT 
		b.id,
		b.title,
		b.description,
		b.published_year,
		b.total_pages,
		b.price,
		b.stock,
		b.image_url,
		b.writing_type,
		b.view_count,
		u.id,
		u.first_name,
		u.last_name,
		u.email,
		u.phone_number
		p.id,
		p.name,
		p.email,
		p.phone_number,
		p.address,
		c.id,
		c.name,
		c.description,
		t.id,
		t.name,
		t.surname,
		a.id,
		a.name,
		a.biography,
		l.id,
		l.name
	FROM 
		books b
	JOIN 
		users u
	ON 
		b.provider_id = u.id AND u.deleted_at = 0
	JOIN 
		publishers p
	ON
		b.publisher_id = p.id AND p.deleted_at = 0
	JOIN 
		categories c
	ON 
		b.category_id = c.id AND c.deleted_at = 0
	JOIN 
		translators t
	ON 	
		b.translator_id = t.id AND t.deleted_at = 0
	JOIN 
		authors a
	ON
		b.author_id = a.id AND a.deleted_at = 0
	JOIN 
		languages l
	ON
		b.language_id = l.id AND l.deleted_at = 0
	WHERE 
		b.deleted_at = 0
	` 

	var args []interface{}
	var conditions []string

	if req.PriceFrom >= 0{
		conditions = append(conditions, " b.price >= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.PriceFrom)
	}
	if req.PriceTo > 0{
		conditions = append(conditions, " b.price <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.PriceTo)
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
		log.Println("books not found")
		return nil, errors.New("books list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving books: ", err)
		return nil, err
	}

	for rows.Next() {
		book := bp.BooksRes{
			Provider: &bp.UsersRes{},
			Publisher: &bp.PublishersRes{},
			Category: &bp.CategoriesRes{},
			Translator: &bp.TranslatorsRes{},
			Author: &bp.AuthorsRes{},
			Language: &bp.LanguagesRes{},
		}

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.PublishedYear,
			&book.TotalPages,
			&book.Price,
			&book.Stock,
			&book.ImageUrl,
			&book.WritingType,
			&book.ViewCount,
			&book.Provider.Id,
			&book.Provider.FirstName,
			&book.Provider.LastName,
			&book.Provider.Email,
			&book.Provider.PhoneNumber,
			&book.Publisher.Id,
			&book.Publisher.Name,
			&book.Publisher.Email,
			&book.Publisher.PhoneNumber,
			&book.Publisher.Address,
			&book.Category.Id,
			&book.Category.Name,
			&book.Category.Description,
			&book.Translator.Id,
			&book.Translator.Name,
			&book.Translator.Surname,
			&book.Author.Id,
			&book.Author.Name,
			&book.Author.Biography,
			&book.Language.Id,
			&book.Language.Name,
		)

		if err != nil {
			log.Println("Error while scanning all books: ", err)
			return nil, err
		}

		books.Books = append(books.Books, &book)
	}

	log.Println("Successfully fetched all books")
	return &books, nil
}
func (r *BooksRepo) Update(req *bp.BooksUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		books
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Book.Title != "" && req.Book.Title != "string" {
		conditions = append(conditions, " title = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.Title)
	}
	if req.Book.Description != "" && req.Book.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.Description)
	}
	if req.Book.PublishedYear != "" && req.Book.PublishedYear != "string" {
		conditions = append(conditions, " published_year = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.PublishedYear)
	}
	if req.Book.ImageUrl != "" && req.Book.ImageUrl != "string" {
		conditions = append(conditions, " image_url = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.ImageUrl)
	}
	if req.Book.WritingType != "" && req.Book.WritingType != "string" {
		conditions = append(conditions, " writing_type = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.WritingType)
	}
	if req.Book.Stock > 0 {
		conditions = append(conditions, " stock = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.Stock)
	}
	if req.Book.Price > 0 {
		conditions = append(conditions, " price = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.Price)
	}
	if req.Book.TotalPages > 0 {
		conditions = append(conditions, " total_pages = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.TotalPages)
	}
	if req.Book.ViewCount == 1 {
		conditions = append(conditions, " view_count = view_count + $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Book.ViewCount)
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
		log.Println("Error while updating book: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("book with id %s couldn't update", req.Id)
	}

	log.Println("Successfully updated book")

	return nil, nil

}
func (r *BooksRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		books
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting book: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("book with id %s not found", req.Id)
	}

	log.Println("Successfully deleted book")

	return nil, nil
}
