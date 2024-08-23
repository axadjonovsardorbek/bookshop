package service

import (
	bp "book/genproto/books"
	st "book/storage/postgres"
	"context"
)

type BooksService struct {
	storage st.Storage
	bp.UnimplementedBooksServiceServer
}

func NewBooksService(storage *st.Storage) *BooksService {
	return &BooksService{storage: *storage}
}

func (s *BooksService) Create(ctx context.Context, req *bp.BooksCreateReq) (*bp.Void, error){
	return s.storage.BooksS.Create(req)
}
func (s *BooksService) GetById(ctx context.Context, req *bp.ById) (*bp.BooksRes, error){
	return s.storage.BooksS.GetById(req)
}
func (s *BooksService) GetAll(ctx context.Context, req *bp.BooksGetAllReq) (*bp.BooksGetAllRes, error){
	return s.storage.BooksS.GetAll(req)
}
func (s *BooksService) Update(ctx context.Context, req *bp.BooksUpdateReq) (*bp.Void, error){
	return s.storage.BooksS.Update(req)
}
func (s *BooksService) Delete(ctx context.Context, req *bp.ById) (*bp.Void, error){
	return s.storage.BooksS.Delete(req)
}