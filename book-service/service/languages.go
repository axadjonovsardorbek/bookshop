package service

import (
	bp "book/genproto/books"
	st "book/storage/postgres"
	"context"
)

type LanguagesService struct {
	storage st.Storage
	bp.UnimplementedLanguagesServiceServer
}

func NewLanguagesService(storage *st.Storage) *LanguagesService {
	return &LanguagesService{storage: *storage}
}

func (s *LanguagesService) Create(ctx context.Context, req *bp.LanguagesRes) (*bp.Void, error) {
	return s.storage.LanguagesS.Create(req)
}
func (s *LanguagesService) GetById(ctx context.Context, req *bp.ById) (*bp.LanguagesRes, error) {
	return s.storage.LanguagesS.GetById(req)
}
func (s *LanguagesService) GetAll(ctx context.Context, req *bp.LanguagesGetAllReq) (*bp.LanguagesGetAllRes, error) {
	return s.storage.LanguagesS.GetAll(req)
}
func (s *LanguagesService) Update(ctx context.Context, req *bp.LanguagesUpdateReq) (*bp.Void, error) {
	return s.storage.LanguagesS.Update(req)
}
func (s *LanguagesService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.LanguagesS.Delete(req)
}
