package storage

import (
	bp "book/genproto/books"
)

type AuthorsI interface{
    Create(*bp.AuthorsRes) (*bp.Void, error)
    GetById(*bp.ById) (*bp.AuthorsGetByIdRes, error)
    GetAll(*bp.AuthorsGetAllReq) (*bp.AuthorsGetAllRes, error)
    Update(*bp.AuthorsUpdateReq) (*bp.Void, error)
    Delete(*bp.ById) (*bp.Void, error)
}

type BooksI interface{
    Create(*bp.BooksRes) (*bp.Void, error)
    GetById(*bp.ById) (*bp.BooksGetByIdRes, error)
    GetAll(*bp.BooksGetAllReq) (*bp.BooksGetAllRes, error)
    Update(*bp.BooksUpdateReq) (*bp.Void, error)
    Delete(*bp.ById) (*bp.Void, error)
}

type PublishersI interface{
    Create(*bp.PublishersRes) (*bp.Void, error)
    GetById(*bp.ById) (*bp.PublishersGetByIdRes, error)
    GetAll(*bp.PublishersGetAllReq) (*bp.PublishersGetAllRes, error)
    Update(*bp.PublishersUpdateReq) (*bp.Void, error)
    Delete(*bp.ById) (*bp.Void, error)
}

type TranslatorsI interface{
    Create(*bp.TranslatorsRes) (*bp.Void, error)
    GetById(*bp.ById) (*bp.TranslatorsGetByIdRes, error)
    GetAll(*bp.TranslatorsGetAllReq) (*bp.TranslatorsGetAllRes, error)
    Update(*bp.TranslatorsUpdateReq) (*bp.Void, error)
    Delete(*bp.ById) (*bp.Void, error)
}