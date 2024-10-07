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
	"strconv"
)

type CommentHandler interface {
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	GetCommentByID(c *gin.Context)
	GetCommentByUsername(c *gin.Context)
	ListComments(c *gin.Context)
	DeleteComment(c *gin.Context)
	GetCommentByPostID(c *gin.Context)
	GetAllUserComments(c *gin.Context)
	GetMostlikeCommentPost(c *gin.Context)
}

type commentHandler struct {
	commentService pb.PostServiceClient
	logger         *slog.Logger
}

func NewCommentHandler(commentService service.Service, logger *slog.Logger) CommentHandler {
	commentClient := commentService.PostService()
	if commentClient == nil {
		log.Fatalf("create comment client failed")
		return nil
	}
	return &commentHandler{
		commentService: commentClient,
		logger:         logger,
	}
}

// CreateComment godoc
// @Summary Create Comment
// @Description Create a new comment
// @Security BearerAuth
// @Tags Comment
// @Accept json
// @Produce json
// @Param Create body models.CommentPost true "Create comment"
// @Success 201 {object} post.CommentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/create [post]
func (h *commentHandler) CreateComment(c *gin.Context) {
	var comment pb.CommentPost

	if err := c.ShouldBindJSON(&comment); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserId = c.MustGet("user_id").(string)
	rep, err := h.commentService.CreateComment(context.Background(), &comment)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"comment": rep})
}

// UpdateComment godoc
// @Summary Update Comment
// @Description Update a comment
// @Security BearerAuth
// @Tags Comment
// @Accept json
// @Produce json
// @Param Update body models.UpdateAComment true "Update comment"
// @Success 200 {object} models.CommentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/update [put]
func (h *commentHandler) UpdateComment(c *gin.Context) {
	var comment pb.UpdateAComment
	if err := c.ShouldBindJSON(&comment); err != nil {
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
	comment.UserId = cl["user_id"].(string)
	rep, err := h.commentService.UpdateComment(context.Background(), &comment)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"comment": rep})
}

// GetCommentByID godoc
// @Summary Get Comment by ID
// @Description Get a comment by its ID
// @Security BearerAuth
// @Tags Comment
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} models.CommentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/getBy/{id} [get]
func (h *commentHandler) GetCommentByID(c *gin.Context) {
	var commentId pb.CommentId

	commentId.CommentId = c.Param("id")

	rep, err := h.commentService.GetCommentByID(context.Background(), &commentId)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": rep})
}

// GetCommentByUsername godoc
// @Summary Get Comment by ID
// @Description Get a comment by its ID
// @Security BearerAuth
// @Tags Comment
// @Produce json
// @Success 200 {object} models.CommentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/getByUser/{id} [get]
func (h *commentHandler) GetCommentByUsername(c *gin.Context) {
	var commentId pb.Username

	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentId.Username = cl["user_id"].(string)

	rep, err := h.commentService.GetCommentByUsername(context.Background(), &commentId)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": rep})
}

// ListComments godoc
// @Summary List Comment
// @Description Get a list of Comment with optional filtering
// @Security BearerAuth
// @Tags Comment
// @Produce json
// @Param filter query models.CommentList false "Filter Comment"
// @Success 200 {object} models.CommentsR
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/list [get]
func (h *commentHandler) ListComments(c *gin.Context) {
	var comment pb.CommentList
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

	comment.Limit = int64(limits)
	comment.Offset = int64(offsets)
	comment.PostId = c.Query("id")

	commentList, err := h.commentService.ListComments(context.Background(), &comment)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": commentList})
}

// DeleteComment godoc
// @Summary Delete Comment
// @Description Delete a Comment by its ID
// @Security BearerAuth
// @Tags Comment
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/delete{id} [delete]
func (h *commentHandler) DeleteComment(c *gin.Context) {
	var commentId pb.CommentId

	commentId.CommentId = c.Param("id")

	rep, err := h.commentService.DeleteComment(context.Background(), &commentId)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": rep})
}

// GetCommentByPostID godoc
// @Summary Get Comment by Country
// @Description Get a list of Comment filtered by country
// @Security BearerAuth
// @Tags Comment
// @Produce json
// @Param id path string true "Comment code"
// @Success 200 {object} models.CommentsR
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get_comment_postId/{id} [get]
func (h *commentHandler) GetCommentByPostID(c *gin.Context) {
	var commentId pb.PostId

	commentId.Id = c.Param("id")

	rep, err := h.commentService.GetCommentByPostID(context.Background(), &commentId)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": rep})
}

// GetAllUserComments godoc
// @Summary Get Comment by Country
// @Description Get a list of Comment filtered by country
// @Security BearerAuth
// @Tags Comment
// @Produce json
// @Success 200 {object} models.CommentsR
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/country [get]
func (h *commentHandler) GetAllUserComments(c *gin.Context) {
	var comment pb.Username

	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Username = cl["user_id"].(string)
	rep, err := h.commentService.GetAllUserComments(context.Background(), &comment)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": rep})
}

// GetMostlikeCommentPost godoc
// @Summary Get Posts by Country
// @Description Get a list of posts filtered by country
// @Security BearerAuth
// @Tags Comment
// @Produce json
// @Param id path string true "Country code"
// @Success 200 {object} models.PostListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/most_like_post/{id} [get]
func (h *commentHandler) GetMostlikeCommentPost(c *gin.Context) {
	var commentId pb.PostId
	commentId.Id = c.Param("id")
	rep, err := h.commentService.GetMostlikeCommentPost(context.Background(), &commentId)
	if err != nil {
		h.logger.Error("Error occurred while creating comment", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": rep})
}
