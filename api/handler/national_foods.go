package handler

import (
	pb "api-gateway/genproto/nationality"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type NationalFoodHandler interface {
	CreateNationalFood(c *gin.Context)
	UpdateNationalFood(c *gin.Context)
	GetNationalFoodByID(c *gin.Context)
	DeleteNationalFood(c *gin.Context)
	ListNationalFoods(c *gin.Context)
	AddImageUrll(c *gin.Context)
}

type nationalFoodHandler struct {
	nationalFoodService pb.NationalityServiceClient
	logger              *slog.Logger
}

func NewNationalFoodHandler(service service.Service, logger *slog.Logger) NationalFoodHandler {
	Clinet := service.Nationality()
	return &nationalFoodHandler{
		nationalFoodService: Clinet,
		logger:              logger,
	}
}

func (h *nationalFoodHandler) CreateNationalFood(c *gin.Context) {
	var nat pb.NationalFood
	if err := c.ShouldBindJSON(&nat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.nationalFoodService.CreateNationalFood(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while creating national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"response": resp})
}

func (h *nationalFoodHandler) UpdateNationalFood(c *gin.Context) {
	var nat pb.UpdateNationalFood
	if err := c.ShouldBindJSON(&nat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.nationalFoodService.UpdateNationalFoods(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while updating national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"response": resp})
}

func (h *nationalFoodHandler) GetNationalFoodByID(c *gin.Context) {
	id := c.Param("id")
	nat := pb.NationalFoodId{
		Id: id,
	}
	resp, err := h.nationalFoodService.GetNationalFoodByID(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while getting national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}

func (h *nationalFoodHandler) DeleteNationalFood(c *gin.Context) {
	id := c.Param("id")
	nat := pb.NationalFoodId{
		Id: id,
	}
	resp, err := h.nationalFoodService.DeleteNationalFood(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while deleting national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}

func (h *nationalFoodHandler) ListNationalFoods(c *gin.Context) {
	var post pb.NationalFoodList

	limit := c.Query("limit")
	offset := c.Query("offset")

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		offsets = 1
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	post.Limit = int64(limits)
	post.Offset = int64(offsets)

	post.Country = c.Query("country")
	resp, err := h.nationalFoodService.ListNationalFood(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while listing national foods", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}

func (h *nationalFoodHandler) AddImageUrll(c *gin.Context) {
	var nat pb.NationalFoodImage
	if err := c.ShouldBindJSON(&nat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.nationalFoodService.AddNationalFoodImage(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while adding national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"response": resp})
}
