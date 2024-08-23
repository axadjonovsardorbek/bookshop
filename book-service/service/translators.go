package service

import (
	bp "book/genproto/books"
	st "book/storage/postgres"
	"context"
)

type TranslatorsService struct {
	storage st.Storage
	bp.UnimplementedTranslatorsServiceServer
}

func NewTranslatorsService(storage *st.Storage) *TranslatorsService {
	return &TranslatorsService{storage: *storage}
}

func (s *TranslatorsService) Create(ctx context.Context, req *bp.TranslatorsRes) (*bp.Void, error){
	return s.storage.TranslatorsS.Create(req)
}
func (s *TranslatorsService) GetById(ctx context.Context, req *bp.ById) (*bp.TranslatorsRes, error){
	return s.storage.TranslatorsS.GetById(req)
}
func (s *TranslatorsService) GetAll(ctx context.Context, req *bp.TranslatorsGetAllReq) (*bp.TranslatorsGetAllRes, error){
	return s.storage.TranslatorsS.GetAll(req)
}
func (s *TranslatorsService) Update(ctx context.Context, req *bp.TranslatorsUpdateReq) (*bp.Void, error){
	return s.storage.TranslatorsS.Update(req)
}
func (s *TranslatorsService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error){
	return s.storage.TranslatorsS.Delete(req)
}