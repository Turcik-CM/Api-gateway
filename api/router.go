package api

import (
	_ "api-gateway/api/docs"
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

// @title Api-Geteway service for Turk-SM
// @version 1.0
// @description API for Api-Geteway Service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
// @BasePath
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
	att := handler.NewAttractionsHandler(a, log)
	nat := handler.NewNationalFoodHandler(a, log)
	his := handler.NewHistoryHandler(a, log)

	attraction := router.Group("attraction")
	{
		attraction.POST("/create", att.CreateAttraction)
		attraction.PUT("/update", att.UpdateAttraction)
		attraction.GET("/getBy/:id", att.GetAttractionByID)
		attraction.DELETE("/delete/:id", att.DeleteAttraction)
		attraction.GET("/list", att.ListAttractions)
		attraction.GET("/list_search", att.SearchAttractions)
		attraction.POST("/add-image", att.AddImageUrl)
		attraction.DELETE("/remove-image/:id", att.RemoveHistoricalImage)

	}

	nationalFood := router.Group("national")
	{
		nationalFood.POST("/create", nat.CreateNationalFood)
		nationalFood.PUT("/update", nat.UpdateNationalFood)
		nationalFood.GET("/getBy/:id", nat.GetNationalFoodByID)
		nationalFood.DELETE("/delete/:id", nat.DeleteNationalFood)
		nationalFood.GET("/list", nat.ListNationalFoods)
		nationalFood.POST("/add-image", nat.AddImageUrll)
	}

	history := router.Group("historical")
	{
		history.POST("/create", his.AddHistorical)
		history.PUT("/update", his.UpdateHistoricals)
		history.GET("/getBy/:id", his.GetHistoricalByID)
		history.DELETE("/delete/:id", his.DeleteHistorical)
		history.GET("/list", his.ListHistorical)
		history.GET("/list_search", his.SearchHistorical)
		history.POST("/add-image", his.AddHistoricalImage)
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", user.Create)
		userGroup.GET("/get_profile", user.GetProfile)
		userGroup.PUT("/update_profile", user.UpdateProfile)
		userGroup.PUT("/change_password", user.ChangePassword)
		userGroup.PUT("/change_profile_image", user.ChangeProfileImage)
		userGroup.GET("/fetch_users", user.FetchUsers)
		userGroup.GET("/list_of_following", user.ListOfFollowing)
		userGroup.GET("/list_of_followers", user.ListOfFollowers)
		userGroup.DELETE("/delete/:user_id", user.DeleteUser)
		userGroup.GET("/get_profile_by_id/:user_id", user.GetProfileById)
		userGroup.PUT("/update_profile_by_id/:user_id", user.UpdateProfileById)
		userGroup.PUT("/change_profile_image_by_id/:user_id", user.ChangeProfileImageById)
		userGroup.POST("/follow/:user_id", user.Follow)
		userGroup.DELETE("/unfollow/:user_id", user.Unfollow)
		userGroup.GET("/followers/:user_id", user.GetUserFollowers)
		userGroup.GET("/follows/:user_id", user.GetUserFollows)
		userGroup.GET("/most_popular", user.MostPopularUser)
	}

	postGroup := router.Group("/post")
	{
		postGroup.POST("/create", post.CreatePost)
		postGroup.PUT("/update", post.UpdatePost)
		postGroup.DELETE("/delete/:id", post.DeletePost)
		postGroup.GET("/getBy/:id", post.GetPostByID)
		postGroup.GET("/list", post.ListPosts)
		postGroup.POST("/add-image", post.AddImageToPost)
		postGroup.DELETE("/remove-image/:id", post.RemoveImageFromPost)
		postGroup.GET("/country/:c", post.GetPostByCountry)
	}

	chatGroup := router.Group("/chat")
	{
		chatGroup.POST("/create", chat.StartMessaging)
		chatGroup.POST("/create_message", chat.SendMessage)
		chatGroup.POST("/message_true", chat.MessageMarcTrue)
		chatGroup.GET("/get_user", chat.GetUserChats)
		chatGroup.GET("/get_massage/:id", chat.GetUnreadMessages)
		chatGroup.PUT("/update", chat.UpdateMessage)
		chatGroup.GET("/get_today/:id", chat.GetTodayMessages)
		chatGroup.DELETE("/delete_massage/:id", chat.DeleteMessage)
		chatGroup.DELETE("/delete_chat/:id", chat.DeleteChat)
		chatGroup.GET("/list", chat.GetChatMessages)
	}

	likeGroup := router.Group("/like")
	{
		likeGroup.POST("/create", like.AddLikePost)
		likeGroup.DELETE("/delete/:post_id", like.DeleteLikePost)
		likeGroup.POST("/comment/create", like.AddLikeComment)
		likeGroup.DELETE("/comment/delete/:commit_id", like.DeleteLikeComment)
		likeGroup.GET("/post/count/:post_id", like.GetPostLikeCount)
		likeGroup.GET("/comment/most-liked/:post_id", like.GetMostLikedComment)
		likeGroup.GET("/post/users/:post_id", like.GetUsersWhichLikePost)
		likeGroup.GET("/comment/users/:comment_id", like.GetUsersWhichLikeComment)
	}

	commentGroup := router.Group("/comment")
	{
		commentGroup.POST("/create", comment.CreateComment)
		commentGroup.PUT("/update", comment.UpdateComment)
		commentGroup.GET("/getBy/:id", comment.GetCommentByID)
		commentGroup.GET("/getByUser/:id", comment.GetCommentByUsername)
		commentGroup.GET("/list", comment.ListComments)
		commentGroup.DELETE("/delete/:id", comment.DeleteComment)
		commentGroup.GET("/get_comment_postId/:id", comment.GetCommentByPostID)
		commentGroup.GET("/country", comment.GetAllUserComments)
		commentGroup.GET("/most_like_post/:id", comment.GetMostlikeCommentPost)
	}
	return router
}
