package service

import (
	bp "book/genproto/books"
	st "book/storage/postgres"
	"context"
)

type CategoriesService struct {
	storage st.Storage
	bp.UnimplementedCategoriesServiceServer
}

func NewCategoriesService(storage *st.Storage) *CategoriesService {
	return &CategoriesService{storage: *storage}
}

func (s *CategoriesService) Create(ctx context.Context, req *bp.CategoriesRes) (*bp.Void, error) {
	return s.storage.CategoriesS.Create(req)
}
func (s *CategoriesService) GetById(ctx context.Context, req *bp.ById) (*bp.CategoriesRes, error) {
	return s.storage.CategoriesS.GetById(req)
}
func (s *CategoriesService) GetAll(ctx context.Context, req *bp.CategoriesGetAllReq) (*bp.CategoriesGetAllRes, error) {
	return s.storage.CategoriesS.GetAll(req)
}
func (s *CategoriesService) Update(ctx context.Context, req *bp.CategoriesUpdateReq) (*bp.Void, error) {
	return s.storage.CategoriesS.Update(req)
}
func (s *CategoriesService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.CategoriesS.Delete(req)
}
