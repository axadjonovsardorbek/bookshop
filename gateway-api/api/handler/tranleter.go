package handler

import (
	pb "api/genproto/books"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTranslator godoc
// @Summary Create a new translator
// @Description Create a new translator with the provided details
// @Tags translators
// @Accept  json
// @Produce  json
// @Param translator body pb.TranslatorsRes true "Translator Details"
// @Success 200 {object} pb.TranslatorsRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /translators [post]
func (h *Handler) CreateTranslator(c *gin.Context) {
	req := pb.TranslatorsRes{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Translater.Create(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetByIdTranslator godoc
// @Summary Get translator by ID
// @Description Get a translator by its ID
// @Tags translators
// @Accept  json
// @Produce  json
// @Param id path string true "Translator ID"
// @Success 200 {object} pb.TranslatorsRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /translators/{id} [get]
func (h *Handler) GetByIdTranslator(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Translater.GetById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// UpdateTranslator godoc
// @Summary Update a translator
// @Description Update a translator's details
// @Tags translators
// @Accept  json
// @Produce  json
// @Param translator body pb.TranslatorsUpdateReq true "Translator Update Details"
// @Success 200 {object} pb.TranslatorsRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /translators [put]
func (h *Handler) UpdateTranslator(c *gin.Context) {
	req := pb.TranslatorsUpdateReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Translater.Update(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// GetAllTranslators godoc
// @Summary Get all translators
// @Description Get all translators with optional filtering
// @Tags translators
// @Accept  json
// @Produce  json
// @Param name query string false "Translator Name"
// @Param surname query string false "Translator Surname"
// @Param page query int false "Page Number"
// @Success 200 {object} pb.TranslatorsGetAllRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /translators [get]
func (h *Handler) GetAllTranslators(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with query page"})
		return
	}

	req := pb.TranslatorsGetAllReq{
		Name:    name,
		Surname: surname,
		Page: &pb.Filter{
			Page: int32(page),
		},
	}
	res, err := h.Translater.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}

// DeleteTranslator godoc
// @Summary Delete a translator
// @Description Delete a translator by ID
// @Tags translators
// @Accept  json
// @Produce  json
// @Param id path string true "Translator ID"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /translators/{id} [delete]
func (h *Handler) DeleteTranslator(c *gin.Context) {
	req := pb.ById{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	res, err := h.Translater.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(200, res)
}
