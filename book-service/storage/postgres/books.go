package postgres

import (
	"database/sql"
	bp "book/genproto/books"
)

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{db: db}
}

func (r *BooksRepo)Create(req *bp.BooksRes) (*bp.Void, error){
	return nil, nil
}
func (r *BooksRepo)GetById(req *bp.ById) (*bp.BooksGetByIdRes, error){
	return nil, nil
}
func (r *BooksRepo)GetAll(req *bp.BooksGetAllReq) (*bp.BooksGetAllRes, error){
	return nil, nil
}
func (r *BooksRepo)Update(req *bp.BooksUpdateReq) (*bp.Void, error){
	return nil, nil
}
func (r *BooksRepo)Delete(req *bp.ById) (*bp.Void, error){
	return nil, nil
}