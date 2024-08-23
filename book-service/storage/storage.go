package storage

import (
	bp "book/genproto/books"
)

type AuthorsI interface {
	Create(*bp.AuthorsRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.AuthorsRes, error)
	GetAll(*bp.AuthorsGetAllReq) (*bp.AuthorsGetAllRes, error)
	Update(*bp.AuthorsUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type BooksI interface {
	Create(*bp.BooksCreateReq) (*bp.Void, error)
	GetById(*bp.ById) (*bp.BooksRes, error)
	GetAll(*bp.BooksGetAllReq) (*bp.BooksGetAllRes, error)
	Update(*bp.BooksUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type PublishersI interface {
	Create(*bp.PublishersRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.PublishersRes, error)
	GetAll(*bp.PublishersGetAllReq) (*bp.PublishersGetAllRes, error)
	Update(*bp.PublishersUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type TranslatorsI interface {
	Create(*bp.TranslatorsRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.TranslatorsRes, error)
	GetAll(*bp.TranslatorsGetAllReq) (*bp.TranslatorsGetAllRes, error)
	Update(*bp.TranslatorsUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type VacanciesI interface {
	Create(req *bp.VacanciesCreateReq) (*bp.Void, error)
	GetById(req *bp.ById) (*bp.VacanciesRes, error)
	GetAll(req *bp.VacanciesGetAllReq) (*bp.VacanciesGetAllRes, error)
	Update(req *bp.VacanciesUpdateReq) (*bp.Void, error)
	Delete(req *bp.ById) (*bp.Void, error)
}

type LanguagesI interface {
	Create(req *bp.LanguagesRes) (*bp.Void, error)
	GetById(req *bp.ById) (*bp.LanguagesRes, error)
	GetAll(req *bp.LanguagesGetAllReq) (*bp.LanguagesGetAllRes, error)
	Update(req *bp.LanguagesUpdateReq) (*bp.Void, error)
	Delete(req *bp.ById) (*bp.Void, error)
}

type CategoriesI interface {
	Create(req *bp.CategoriesRes) (*bp.Void, error)
	GetById(req *bp.ById) (*bp.CategoriesRes, error)
	GetAll(req *bp.CategoriesGetAllReq) (*bp.CategoriesGetAllRes, error)
	Update(req *bp.CategoriesUpdateReq) (*bp.Void, error)
	Delete(req *bp.ById) (*bp.Void, error)
}
