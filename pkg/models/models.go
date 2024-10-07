package models

type Void struct {
}

type Search struct {
	Action string `json:"string"`
}

// ChatResponse represents chat response structure
type ChatResponse struct {
	ID        string `json:"id" db:"id" default:"0"`
	User2ID   string `json:"user2_id" db:"user2_id" default:"0"`
	CreatedAt string `json:"created_at" db:"created_at" default:"2023-01-01T00:00:00Z"`
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

type List struct {
	Limit  int64  `json:"limit" default:"10"`
	Page   int64  `json:"page" default:"0"`
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
	Total    string            `json:"total"`
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
	Post  []PostResponse `json:"post"`
	Total string         `json:"total"`
}

// Message represents a simple message response
type Message struct {
	Massage string `json:"massage"`
}

// PostCountry represents a country related to a post
type PostCountry struct {
	Country string `form:"country_name" json:"country" default:"Uzbekistan"`
}

// LikeList represents pagination for liked posts
type LikeList struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

// PostList represents pagination and filtering options for a list of posts
type PostList struct {
	Limit   int64  `json:"limit" default:"10"`
	Page    int64  `json:"page" default:"0"`
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

type LoginResponse struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

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

type UpdateProfileRequest struct {
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Username    string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
	Bio         string `json:"bio" db:"bio"`
	Phone       string `json:"phone" db:"phone"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" db:"current_password"`
	NewPassword     string `json:"new_password" db:"new_password"`
}

type ChangePasswordResponse struct {
	Message string `json:"message" db:"message"`
}

type URL struct {
	URL string `json:"url" db:"url"`
}

type Error struct {
	Error string `json:"error"`
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
	FoodName    string `form:"food_name"`
	FoodType    string `form:"food_type"`
	Description string `form:"description"`
	CountryId   string `form:"country_id"`
	ImageURL    string `form:"image_url"`
	Ingredients string `form:"ingredients"`
}

type NationalFoodResponse struct {
	ID          string `json:"id"`
	FoodName    string `json:"country"`
	FoodType    string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	CountryId   int32  `json:"rating"`
	Ingredients string `json:"ingredients"`
	CreatedAt   string `json:"created_at"`
}

type NationalFoodList struct {
	Limit     int64  `json:"limit" default:"10"`
	Page      int64  `json:"page" default:"0"`
	CountryId string `json:"country_id" default:"Uzbekistan"`
}

type NationalFoodListResponse struct {
	NationalFood []NationalFoodResponse `json:"national_food"`
	Total        string                 `json:"total"`
}

type Attraction struct {
	City        string `form:"city" default:"Uzbekistan"`
	Name        string `form:"name" default:"0"`
	Description string `form:"description" default:"0"`
	Category    string `form:"category" default:"culture"`
	ImageURL    string `form:"image_url" default:"0"`
	Location    string `form:"location" default:"0"`
}

type AttractionList struct {
	City        string `json:"city"`
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
	Total       string               `json:"total"`
}

type UpdateNationalFood struct {
	ID          string `json:"id"`
	FoodName    string `json:"food_name"`
	FoodType    string `json:"food_type"`
	Description string `json:"description"`
	CountryId   string `json:"country_id"`
	ImageURL    string `json:"image_url"`
	Ingredients string `json:"ingredients"`
}

type AttractionSearch struct {
	SearchTerm string `json:"search_term" default:"dodi"`
	Limit      string `json:"limit" default:"10"`
	Offset     string `json:"offset" default:"0"`
}

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

type UpdateCountry struct {
	Id   string `json:"id" default:"0" form:"id"`
	Name string `json:"name" default:"0" form:"name"`
}

type GetCountryResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type UpdateCountryResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ListCountriesRequest struct {
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
	Name  string `json:"name"`
}

type ListCountriesResponse struct {
	Countries []Country `json:"countries"`
	Total     string    `json:"total"`
}

type Country struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CreateCityRequest struct {
	CountryID string `json:"country_id"`
	Name      string `json:"name"`
}

type CreateCityResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateAttractionTypeRequest struct {
	Name     string `json:"name"`
	Activity int32  `json:"activity"`
}

type CreateAttractionTypeResponse struct {
	AttractionType AttractionType `json:"attraction_type"`
}

type GetAttractionTypeResponse struct {
	AttractionType AttractionType `json:"attraction_type"`
}

type UpdateAttractionTypeRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Activity int32  `json:"activity"`
}

type UpdateAttractionTypeResponse struct {
	AttractionType AttractionType `json:"attraction_type"`
}

type ListAttractionTypesRequest struct {
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
	Name  string `json:"name"`
}

type ListAttractionTypesResponse struct {
	AttractionTypes []AttractionType `json:"attraction_types"`
	Total           string           `json:"total"`
}

type AttractionType struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Activity int32  `json:"activity"`
}

type FilterCountry struct {
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
	Name  string `json:"name"`
}
