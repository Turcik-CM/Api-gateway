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

	userGroup := router.Group("/user")
	{
		// Create a new user
		userGroup.POST("/create", user.Create)

		// Get user profile
		userGroup.GET("/get_profile", user.GetProfile)

		// Update user profile
		userGroup.PUT("/update_profile", user.UpdateProfile)

		// Change user password
		userGroup.PUT("/change_password", user.ChangePassword)

		// Change user profile image
		userGroup.PUT("/change_profile_image", user.ChangeProfileImage)

		// Fetch users with filtering options
		userGroup.GET("/fetch_users", user.FetchUsers)

		// List of users that the current user is following
		userGroup.GET("/list_of_following", user.ListOfFollowing)

		// List of followers for the current user
		userGroup.GET("/list_of_followers", user.ListOfFollowers)

		// Delete a user account (Admin)
		userGroup.DELETE("/delete/:user_id", user.DeleteUser)

		// Get user profile by user ID (Admin)
		userGroup.GET("/get_profile_by_id/:user_id", user.GetProfileById)

		// Update user profile by user ID (Admin)
		userGroup.PUT("/update_profile_by_id/:user_id", user.UpdateProfileById)

		// Change user profile image by user ID (Admin)
		userGroup.PUT("/change_profile_image_by_id/:user_id", user.ChangeProfileImageById)

		// Follow a user
		userGroup.POST("/follow/:user_id", user.Follow)

		// Unfollow a user
		userGroup.DELETE("/unfollow/:user_id", user.Unfollow)

		// Get followers of a user
		userGroup.GET("/followers/:user_id", user.GetUserFollowers)

		// Get users followed by a user
		userGroup.GET("/follows/:user_id", user.GetUserFollows)

		// Get the most popular user
		userGroup.GET("/most_popular", user.MostPopularUser)
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
		chatGroup.POST("/create", chat.StartMessaging)              // @Router /chat/create [post]
		chatGroup.POST("/create_message", chat.SendMessage)         // @Router /chat/create_message [post]
		chatGroup.POST("/message_true", chat.MessageMarcTrue)       // @Router /chat/message_true [post]
		chatGroup.GET("/get_user", chat.GetUserChats)               // @Router /chat/get_user [get]
		chatGroup.GET("/get_massage", chat.GetUnreadMessages)       // @Router /chat/get_massage [get]
		chatGroup.PUT("/update", chat.UpdateMessage)                // @Router /chat/update [put]
		chatGroup.GET("/get_today", chat.GetTodayMessages)          // @Router /chat/get_today [get]
		chatGroup.DELETE("/delete_massage/:id", chat.DeleteMessage) // @Router /chat/delete_massage{id} [delete]
		chatGroup.DELETE("/delete_chat/:id", chat.DeleteChat)       // @Router /chat/delete_chat{id} [delete]
		chatGroup.GET("/list", chat.GetChatMessages)                // @Router /chat/list [get]
	}

	likeGroup := router.Group("/like")
	{
		// Add Like to a Post
		// @Router /like/create [post]
		likeGroup.POST("/create", like.AddLikePost)

		// Delete Like from a Post
		// @Router /like/delete/:post_id [delete]
		likeGroup.DELETE("/delete/:post_id", like.DeleteLikePost)

		// Add Like to a Comment
		// @Router /like/comment/create [post]
		likeGroup.POST("/comment/create", like.AddLikeComment)

		// Delete Like from a Comment
		// @Router /like/comment/delete/:commit_id [delete]
		likeGroup.DELETE("/comment/delete/:commit_id", like.DeleteLikeComment)

		// Get Post Like Count
		// @Router /like/post/count/:post_id [get]
		likeGroup.GET("/post/count/:post_id", like.GetPostLikeCount)

		// Get Most Liked Comment
		// @Router /like/comment/most-liked/:post_id [get]
		likeGroup.GET("/comment/most-liked/:post_id", like.GetMostLikedComment)

		// Get Users Who Liked a Post
		// @Router /like/post/users/:post_id [get]
		likeGroup.GET("/post/users/:post_id", like.GetUsersWhichLikePost)

		// Get Users Who Liked a Comment
		// @Router /like/comment/users/:comment_id [get]
		likeGroup.GET("/comment/users/:comment_id", like.GetUsersWhichLikeComment)
	}

	commentGroup := router.Group("/comment")
	{
		commentGroup.POST("/create", comment.CreateComment)                     // @Router /comment/create [post]
		commentGroup.PUT("/update", comment.UpdateComment)                      // @Router /comment/update [put]
		commentGroup.GET("/getBy/:id", comment.GetCommentByID)                  // @Router /comment/getBy{id} [get]
		commentGroup.GET("/getByUser/:id", comment.GetCommentByUsername)        // @Router /comment/getByUser{id} [get]
		commentGroup.GET("/list", comment.ListComments)                         // @Router /comment/list [get]
		commentGroup.DELETE("/delete/:id", comment.DeleteComment)               // @Router /comment/delete{id} [delete]
		commentGroup.GET("/get_comment_postId/:id", comment.GetCommentByPostID) // @Router /comment/get_comment_postId/{id} [get]
		commentGroup.GET("/user_comments", comment.GetAllUserComments)          // @Router /comment/user_comments [get]
		commentGroup.GET("/most_like_post/:id", comment.GetMostlikeCommentPost) // @Router /comment/most_like_post/{id} [get]
	}
	return router
}
