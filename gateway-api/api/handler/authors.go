package handler

import (
	pb "api/genproto/books"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the provided details
// @Tags authors
// @Accept  json
// @Produce  json
// @Param author body pb.AuthorsRes true "Author Details"
// @Success 200 {object} pb.AuthorsRes
// @Failure 400 {object}  string 
// @Failure 500 {object} string
// @Router /authors [post]
func (h *Handler) CreateAuthor(c *gin.Context) {
	req := pb.AuthorsRes{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Author.Create(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetByIdAuthor godoc
// @Summary Get author by ID
// @Description Get an author by their ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path string true "Author ID"
// @Success 200 {object} pb.AuthorsGetByIdRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors/{id} [get]
func (h *Handler) GetByIdAuthor(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Author.GetById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// UpdateAuthor godoc
// @Summary Update an author
// @Description Update an author's details
// @Tags authors
// @Accept  json
// @Produce  json
// @Param author body pb.AuthorsUpdateReq true "Author Update Details"
// @Success 200 {object} pb.AuthorsRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors [put]
func (h *Handler) UpdateAuthor(c *gin.Context) {
	req := pb.AuthorsUpdateReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Author.Update(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetAllAuthors godoc
// @Summary Get all authors
// @Description Get all authors with optional filtering
// @Tags authors
// @Accept  json
// @Produce  json
// @Param name query string false "Author Name"
// @Param page query int false "Page Number"
// @Success 200 {object} pb.AuthorsGetAllRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors [get]
func (h *Handler) GetAllAuthors(c *gin.Context) {
	name := c.Query("name")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with query page"})
		return
	}

	req := pb.AuthorsGetAllReq{
		Name: name,
		Page: &pb.Filter{
			Page: int32(page),
		},
	}
	res, err := h.Author.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// DeleteAuthor godoc
// @Summary Delete an author
// @Description Delete an author by ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path string true "Author ID"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors/{id} [delete]
func (h *Handler) DeleteAuthor(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Author.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}
