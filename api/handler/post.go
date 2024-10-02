package handler

import (
	pb "api-gateway/genproto/post"
	"api-gateway/pkg/minio"
	"api-gateway/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type PostHandler interface {
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	GetPostByID(c *gin.Context)
	ListPosts(c *gin.Context)
	DeletePost(c *gin.Context)
	UpdateImageToPost(c *gin.Context)
	RemoveImageFromPost(c *gin.Context)
	GetPostByCountry(c *gin.Context)
}

type postHandler struct {
	postService pb.PostServiceClient
	logger      *slog.Logger
}

func NewPostHandler(postService service.Service, logger *slog.Logger) PostHandler {
	postClient := postService.PostService()
	if postClient == nil {
		log.Fatalf("postService postService postClient is nil")
		return nil
	}
	return &postHandler{
		postService: postClient,
		logger:      logger,
	}
}

// CreatePost godoc
// @Summary Create Post
// @Description Create a new post
// @Security BearerAuth
// @Tags Posts
// @Accept json
// @Produce json
// @Param Create body models.Post true "Create post"
// @Param file formData file true "Upload image"
// @Success 201 {object} models.PostResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/create [post]
func (h *postHandler) CreatePost(c *gin.Context) {
	var post pb.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.Error("Error occurred while getting file from form", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := minio.UploadPost(file)
	if err != nil {
		h.logger.Error("Error occurred while uploading file", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.ImageUrl = url
	post.UserId = c.MustGet("user_id").(string)

	req, err := h.postService.CreatePost(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while creating post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"post": req})
}

// UpdatePost godoc
// @Summary Update Post
// @Description Update a post
// @Security BearerAuth
// @Tags Posts
// @Accept json
// @Produce json
// @Param Update body models.UpdateAPost true "Update post"
// @Success 200 {object} models.PostResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/update [put]
func (h *postHandler) UpdatePost(c *gin.Context) {
	var post pb.UpdateAPost
	if err := c.ShouldBindJSON(&post); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := h.postService.UpdatePost(c.Request.Context(), &post)
	if err != nil {
		h.logger.Error("Error occurred while updating post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": req})
}

// GetPostByID godoc
// @Summary Get Post by ID
// @Description Get a post by its ID
// @Security BearerAuth
// @Tags Posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.PostResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/getBy/{id} [get]
func (h *postHandler) GetPostByID(c *gin.Context) {
	var post pb.PostId
	post.Id = c.Param("id")
	req, err := h.postService.GetPostByID(c.Request.Context(), &post)
	if err != nil {
		h.logger.Error("Error occurred while getting post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": req})
}

// ListPosts godoc
// @Summary List Posts
// @Description Get a list of posts with optional filtering
// @Security BearerAuth
// @Tags Posts
// @Produce json
// @Param filter query models.PostList false "Filter posts"
// @Success 200 {object} models.PostListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/list [get]
func (h *postHandler) ListPosts(c *gin.Context) {
	var post pb.PostList

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
	post.Hashtag = c.Query("hashtag")

	req, err := h.postService.ListPosts(c.Request.Context(), &post)
	if err != nil {
		h.logger.Error("Error occurred while listing posts", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": req})
}

// DeletePost godoc
// @Summary Delete Post
// @Description Delete a post by its ID
// @Security BearerAuth
// @Tags Posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/delete/{id} [delete]
func (h *postHandler) DeletePost(c *gin.Context) {
	var post pb.PostId
	post.Id = c.Param("id")
	req, err := h.postService.DeletePost(c.Request.Context(), &post)
	if err != nil {
		h.logger.Error("Error occurred while deleting post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": req})
}

// AddImageToPost godoc
// @Summary Add Image to Post
// @Description Add an image to a post by post ID
// @Security BearerAuth
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "post id"
// @Param file formData file true "Upload image"
// @Success 200 {object} models.PostResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/image/{id} [put]
func (h *postHandler) UpdateImageToPost(c *gin.Context) {
	var post pb.ImageUrl

	id := c.Param("id")

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.Error("Error occurred while getting file from form", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := minio.UploadPost(file)
	if err != nil {
		h.logger.Error("Error occurred while uploading file", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Url = url
	post.PostId = id

	req, err := h.postService.AddImageToPost(c.Request.Context(), &post)
	if err != nil {
		h.logger.Error("Error occurred while adding image to post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	c.JSON(http.StatusOK, gin.H{"post": req})
}

// RemoveImageFromPost godoc
// @Summary Remove Image from Post
// @Description Remove an image from a post by post ID
// @Security BearerAuth
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Image URL"
// @Success 200 {object} models.PostResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/remove-image/{id} [delete]
func (h *postHandler) RemoveImageFromPost(c *gin.Context) {
	var post pb.ImageUrl
	post.PostId = c.Param("id")
	req, err := h.postService.RemoveImageFromPost(c.Request.Context(), &post)
	if err != nil {
		h.logger.Error("Error occurred while removing image from post", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": req})
}

// GetPostByCountry godoc
// @Summary Get Posts by Country
// @Description Get a list of posts filtered by country
// @Security BearerAuth
// @Tags Posts
// @Produce json
// @Param country path string true "Country code"
// @Success 200 {object} models.PostListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /post/country/{country} [get]
func (h *postHandler) GetPostByCountry(c *gin.Context) {
	var post pb.PostCountry

	p := c.Param("country")

	post.Country = p

	req, err := h.postService.GetPostByCountry(c.Request.Context(), &post)
	if err != nil {
		h.logger.Error("Error occurred while getting post by country", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	c.JSON(http.StatusOK, gin.H{"post": req})
}
