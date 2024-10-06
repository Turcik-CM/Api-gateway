package handler

import (
	pb "api-gateway/genproto/post"
	"api-gateway/pkg/minio"
	"api-gateway/pkg/models"
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
// @Summary Create a new post
// @Description Create a new post, including an optional image upload
// @Security BearerAuth
// @Tags Posts
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload image file (optional)"
// @Param content formData string true "Content of the post"
// @Param country formData string true "Country of the post"
// @Param description formData string true "Description of the post"
// @Param hashtag formData string true "Hashtag for the post"
// @Param location formData string true "Location for the post"
// @Param title formData string true "Title of the post"
// @Success 201 {object} models.PostResponse "Post successfully created"
// @Failure 400 {object} models.Error "Bad request, validation error or invalid file"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /post/create [post]
func (h *postHandler) CreatePost(c *gin.Context) {
	log.Println("Request received")

	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		log.Println("Failed to bind form data")
		h.logger.Error("Error occurred while binding form data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var url string
	file, err := c.FormFile("file")
	if err == nil {
		url, err = minio.UploadPost(file)
		if err != nil {
			log.Println("Error occurred while uploading file")
			h.logger.Error("Error occurred while uploading file:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		post.ImageUrl = url
	}

	userId, exists := c.Get("user_id")
	if !exists {
		log.Println("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	post.UserId = userId.(string)

	resp := pb.Post{
		Location:    post.Location,
		Country:     post.Country,
		ImageUrl:    post.ImageUrl,
		UserId:      post.UserId,
		Title:       post.Title,
		Description: post.Description,
		Hashtag:     post.Hashtag,
		Content:     post.Content,
	}

	req, err := h.postService.CreatePost(context.Background(), &resp)
	if err != nil {
		log.Println("Error occurred while creating post in service")
		h.logger.Error("Error occurred while creating post:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created post response
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
	fmt.Println(req)
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
