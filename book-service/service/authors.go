package service

import (
	bp "book/genproto/books"
	st "book/storage/postgres"
	"context"
)

type AuthorsService struct {
	storage st.Storage
	bp.UnimplementedAuthorsServiceServer
}

func NewAuthorsService(storage *st.Storage) *AuthorsService {
	return &AuthorsService{storage: *storage}
}

func (s *AuthorsService) Create(ctx context.Context, req *bp.AuthorsRes) (*bp.Void, error){
	return s.storage.AuthorsS.Create(req)
}
func (s *AuthorsService) GetById(ctx context.Context, req *bp.ById) (*bp.AuthorsRes, error){
	return s.storage.AuthorsS.GetById(req)
}
func (s *AuthorsService) GetAll(ctx context.Context, req *bp.AuthorsGetAllReq) (*bp.AuthorsGetAllRes, error){
	return s.storage.AuthorsS.GetAll(req)
}
func (s *AuthorsService) Update(ctx context.Context, req *bp.AuthorsUpdateReq) (*bp.Void, error){
	return s.storage.AuthorsS.Update(req)
}
func (s *AuthorsService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error){
	return s.storage.AuthorsS.Delete(req)
}