package handler

import (
	pb "api/genproto/books"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePublisher godoc
// @Summary Create a new publisher
// @Description Create a new publisher with the provided details
// @Tags publishers
// @Accept  json
// @Produce  json
// @Param publisher body pb.PublishersRes true "Publisher Details"
// @Success 200 {object} pb.PublishersRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /publishers [post]
func (h *Handler) CreatePublisher(c *gin.Context) {
	req := pb.PublishersRes{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Publisher.Create(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetByIdPublisher godoc
// @Summary Get publisher by ID
// @Description Get a publisher by its ID
// @Tags publishers
// @Accept  json
// @Produce  json
// @Param id path string true "Publisher ID"
// @Success 200 {object} pb.PublishersRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /publishers/{id} [get]
func (h *Handler) GetByIdPublisher(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Publisher.GetById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// UpdatePublisher godoc
// @Summary Update a publisher
// @Description Update a publisher's details
// @Tags publishers
// @Accept  json
// @Produce  json
// @Param publisher body pb.PublishersUpdateReq true "Publisher Update Details"
// @Success 200 {object} pb.PublishersRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /publishers [put]
func (h *Handler) UpdatePublisher(c *gin.Context) {
	req := pb.PublishersUpdateReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Publisher.Update(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetAllPublishers godoc
// @Summary Get all publishers
// @Description Get all publishers with optional filtering
// @Tags publishers
// @Accept  json
// @Produce  json
// @Param name query string false "Publisher Name"
// @Param address query string false "Publisher Address"
// @Param page query int false "Page Number"
// @Success 200 {object} pb.PublishersGetAllRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /publishers [get]
func (h *Handler) GetAllPublishers(c *gin.Context) {
	name := c.Query("name")
	address := c.Query("address")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with query page"})
		return
	}

	req := pb.PublishersGetAllReq{
		Name:    name,
		Address: address,
		Page: &pb.Filter{
			Page: int32(page),
		},
	}
	res, err := h.Publisher.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// DeletePublisher godoc
// @Summary Delete a publisher
// @Description Delete a publisher by ID
// @Tags publishers
// @Accept  json
// @Produce  json
// @Param id path string true "Publisher ID"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /publishers/{id} [delete]
func (h *Handler) DeletePublisher(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Publisher.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}
