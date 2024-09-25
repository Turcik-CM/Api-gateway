package models

type Void struct {
}

// Search represents a search action
type Search struct {
	Action string `json:"action"`
}

// UpdateMs represents a message update request
type UpdateMs struct {
	MessageID string `json:"message_id" db:"message_id"`
	Text      string `json:"text"`
}

// ChatResponse represents chat response structure
type ChatResponse struct {
	ID        string `json:"id" db:"id"`
	User2ID   string `json:"user2_id" db:"user2_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

// ChatResponseList represents a list of chats
type ChatResponseList struct {
	Chat []ChatResponse `json:"chat"`
}

// MassageResponse represents a message response structure
type MassageResponse struct {
	ID          string `json:"id" db:"id"`
	ChatID      string `json:"chat_id" db:"chat_id"`
	SenderID    string `json:"sender_id" db:"sender_id"`
	ContentType string `json:"content_type" db:"content_type"`
	Content     string `json:"content" db:"content"`
	IsRead      bool   `json:"is_read" db:"is_read"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

// MassageResponseList represents a list of messages
type MassageResponseList struct {
	Massage []MassageResponse `json:"massage"`
}

// List represents a pagination request for messages or posts
type List struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	ChatID string `json:"chat_id"`
}

// MassageId represents a structure for message identification
type MassageId struct {
	MassageID string `json:"massage_id"`
}

// ChatId represents a structure for chat identification
type ChatId struct {
	ChatID string `json:"chat_id"`
}

// MassageTrue represents a request to mark a message as read
type MassageTrue struct {
	ChatID string `json:"chat_id"`
}

// CreateMassage represents the structure for creating a message
type CreateMassage struct {
	ChatID      string `json:"chat_id"`
	SenderID    string `json:"sender_id"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
}

// CreateChat represents the structure for creating a chat
type CreateChat struct {
	User2ID string `json:"user2_id"`
}

// CommentResponse represents a single comment structure
type CommentResponse struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	PostID    string `json:"post_id" db:"post_id"`
	Content   string `json:"content" db:"content"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

// CommentsR represents a list of comments
type CommentsR struct {
	Comments []CommentResponse `json:"comments"`
}

// Username represents a username structure
type Username struct {
	Username string `json:"username"`
}

// Users represents a list of usernames
type Users struct {
	Users []string `json:"users"`
}

// ImageUrl represents a post image URL structure
type ImageUrl struct {
	PostID string `json:"post_id"`
	URL    string `json:"url"`
}

// UserPostId represents the user and post identification
type UserPostId struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

// CommentAllResponse represents a list of comments
type CommentAllResponse struct {
	Comments []CommentResponse `json:"comments"`
}

// CommentList represents pagination for comments related to a post
type CommentList struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	PostID string `json:"post_id"`
}

// CommentPost represents the structure for posting a comment
type CommentPost struct {
	PostID    string `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// UpdateAComment represents the structure for updating a comment
type UpdateAComment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// LikePost represents a structure for liking a post
type LikePost struct {
	PostID string `json:"post_id"`
}

// LikeCommit represents a structure for liking a comment
type LikeCommit struct {
	CommitID string `json:"commit_id"`
}

// LikeResponse represents a response for a post like
type LikeResponse struct {
	PostID    string `json:"post_id"`
	CreatedAt string `json:"created_at"`
}

// LikeComResponse represents a response for a comment like
type LikeComResponse struct {
	CommitID  string `json:"commit_id"`
	CreatedAt string `json:"created_at"`
}

// CommentId represents a structure for comment identification
type CommentId struct {
	CommentID string `json:"comment_id"`
}

// PostListResponse represents a list of posts
type PostListResponse struct {
	Post []PostResponse `json:"post"`
}

// Message represents a simple message response
type Message struct {
	Massage string `json:"massage"`
}

// PostCountry represents a country related to a post
type PostCountry struct {
	Country string `json:"country"`
}

// LikeList represents pagination for liked posts
type LikeList struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

// PostList represents pagination and filtering options for a list of posts
type PostList struct {
	Limit   int64  `json:"limit"`
	Offset  int64  `json:"offset"`
	Country string `json:"country"`
	Hashtag string `json:"hashtag"`
}

// PostId represents the post identification
type PostId struct {
	ID string `json:"id"`
}

// UpdateAPost represents the structure for updating a post
type UpdateAPost struct {
	ID       string `json:"id"`
	Country  string `json:"country"`
	Location string `json:"location"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Hashtag  string `json:"hashtag"`
	ImageURL string `json:"image_url"`
}

// Post represents the structure for creating a post
type Post struct {
	Country     string `json:"country"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Hashtag     string `json:"hashtag"`
	Content     string `json:"content"`
}

// LikeCount represents the count of likes for a post
type LikeCount struct {
	ID    string `json:"id"`
	Count string `json:"count"`
}

// PostResponse represents the structure of a single post response
type PostResponse struct {
	ID          string `json:"id"`
	Country     string `json:"country"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Hashtag     string `json:"hashtag"`
	Content     string `json:"content"`
	ImageURL    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

//-----------------user---------------------------

// DFollowRes represents the data of a follower and following relationship, including when it was unfollowed.
type DFollowRes struct {
	FollowingID  string `json:"following_id" db:"following_id"`
	UnfollowedAt string `json:"unfollowed_at" db:"unfollowed_at"`
}

// Count represents a count of something with a description.
type Count struct {
	Description string `json:"description" db:"description"`
	Count       int64  `json:"count" db:"count"`
}

// FollowReq represents a follow request between two users.
type FollowReq struct {
	FollowingID string `json:"following_id" db:"following_id"`
}

// FollowRes represents a follow response with the time it was followed.
type FollowRes struct {
	FollowingID string `json:"following_id" db:"following_id"`
	FollowedAt  string `json:"followed_at" db:"followed_at"`
}

// Id represents a generic user identifier.
type Id struct {
}

// CreateRequest represents a request for user creation with various user attributes.
type CreateRequest struct {
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Phone       string `json:"phone" db:"phone"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Username    string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
	Bio         string `json:"bio" db:"bio"`
}

// UserResponse represents the response of user details.
type UserResponse struct {
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Username    string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
	Bio         string `json:"bio" db:"bio"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

// LoginRequest represents a request for logging in.
type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

// LoginResponse represents the response after a successful login.
type LoginResponse struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

// GetProfileResponse represents the response when retrieving a user profile.
type GetProfileResponse struct {
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Email          string `json:"email" db:"email"`
	PhoneNumber    string `json:"phone_number" db:"phone_number"`
	Username       string `json:"username" db:"username"`
	Nationality    string `json:"nationality" db:"nationality"`
	Bio            string `json:"bio" db:"bio"`
	ProfileImage   string `json:"profile_image" db:"profile_image"`
	FollowersCount int32  `json:"followers_count" db:"followers_count"`
	FollowingCount int32  `json:"following_count" db:"following_count"`
	PostsCount     int32  `json:"posts_count" db:"posts_count"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}

// UpdateProfileRequest represents the request to update a user's profile.
type UpdateProfileRequest struct {
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	Username     string `json:"username" db:"username"`
	Nationality  string `json:"nationality" db:"nationality"`
	Bio          string `json:"bio" db:"bio"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
	Phone        string `json:"phone" db:"phone"`
}

// Filter represents filtering options with pagination for admin-specific requests.
type Filter struct {
	Page      int32  `json:"page" db:"page"`
	Limit     int32  `json:"limit" db:"limit"`
	FirstName string `json:"first_name" db:"first_name"`
}

// UserResponses represents a list of user responses.
type UserResponses struct {
	Users []UserResponse `json:"users" db:"users"`
}

// ChangePasswordRequest represents the request to change a user's password.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" db:"current_password"`
	NewPassword     string `json:"new_password" db:"new_password"`
}

// ChangePasswordResponse represents the response after a successful password change.
type ChangePasswordResponse struct {
	Message string `json:"message" db:"message"`
}

// URL represents a URL and associated user ID.
type URL struct {
	URL    string `json:"url" db:"url"`
	UserID string `json:"user_id" db:"user_id"`
}

// Ids represents the follower and following relationship.
type Ids struct {
	FollowerID  string `json:"follower_id" db:"follower_id"`
	FollowingID string `json:"following_id" db:"following_id"`
}

// FollowUser represents a user being followed, with their username and ID.
type FollowUser struct {
	Username string `json:"username" db:"username"`
	ID       string `json:"id" db:"id"`
}

// Follows represents a list of users being followed.
type Follows struct {
	Following []FollowUser `json:"following" db:"following"`
}

type Error struct {
	Error string `json:"error"`
}
