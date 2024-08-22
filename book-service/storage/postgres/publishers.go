package postgres

import (
	"database/sql"
	bp "book/genproto/books"
)

type PublishersRepo struct {
	db *sql.DB
}

func NewPublishersRepo(db *sql.DB) *PublishersRepo {
	return &PublishersRepo{db: db}
}

func (r *PublishersRepo)Create(req *bp.PublishersRes) (*bp.Void, error){
	return nil, nil
}
func (r *PublishersRepo)GetById(req *bp.ById) (*bp.PublishersGetByIdRes, error){
	return nil, nil
}
func (r *PublishersRepo)GetAll(req *bp.PublishersGetAllReq) (*bp.PublishersGetAllRes, error){
	return nil, nil
}
func (r *PublishersRepo)Update(req *bp.PublishersUpdateReq) (*bp.Void, error){
	return nil, nil
}
func (r *PublishersRepo)Delete(req *bp.ById) (*bp.Void, error){
	return nil, nil
}