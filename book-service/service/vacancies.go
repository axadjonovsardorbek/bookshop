package service

import (
	bp "book/genproto/books"
	st "book/storage/postgres"
	"context"
)

type VacanciesService struct {
	storage st.Storage
	bp.UnimplementedVacanciesServiceServer
}

func NewVacanciesService(storage *st.Storage) *VacanciesService {
	return &VacanciesService{storage: *storage}
}

func (s *VacanciesService) Create(ctx context.Context, req *bp.VacanciesCreateReq) (*bp.Void, error) {
	return s.storage.VacanciesS.Create(req)
}
func (s *VacanciesService) GetById(ctx context.Context, req *bp.ById) (*bp.VacanciesRes, error) {
	return s.storage.VacanciesS.GetById(req)
}
func (s *VacanciesService) GetAll(ctx context.Context, req *bp.VacanciesGetAllReq) (*bp.VacanciesGetAllRes, error) {
	return s.storage.VacanciesS.GetAll(req)
}
func (s *VacanciesService) Update(ctx context.Context, req *bp.VacanciesUpdateReq) (*bp.Void, error) {
	return s.storage.VacanciesS.Update(req)
}
func (s *VacanciesService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error) {
	return s.storage.VacanciesS.Delete(req)
}
