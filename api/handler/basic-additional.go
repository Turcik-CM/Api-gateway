package handler

import (
	pb "api-gateway/genproto/post"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
)

type BasicAdditional interface {
	GetUserRecommendation(c *gin.Context)
	GetPostsByUsername(c *gin.Context)
	SearchPost(c *gin.Context)
}

type basicAdditional struct {
	basic  pb.PostServiceClient
	logger *slog.Logger
}

func NewBasicAdditional(service service.Service, logger *slog.Logger) BasicAdditional {
	basicClent := service.PostService()
	if basicClent == nil {
		log.Fatalf("Error creating Basic additional service")
		return nil
	}
	return &basicAdditional{
		basic:  basicClent,
		logger: logger,
	}
}

// GetUserRecommendation godoc
// @Summary Get Recommendation by ID
// @Description Get Recommendation by its ID
// @Security BearerAuth
// @Tags Posts
// @Produce json
// @Param user_name path string true "Basic ID"
// @Success 200 {object} models.PostListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /basic/username/{user_name} [get]
func (b *basicAdditional) GetUserRecommendation(c *gin.Context) {
	user := c.Param("user_name")

	res := pb.Username{
		Username: user,
	}

	req, err := b.basic.GetUserRecommendation(context.Background(), &res)
	if err != nil {
		b.logger.Error("Error getting recommendation")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetPostsByUsername godoc
// @Summary Get PostsByUsername by ID
// @Description Get PostsByUsername by its ID
// @Security BearerAuth
// @Tags Posts
// @Produce json
// @Param user_name path string true "Basic ID"
// @Success 200 {object} models.PostListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /basic/getByUser/{user_name} [get]
func (b *basicAdditional) GetPostsByUsername(c *gin.Context) {
	user := c.Param("user_name")

	res := pb.Username{
		Username: user,
	}

	req, err := b.basic.GetPostsByUsername(context.Background(), &res)
	if err != nil {
		b.logger.Error("Error getting recommendation")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// SearchPost godoc
// @Summary Get SearchPost by ID
// @Description Get SearchPost by its ID
// @Security BearerAuth
// @Tags Posts
// @Produce json
// @Param action path string true "SearchPost ID"
// @Success 200 {object} models.PostListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /basic/search/{action} [get]
func (b *basicAdditional) SearchPost(c *gin.Context) {
	var action pb.Search

	s := c.Param("action")

	action.Action = s

	res, err := b.basic.SearchPost(context.Background(), &action)
	if err != nil {
		b.logger.Error("Error getting recommendation")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
