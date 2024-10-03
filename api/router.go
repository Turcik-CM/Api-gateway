package api

import (
	_ "api-gateway/api/docs"
	"api-gateway/api/handler"
	"api-gateway/api/middleware"
	"api-gateway/pkg/config"
	"api-gateway/service"
	"api-gateway/service/redis"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

// @title API Gateway Service for Turk-SM
// @version 1.0
// @description API for API Gateway Service
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token here
func NewRouter(cfg *config.Config, log *slog.Logger, casbin *casbin.Enforcer, redis *redis.RedisStorage) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	auth := handler.NewAuthHandler(log, a, redis)
	country := handler.NewCountriesHandlers(a, log)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login/email", auth.LoginEmail)
		authGroup.POST("/login/username", auth.LoginUsername)
		authGroup.POST("/accept-code", auth.AcceptCodeToRegister)
		authGroup.POST("/forgot-password", auth.ForgotPassword)
		authGroup.POST("/register-admin", auth.RegisterAdmin)
		authGroup.POST("/reset-password", auth.ResetPassword)
	}

	router1 := router.Group("")
	router1.Use(middleware.PermissionMiddleware(casbin))

	attraction := router1.Group("attraction")
	{
		attraction.POST("/create", att.CreateAttraction)
		attraction.PUT("/update", att.UpdateAttraction)
		attraction.GET("/getBy/:id", att.GetAttractionByID)
		attraction.DELETE("/delete/:id", att.DeleteAttraction)
		attraction.GET("/list", att.ListAttractions)
		attraction.GET("/list_search", att.SearchAttractions)
		attraction.PUT("/image/:id", att.UpdateImage)
		attraction.DELETE("/remove-image/:id", att.RemoveHistoricalImage)

	}

	attractionType := router1.Group("attraction-type")
	{
		attractionType.POST("/create", att.CreateAttractionType)
		attractionType.PUT("/update", att.UpdateAttractionType)
		attractionType.GET("/get/:id", att.GetAttractionByIDType)
		attractionType.DELETE("/delete/:id", att.DeleteAttractionType)
		attractionType.GET("/list", att.ListAttractionsType)
	}

	nationalCountry := router1.Group("/country")
	{
		nationalCountry.POST("/create", country.CreateCountry)
		nationalCountry.PUT("/update", country.UpdateCountry)
		nationalCountry.GET("/get/:id", country.GetCountryByID)
		nationalCountry.DELETE("/delete/:id", country.DeleteCountry)
		nationalCountry.GET("/list", country.ListCountries)
	}

	nationalFood := router1.Group("/national")
	{
		nationalFood.POST("/create", nat.CreateNationalFood)
		nationalFood.PUT("/update", nat.UpdateNationalFood)
		nationalFood.GET("/getBy/:id", nat.GetNationalFoodByID)
		nationalFood.DELETE("/delete/:id", nat.DeleteNationalFood)
		nationalFood.GET("/list", nat.ListNationalFoods)
		nationalFood.PUT("/image/:id", nat.UpdateImage)
	}

	history := router1.Group("historical")
	{
		history.POST("/create", his.AddHistorical)
		history.PUT("/update", his.UpdateHistoricals)
		history.GET("/getBy/:id", his.GetHistoricalByID)
		history.DELETE("/delete/:id", his.DeleteHistorical)
		history.GET("/list", his.ListHistorical)
		history.GET("/list_search", his.SearchHistorical)
		history.PUT("/image/:id", his.UpdateHisImage)
	}

	admin := router1.Group("admin")
	{
		admin.GET("/fetch_users", user.FetchUsers)
		admin.POST("/create-user", user.Create)
		admin.DELETE("/delete-user/:id", user.DeleteUser)
		admin.GET("/user-by-id/:id", user.GetProfileById)
	}

	userGroup := router1.Group("/user")
	{
		userGroup.GET("/get-profile", user.GetProfile)
		userGroup.PUT("/update-profile", user.UpdateProfile)
		userGroup.PUT("/change-password", user.ChangePassword)
		userGroup.PUT("/change-profile-image", user.ChangeProfileImage)
		userGroup.GET("/list-of-following", user.ListOfFollowing)
		userGroup.GET("/list-of-followers", user.ListOfFollowers)
		userGroup.DELETE("/delete", user.DeleteProfile)
		userGroup.POST("/follow", user.Follow)
		userGroup.DELETE("/unfollow/:id", user.Unfollow)
		userGroup.GET("/most-popular-user", user.MostPopularUser)
	}

	postGroup := router1.Group("/post")
	{
		postGroup.POST("/create", post.CreatePost)
		postGroup.PUT("/update", post.UpdatePost)
		postGroup.DELETE("/delete/:id", post.DeletePost)
		postGroup.GET("/getBy/:id", post.GetPostByID)
		postGroup.GET("/list", post.ListPosts)
		postGroup.POST("/image/:id", post.UpdatePost)
		postGroup.DELETE("/remove-image/:id", post.RemoveImageFromPost)
		postGroup.GET("/country/:country", post.GetPostByCountry)
	}

	chatGroup := router1.Group("/chat")
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

	likeGroup := router1.Group("/like")
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

	commentGroup := router1.Group("/comment")
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
