package postgres

import (
	"database/sql"
	bp "book/genproto/books"
)

type AuthorsRepo struct {
	db *sql.DB
}

func NewAuthorsRepo(db *sql.DB) *AuthorsRepo {
	return &AuthorsRepo{db: db}
}

func (r *AuthorsRepo) Create(req *bp.AuthorsRes) (*bp.Void, error) {
	return nil, nil
}
func (r *AuthorsRepo) GetById(req *bp.ById) (*bp.AuthorsGetByIdRes, error) {
	return nil, nil
}
func (r *AuthorsRepo) GetAll(req *bp.AuthorsGetAllReq) (*bp.AuthorsGetAllRes, error) {
	return nil, nil
}
func (r *AuthorsRepo) Update(req *bp.AuthorsUpdateReq) (*bp.Void, error) {
	return nil, nil
}
func (r *AuthorsRepo) Delete(req *bp.ById) (*bp.Void, error) {
	return nil, nil
}
