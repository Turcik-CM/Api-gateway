package handler

import (
	"api-gateway/service"
	"log/slog"
)

type Handler struct {
	Post    PostHandler
	Chat    ChatHandler
	Like    LikeHandler
	Comment CommentHandler
	User    UserHandler
}

func (h *Handler) ChatHandler() ChatHandler {
	return h.Chat
}

func (h *Handler) CommentHandler() CommentHandler {
	return h.Comment
}

func (h *Handler) UserHandler() UserHandler {
	return h.User
}

func (h *Handler) LikeHandler() LikeHandler {
	return h.Like
}

func (h *Handler) PostHandler() PostHandler {
	return h.Post
}

type SHandler interface {
	PostHandler() PostHandler
	CommentHandler() CommentHandler
	UserHandler() UserHandler
	LikeHandler() LikeHandler
	ChatHandler() ChatHandler
}

func NewMainHandler(Service service.Service, logger *slog.Logger) SHandler {
	return &Handler{
		Post:    NewPostHandler(Service, logger),
		Chat:    NewChatHandler(Service, logger),
		Comment: NewCommentHandler(Service, logger),
		Like:    NewLikeHandler(Service, logger),
		User:    NewUserHandler(Service, logger),
	}
}
