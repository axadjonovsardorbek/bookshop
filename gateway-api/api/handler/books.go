package handler

import (
	pb "api/genproto/books"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBooks godoc
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body pb.BooksRes true "Book Details"
// @Success 200 {object} pb.BooksRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books [post]
func (h *Handler) CreateBooks(c *gin.Context) {
	req := pb.BooksRes{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	
	res, err := h.Book.Create(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetByIdBook godoc
// @Summary Get book by ID
// @Description Get a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} pb.BooksRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books/{id} [get]
func (h *Handler) GetByIdBook(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Book.GetById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book's details
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body pb.BooksUpdateReq true "Book Update Details"
// @Success 200 {object} pb.BooksRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	req := pb.BooksUpdateReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Book.Update(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get all books with optional filtering
// @Tags books
// @Accept  json
// @Produce  json
// @Param price_from query int false "Minimum Price"
// @Param price_to query int false "Maximum Price"
// @Param page query int false "Page Number"
// @Success 200 {object} pb.BooksGetAllRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books [get]
func (h *Handler) GetAllBook(c *gin.Context) {
	priceFrom, err := strconv.Atoi(c.Query("price_from"))
	if err != nil {
		c.JSON(400, gin.H{"Error": "Error with query price_from"})
		return
	}
	priceTo, err := strconv.Atoi(c.Query("price_to"))
	if err != nil {
		c.JSON(400, gin.H{"Error": "Error with query price_to"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(400, gin.H{"Error": "Error with query page"})
		return
	}
	req := pb.BooksGetAllReq{
		PriceFrom: int32(priceFrom),
		PriceTo:   int32(priceTo),
		Page: &pb.Filter{
			Page: int32(page),
		},
	}
	res, err := h.Book.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// DeleteBooks godoc
// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books/{id} [delete]
func (h *Handler) DeleteBooks(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Book.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}
