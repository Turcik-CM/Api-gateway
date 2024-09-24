package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"
	"api-gateway/pkg/config"
	"api-gateway/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

func NewRouter(cfg *config.Config, log *slog.Logger, casbin *casbin.Enforcer) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.PermissionMiddleware(casbin))

	a, err := service.NewService(cfg)
	if err != nil {
		log.Info("Error while creating service", err)
		return nil
	}

	post := handler.NewPostHandler(a, log)
	chat := handler.NewChatHandler(a, log)
	like := handler.NewLikeHandler(a, log)
	comment := handler.NewCommentHandler(a, log)
	user := handler.NewUserHandler(a, log)

	userGroup := router.Group("/user")
	{

	}

	postGroup := router.Group("/post")
	{
		postGroup.POST("/create", post.CreatePost)
		postGroup.PUT("/update", post.UpdatePost)
		postGroup.DELETE("/delete/:id", post.DeletePost)
		postGroup.GET("/getBy/:id", post.GetPostByID)
		postGroup.GET("/list", post.ListPosts)
		postGroup.PUT("/add-image", post.AddImageToPost)
		postGroup.DELETE("/remove-image", post.RemoveImageFromPost)
		postGroup.GET("/country/:country", post.GetPostByCountry)
	}

	chatGroup := router.Group("/chat")
	{

	}

	likeGroup := router.Group("/like")
	{

	}

	commentGroup := router.Group("/comment")
	{

	}
}
