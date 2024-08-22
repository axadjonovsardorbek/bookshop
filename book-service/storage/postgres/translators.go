package postgres

import (
	"database/sql"
	bp "book/genproto/books"
)

type TranslatorsRepo struct {
	db *sql.DB
}

func NewTranslatorsRepo(db *sql.DB) *TranslatorsRepo {
	return &TranslatorsRepo{db: db}
}

func (r *TranslatorsRepo)Create(req *bp.TranslatorsRes) (*bp.Void, error){
	return nil, nil
}
func (r *TranslatorsRepo)GetById(req *bp.ById) (*bp.TranslatorsGetByIdRes, error){
	return nil, nil
}
func (r *TranslatorsRepo)GetAll(req *bp.TranslatorsGetAllReq) (*bp.TranslatorsGetAllRes, error){
	return nil, nil
}
func (r *TranslatorsRepo)Update(req *bp.TranslatorsUpdateReq) (*bp.Void, error){
	return nil, nil
}
func (r *TranslatorsRepo)Delete(req *bp.ById) (*bp.Void, error){
	return nil, nil
}