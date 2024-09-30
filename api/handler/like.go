package handler

import (
	pb "api-gateway/genproto/post"
	t "api-gateway/pkg/token"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
)

type LikeHandler interface {
	AddLikePost(c *gin.Context)
	DeleteLikePost(c *gin.Context)
	AddLikeComment(c *gin.Context)
	DeleteLikeComment(c *gin.Context)
	GetPostLikeCount(c *gin.Context)
	GetMostLikedComment(c *gin.Context)
	GetUsersWhichLikePost(c *gin.Context)
	GetUsersWhichLikeComment(c *gin.Context)
}

type HandlerL struct {
	LikeService pb.PostServiceClient
	logger      *slog.Logger
}

func NewLikeHandler(likeService service.Service, logger *slog.Logger) LikeHandler {
	likeClient := likeService.PostService()
	if likeClient == nil {
		log.Fatalf("Error creating like handler")
		return nil
	}
	return &HandlerL{
		LikeService: likeClient,
		logger:      logger,
	}
}

// AddLikePost godoc
// @Summary Create Like
// @Description Create a new like
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param like body models.LikePost true "Like Post"
// @Success 200 {object} models.LikeResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/create [post]
func (h *HandlerL) AddLikePost(c *gin.Context) {
	var like pb.LikePost
	if err := c.ShouldBindJSON(&like); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	like.UserId = cl["user_id"].(string)

	req, err := h.LikeService.AddLikePost(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while adding like post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// DeleteLikePost godoc
// @Summary Delete Like from a Post
// @Description Delete a like by its ID
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param id path string true "Like ID"
// @Param like body models.LikePost true "Like Post"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/delete/{id} [delete]
func (h *HandlerL) DeleteLikePost(c *gin.Context) {
	var like pb.LikePost

	like.PostId = c.PostForm("id")

	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	like.UserId = cl["user_id"].(string)

	req, err := h.LikeService.DeleteLikePost(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while deleting like post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// AddLikeComment godoc
// @Summary Add Like to a Comment
// @Description Add a like to a comment by comment ID
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param like body models.LikeCommit true "Like Comment"
// @Success 200 {object} models.LikeComResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/comment/create [post]
func (h *HandlerL) AddLikeComment(c *gin.Context) {
	var like pb.LikeCommit
	if err := c.ShouldBindJSON(&like); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	like.UserId = cl["user_id"].(string)
	req, err := h.LikeService.AddLikeComment(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while adding like comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// DeleteLikeComment godoc
// @Summary Delete Like from a Comment
// @Description Remove a like from a comment by comment ID
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param commit_id path string true "Comment ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/comment/delete/{commit_id} [delete]
func (h *HandlerL) DeleteLikeComment(c *gin.Context) {
	var like pb.LikeCommit

	like.CommitId = c.Param("commit_id")

	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	like.UserId = cl["user_id"].(string)
	req, err := h.LikeService.DeleteLikeComment(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while deleting like comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetPostLikeCount godoc
// @Summary Get Post Like Count
// @Description Retrieve the total number of likes for a given post
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 200 {object} models.LikeCount
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/post/count/{post_id} [get]
func (h *HandlerL) GetPostLikeCount(c *gin.Context) {
	var like pb.PostId

	like.Id = c.Param("post_id")

	req, err := h.LikeService.GetPostLikeCount(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while getting post like count", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetMostLikedComment godoc
// @Summary Get Most Liked Comment
// @Description Retrieve the comment with the most likes for a given post
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 200 {object} models.LikeCount
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/comment/most-liked/{post_id} [get]
func (h *HandlerL) GetMostLikedComment(c *gin.Context) {
	var like pb.PostId

	like.Id = c.Param("post_id")

	req, err := h.LikeService.GetMostLikedComment(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while getting most liked comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetUsersWhichLikePost godoc
// @Summary Get Users Who Liked a Post
// @Description Retrieve a list of users who liked a specific post
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 200 {object} models.Users
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/post/users/{post_id} [get]
func (h *HandlerL) GetUsersWhichLikePost(c *gin.Context) {
	var like pb.PostId

	like.Id = c.Param("post_id")

	req, err := h.LikeService.GetUsersWhichLikePost(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while getting users like post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetUsersWhichLikeComment godoc
// @Summary Get Users Who Liked a Comment
// @Description Retrieve a list of users who liked a specific comment
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param comment_id path string true "Comment ID"
// @Success 200 {object} models.Users
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/comment/users/{comment_id} [get]
func (h *HandlerL) GetUsersWhichLikeComment(c *gin.Context) {
	var like pb.CommentId

	like.CommentId = c.Param("comment_id")

	req, err := h.LikeService.GetUsersWhichLikeComment(context.Background(), &like)
	if err != nil {
		h.logger.Error("Error occurred while getting users like comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}
