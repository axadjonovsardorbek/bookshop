package service

import (
	bp "book/genproto/books"
	st "book/storage/postgres"
	"context"
)

type PublishersService struct {
	storage st.Storage
	bp.UnimplementedPublishersServiceServer
}

func NewPublishersService(storage *st.Storage) *PublishersService {
	return &PublishersService{storage: *storage}
}

func (s *PublishersService) Create(ctx context.Context, req *bp.PublishersRes) (*bp.Void, error){
	return s.storage.PublishersS.Create(req)
}
func (s *PublishersService) GetById(ctx context.Context, req *bp.ById) (*bp.PublishersRes, error){
	return s.storage.PublishersS.GetById(req)
}
func (s *PublishersService) GetAll(ctx context.Context, req *bp.PublishersGetAllReq) (*bp.PublishersGetAllRes, error){
	return s.storage.PublishersS.GetAll(req)
}
func (s *PublishersService) Update(ctx context.Context, req *bp.PublishersUpdateReq) (*bp.Void, error){
	return s.storage.PublishersS.Update(req)
}
func (s *PublishersService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error){
	return s.storage.PublishersS.Delete(req)
}