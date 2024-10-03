package models

type Void struct {
}

// Search represents a search action
type Search struct {
	Action string `json:"action" default:"search"`
}

// UpdateMs represents a message update request
type UpdateMs struct {
	MessageID string `json:"message_id" db:"message_id" default:"123456"`
	Text      string `json:"text" default:"Hello"`
}

// ChatResponse represents chat response structure
type ChatResponse struct {
	ID        string `json:"id" db:"id" default:"0"`
	User2ID   string `json:"user2_id" db:"user2_id" default:"0"`
	CreatedAt string `json:"created_at" db:"created_at" default:"2023-01-01T00:00:00Z"`
}

// ChatResponseList represents a list of chats
type ChatResponseList struct {
	Chat []ChatResponse `json:"chat"`
}

// MassageResponse represents a message response structure
type MassageResponse struct {
	ID          string `json:"id" db:"id" default:"0"`
	ChatID      string `json:"chat_id" db:"chat_id" default:"0"`
	SenderID    string `json:"sender_id" db:"sender_id" default:"0"`
	ContentType string `json:"content_type" db:"content_type" default:"text"`
	Content     string `json:"content" db:"content" default:"Hello"`
	IsRead      bool   `json:"is_read" db:"is_read" default:"false"`
	CreatedAt   string `json:"created_at" db:"created_at" default:"2023-01-01T00:00:00Z"`
	UpdatedAt   string `json:"updated_at" db:"updated_at" default:"2023-01-01T00:00:00Z"`
}

// MassageResponseList represents a list of messages
type MassageResponseList struct {
	Massage []MassageResponse `json:"massage"`
}

// List represents a pagination request for messages or posts
type List struct {
	Limit  int64  `json:"limit" default:"10"`
	Offset int64  `json:"offset" default:"0"`
	ChatID string `json:"chat_id" default:"0"`
}

// MassageId represents a structure for message identification
type MassageId struct {
	MassageID string `json:"massage_id" default:"0"`
}

// ChatId represents a structure for chat identification
type ChatId struct {
	ChatID string `json:"chat_id" default:"0"`
}

// MassageTrue represents a request to mark a message as read
type MassageTrue struct {
	ChatID string `json:"chat_id" default:"0"`
}

// CreateMassage represents the structure for creating a message
type CreateMassage struct {
	ChatID      string `json:"chat_id" default:"0"`
	SenderID    string `json:"sender_id" default:"0"`
	ContentType string `json:"content_type" default:"text"`
	Content     string `json:"content" default:"Hello"`
}

// CreateChat represents the structure for creating a chat
type CreateChat struct {
	User2ID string `json:"user2_id" default:"0"`
}

// CommentResponse represents a single comment structure
type CommentResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	PostID    string `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CommentsR represents a list of comments
type CommentsR struct {
	Comments []CommentResponse `json:"comments"`
}

// Username represents a username structure
type Username struct {
	Username string `json:"username" default:"0"`
}

// Users represents a list of usernames
type Users struct {
	Users []string `json:"users"`
}

// ImageUrl represents a post image URL structure
type ImageUrl struct {
	PostID string `form:"post_id" default:"0"`
	URL    string `json:"url" default:"https://example.com/image.jpg"`
}

// UserPostId represents the user and post identification
type UserPostId struct {
	UserID string `json:"user_id" default:"0"`
	PostID string `json:"post_id" default:"0"`
}

// CommentAllResponse represents a list of comments
type CommentAllResponse struct {
	Comments []CommentResponse `json:"comments" `
}

// CommentList represents pagination for comments related to a post
type CommentList struct {
	Limit  int64  `json:"limit" default:"10"`
	Offset int64  `json:"offset" default:"0"`
	PostID string `json:"post_id" default:"0"`
}

// CommentPost represents the structure for posting a comment
type CommentPost struct {
	PostID  string `json:"post_id" default:"0"`
	Content string `json:"content"  default:"dodi"`
}

// UpdateAComment represents the structure for updating a comment
type UpdateAComment struct {
	ID      string `json:"id" default:"0"`
	Content string `json:"content" default:"dodi"`
}

// LikePost represents a structure for liking a post
type LikePost struct {
	PostID string `json:"post_id"  default:"0"`
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
	CommentID string `json:"comment_id" default:"0"`
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
	Country string `json:"country" default:"Uzbekistan"`
}

// LikeList represents pagination for liked posts
type LikeList struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

// PostList represents pagination and filtering options for a list of posts
type PostList struct {
	Limit   int64  `json:"limit" default:"10"`
	Offset  int64  `json:"offset" default:"0"`
	Country string `json:"country" default:"Uzbekistan"`
	Hashtag string `json:"hashtag" default:"dodi"`
}

// PostId represents the post identification
type PostId struct {
	ID string `json:"id" default:"0"`
}

// UpdateAPost represents the structure for updating a post
type UpdateAPost struct {
	ID       string `json:"id" default:"dodi"`
	Country  string `json:"country" default:"Uzbekistan"`
	Location string `json:"location" default:"dodi"`
	Title    string `json:"title" default:"dodi"`
	Content  string `json:"content" default:"dodi"`
	Hashtag  string `json:"hashtag" default:"dodi"`
	ImageURL string `json:"image_url" default:"dodi"`
}

// Post represents the structure for creating a post
type Post struct {
	Title       string `form:"title" default:"dodi"`
	Content     string `form:"content" default:"dodi"`
	Country     string `form:"country" default:"Uzbekistan"`
	Description string `form:"description" default:"dodi"`
	Hashtag     string `form:"hashtag" default:"dodi"`
	Location    string `form:"location" default:"dodi"`
	ImageUrl    string `form:"image_url" default:"dodi"`
	UserId      string `form:"user_id" default:"dodi"`
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
	Role        string `json:"role" db:"role"`
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
	PostsCount     int32  `json:"posts_count" db: db:"follower_id""posts_count"`
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
	URL string `json:"url" db:"url"`
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

//nationality -----------------------------------------------------------------

type HistoricalImage struct {
	ID  string `json:"id" default:"0"`
	URL string `json:"url" default:"0"`
}

type HistoricalCountry struct {
	City string `json:"city"`
}

type HistoricalSearch struct {
	Search string `json:"search" default:"dodi"`
}

type HistoricalListResponse struct {
	Historical []HistoricalResponse `json:"historical"`
}

type HistoricalList struct {
	Limit   int64  `json:"limit" default:"10"`
	Offset  int64  `json:"offset" default:"0"`
	Country string `json:"country" default:"Uzbekistan"`
}

type HistoricalId struct {
	ID string `json:"id" default:"0"`
}

type UpdateHistorical struct {
	ID          string `json:"id" default:"0"`
	Country     string `json:"country" default:"Uzbekistan"`
	City        string `json:"city" default:"Uzbekistan"`
	Name        string `json:"name" default:"dodi"`
	Description string `json:"description" default:"dodi"`
	ImageURL    string `json:"image_url" default:"0"`
}

type Historical struct {
	Country     string `form:"country" default:"Uzbekistan"`
	City        string `form:"city" default:"Uzbekistan"`
	Name        string `form:"name" default:"dodi"`
	Description string `form:"description" default:"dodi"`
	ImageURL    string `form:"image_url" default:"0"`
}

type HistoricalResponse struct {
	ID          string `json:"id"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type NationalFood struct {
	Country     string `form:"country" default:"Uzbekistan"`
	Name        string `form:"name" default:"dodi"`
	Description string `form:"description" default:"dodi"`
	Nationality string `form:"nationality" default:"dodi"`
	ImageURL    string `form:"image_url" default:"dodi"`
	Rating      int32  `form:"rating" default:"1200000"`
	FoodType    string `form:"food_type" default:"dodi"`
	Ingredients string `form:"ingredients" default:"dodi"`
}

type NationalFoodResponse struct {
	ID          string `json:"id"`
	Country     string `json:"country"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Rating      int32  `json:"rating"`
	Ingredients string `json:"ingredients"`
	FoodType    string `json:"food_type"`
	Nationality string `json:"nationality"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type NationalFoodId struct {
	ID string `json:"id" default:"0"`
}

type NationalFoodList struct {
	Limit   int64  `json:"limit" default:"10"`
	Offset  int64  `json:"offset" default:"0"`
	Country string `json:"country" default:"Uzbekistan"`
}

type NationalFoodListResponse struct {
	NationalFood []NationalFoodResponse `json:"national_food"`
}

type NationalFoodImage struct {
	ID       string `json:"id" default:"0"`
	ImageURL string `json:"image_url" default:"0"`
}

type NationalFoodCountry struct {
	Country string `json:"country"`
}

type NationalFoodSearch struct {
	Search string `json:"search"`
}

type RatingResponse struct {
	Rating int32 `json:"rating"`
}

type Attraction struct {
	Country     string `form:"country" default:"Uzbekistan"`
	Name        string `form:"name" default:"0"`
	Description string `form:"description" default:"0"`
	Category    string `form:"category" default:"culture"`
	ImageURL    string `form:"image_url" default:"0"`
	Location    string `form:"location" default:"0"`
}

type AttractionList struct {
	Country     string `json:"country" default:"Uzbekistan"`
	Category    string `json:"category" default:"culture"`
	Name        string `json:"name" default:"dodi"`
	Description string `json:"description" default:"dodi"`
	Limit       int64  `json:"limit" default:"0"`
	Offset      int64  `json:"offset" default:"0"`
}

type AttractionResponse struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
	Location    string `json:"location"`
	ImageURL    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AttractionListResponse struct {
	Attractions []AttractionResponse `json:"attractions"`
}

type AttractionId struct {
	ID string `json:"id" default:"0"`
}

type UpdateNationalFood struct {
	ID          string `json:"id" default:"0"`
	Country     string `json:"country" default:"dodi"`
	Name        string `json:"name" default:"dodi"`
	Description string `json:"description" default:"dodi"`
	ImageURL    string `json:"image_url" default:"0"`
	Rating      int32  `json:"rating" default:"120000"`
	Nationality string `json:"nationality" default:"dodi"`
	FoodType    string `json:"food_type" default:"dodi"`
	Ingredients string `json:"ingredients" default:"dodi"`
}

type UpdateAttraction struct {
	ID          string `json:"id" default:"0"`
	Country     string `json:"country" default:"Uzbekistan"`
	Name        string `json:"name" default:"0"`
	Description string `json:"description" default:"0"`
	Category    string `json:"category" default:"culture"`
	Location    string `json:"location" default:"dodi"`
	ImageURL    string `json:"image_url" default:"0"`
}

type AttractionImage struct {
	ID       string `json:"id" default:"0"`
	ImageURL string `json:"image_url" default:"0"`
}

type AttractionCountry struct {
	Country string `json:"country"`
}

type AttractionSearch struct {
	SearchTerm string `json:"search_term" default:"dodi"`
	Limit      string `json:"limit" default:"10"`
	Offset     string `json:"offset" default:"0"`
}

// auth -----------------------------------
type RegisterRequest struct {
	Email     string `json:"email" db:"email" default:"your email"`
	Phone     string `json:"phone" db:"phone" default:"+123456789123456"`
	FirstName string `json:"first_name" db:"first_name" default:"Tom"`
	LastName  string `json:"last_name" db:"last_name" default:"Joe"`
	Username  string `json:"username" db:"username" default:"tom0011"`
	Country   string `json:"country" db:"country" default:"Uzbekistan"`
	Password  string `json:"password" db:"password" default:"123456"`
	Bio       string `json:"bio" db:"bio" default:"holasela berish shartmas"`
}
type RegisterRequest1 struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Phone     string `json:"phone" db:"phone"`
	Username  string `json:"username" db:"username"`
	Country   string `json:"country" db:"country"`
	Bio       string `json:"bio" db:"bio"`
	Code      string `json:"code" binding:"required"`
}

type RegisterResponse struct {
	Id           string `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	Flag         string `json:"flag" db:"flag"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type LoginEmailRequest struct {
	Email    string `json:"email" db:"email" default:"registerdagi email ni kiritng"`
	Password string `json:"password" db:"password" default:"123456"`
}

type LoginResponse1 struct {
	Id       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
	Country  string `json:"country" db:"country"`
}

type LoginUsernameRequest struct {
	Username string `json:"username" db:"username" default:"tom0011"`
	Password string `json:"password" db:"password" default:"123456"`
}

type Tokens struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type AcceptCode struct {
	Email string `json:"email" default:"code cogan email ni kiriting"`
	Code  string `json:"code" default:"12369"`
}
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPassReq struct {
	Email    string `json:"email" default:"example@gmail.com"`
	Password string `json:"new_password" default:"123369"`
	Code     string `json:"code" default:"123456"`
}
type UpdatePasswordReq struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
