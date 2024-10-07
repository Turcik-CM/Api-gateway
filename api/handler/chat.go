package handler

import (
	pb "api-gateway/genproto/post"
	t "api-gateway/pkg/token"
	"api-gateway/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type ChatHandler interface {
	StartMessaging(c *gin.Context)
	SendMessage(c *gin.Context)
	GetChatMessages(c *gin.Context)
	MessageMarcTrue(c *gin.Context)
	GetUserChats(c *gin.Context)
	GetUnreadMessages(c *gin.Context)
	UpdateMessage(c *gin.Context)
	GetTodayMessages(c *gin.Context)
	DeleteMessage(c *gin.Context)
	DeleteChat(c *gin.Context)
}

type chatHandler struct {
	chatService pb.PostServiceClient
	logger      *slog.Logger
}

func NewChatHandler(chatService service.Service, logger *slog.Logger) ChatHandler {
	chatClient := chatService.PostService()
	if chatClient == nil {
		log.Fatalf("cannot create new chat handler")
		return nil
	}
	return &chatHandler{
		chatService: chatClient,
		logger:      logger,
	}
}

// StartMessaging godoc
// @Summary Create Chat
// @Description Create a new Chat
// @Security BearerAuth
// @Tags Chat
// @Accept json
// @Produce json
// @Param Create body models.CreateChat true "Create Chat"
// @Success 201 {object} post.ChatResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/create [post]
func (h *chatHandler) StartMessaging(c *gin.Context) {
	var chat pb.CreateChat

	if err := c.ShouldBindJSON(&chat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat.User1Id = c.MustGet("user_id").(string)
	req, err := h.chatService.StartMessaging(context.Background(), &chat)
	if err != nil {
		h.logger.Error("Error occurred while starting messaging", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": req})
}

// SendMessage godoc
// @Summary Create Chat
// @Description Create a new Chat
// @Security BearerAuth
// @Tags Chat
// @Accept json
// @Produce json
// @Param Create body models.CreateMassage true "Create Chat"
// @Success 201 {object} post.MassageResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/create_message [post]
func (h *chatHandler) SendMessage(c *gin.Context) {
	var chat pb.CreateMassage

	if err := c.ShouldBindJSON(&chat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat.SenderId = c.MustGet("user_id").(string)

	rep, err := h.chatService.SendMessage(context.Background(), &chat)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})
}

// MessageMarcTrue godoc
// @Summary Create Chat
// @Description Create a new Chat
// @Security BearerAuth
// @Tags Chat
// @Accept json
// @Produce json
// @Param Create body post.MassageTrue true "Create Chat"
// @Success 201 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/message_true [post]
func (h *chatHandler) MessageMarcTrue(c *gin.Context) {
	var chat pb.MassageTrue
	if err := c.ShouldBindJSON(&chat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rep, err := h.chatService.MessageMarcTrue(context.Background(), &chat)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})

}

// GetUserChats godoc
// @Summary Get Chat by ID
// @Description Get a chat by its ID
// @Security BearerAuth
// @Tags Chat
// @Produce json
// @Success 200 {object} post.ChatResponseList
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/get_user [get]
func (h *chatHandler) GetUserChats(c *gin.Context) {
	var user pb.Username

	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = cl["user_id"].(string)
	fmt.Println(user.Username)
	res, err := h.chatService.GetUserChats(context.Background(), &user)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"message": res})
}

// GetUnreadMessages godoc
// @Summary Get Chat by ID
// @Description Get a chat by its ID
// @Security BearerAuth
// @Tags Chat
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} post.MassageResponseList
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/get_massage/{id} [get]
func (h *chatHandler) GetUnreadMessages(c *gin.Context) {
	var user pb.ChatId
	user.ChatId = c.Param("id")
	rep, err := h.chatService.GetUnreadMessages(context.Background(), &user)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})
}

// UpdateMessage godoc
// @Summary Update Chat
// @Description Update a post
// @Security BearerAuth
// @Tags Chat
// @Accept json
// @Produce json
// @Param Update body post.UpdateMs true "Update chat"
// @Success 200 {object} post.MassageResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/update [put]
func (h *chatHandler) UpdateMessage(c *gin.Context) {
	var chat pb.UpdateMs
	if err := c.ShouldBindJSON(&chat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rep, err := h.chatService.UpdateMessage(context.Background(), &chat)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})
}

// GetTodayMessages godoc
// @Summary Get Chat by ID
// @Description Get a chat by its ID
// @Security BearerAuth
// @Tags Chat
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} post.MassageResponseList
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/get_today/{id} [get]
func (h *chatHandler) GetTodayMessages(c *gin.Context) {
	var user pb.ChatId

	user.ChatId = c.Param("id")

	rep, err := h.chatService.GetTodayMessages(context.Background(), &user)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})
}

// DeleteMessage godoc
// @Summary Delete Chat
// @Description Delete a chat by its ID
// @Security BearerAuth
// @Tags Chat
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/delete_massage/{id} [delete]
func (h *chatHandler) DeleteMessage(c *gin.Context) {
	var chat pb.MassageId
	chat.MassageId = c.Param("id")
	rep, err := h.chatService.DeleteMessage(context.Background(), &chat)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})
}

// DeleteChat godoc
// @Summary Delete Chat
// @Description Delete a chat by its ID
// @Security BearerAuth
// @Tags Chat
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/delete_chat/{id} [delete]
func (h *chatHandler) DeleteChat(c *gin.Context) {
	var chat pb.ChatId
	chat.ChatId = c.Param("id")
	rep, err := h.chatService.DeleteChat(context.Background(), &chat)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})
}

// GetChatMessages godoc
// @Summary List Chat
// @Description Get a list of chat with optional filtering
// @Security BearerAuth
// @Tags Chat
// @Produce json
// @Param filter query models.List false "Filter chat"
// @Success 200 {object} post.MassageResponseList
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /chat/list [get]
func (h *chatHandler) GetChatMessages(c *gin.Context) {
	var chat pb.List
	var offset int

	limit := c.Query("limit")
	p := c.Query("page")

	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	if page == 0 || page > 1 {
		offset = 0
	} else {
		offset = (page - 1) * limits
	}

	chat.Limit = int64(limits)
	chat.Offset = int64(offset)
	chat.ChatId = c.Query("chat_id")

	rep, err := h.chatService.GetChatMessages(context.Background(), &chat)
	if err != nil {
		h.logger.Error("Error occurred while sending message", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": rep})
}
