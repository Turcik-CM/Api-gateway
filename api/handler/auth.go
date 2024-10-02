package handler

import (
	"api-gateway/api/email"
	pb "api-gateway/genproto/user"
	"api-gateway/pkg/config"
	"api-gateway/pkg/hashing"
	"api-gateway/pkg/models"
	"api-gateway/pkg/token"
	"api-gateway/service"
	"api-gateway/service/redis"
	"context"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type AuthHandler interface {
	Register(c *gin.Context)
	LoginEmail(c *gin.Context)
	LoginUsername(c *gin.Context)
	AcceptCodeToRegister(c *gin.Context)
	ForgotPassword(c *gin.Context)
	RegisterAdmin(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type authHandler struct {
	authService pb.UserServiceClient
	log         *slog.Logger
	redis       *redis.RedisStorage
}

func NewAuthHandler(logg *slog.Logger, sr service.Service, redis *redis.RedisStorage) AuthHandler {
	authClient := sr.UserService()
	if authClient == nil {
		log.Fatalf("failed to create auth service")
		return nil
	}
	return &authHandler{
		authService: authClient,
		log:         logg,
		redis:       redis,
	}
}

// Register godoc
// @Summary Register Users
// @Description create users
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register body models.RegisterRequest true "register user"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	var auth models.RegisterRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := checkmail.ValidateFormat(auth.Email)
	if err != nil {
		h.log.Error("Invalid email provided", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email provided"})
		return
	}
	code, err := email.Email(auth.Email)
	if err != nil {
		h.log.Error("Invalid email provided", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email provided: " + err.Error()})
		return
	}
	req1 := models.RegisterRequest1{
		FirstName: auth.FirstName,
		LastName:  auth.LastName,
		Email:     auth.Email,
		Phone:     auth.Phone,
		Username:  auth.Username,
		Country:   auth.Country,
		Bio:       auth.Bio,
		Password:  auth.Password,
	}
	req1.Code = code

	err = h.redis.SetRegister(c, req1)
	if err != nil {
		h.log.Error("Failed to register user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	h.log.Info("Successfully saved to redis")

	c.JSON(http.StatusOK, gin.H{"info": "code sent to this email " + req1.Email})
}

// AcceptCodeToRegister godoc
// @Summary Accept code to register
// @Description it accepts code to register
// @Tags Auth
// @Param token body models.AcceptCode true "enough"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/accept-code [post]
func (h *authHandler) AcceptCodeToRegister(c *gin.Context) {
	var req models.AcceptCode
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid data provided", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	register, err := h.redis.GetRegister(ctx, req.Email)
	if err != nil {
		h.log.Error("Failed to get register from redis", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get register from redis; " + err.Error()})
		return
	}

	if register.Code != req.Code {
		h.log.Error("Invalid code", "code", req.Code)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code"})
		return
	}

	res := pb.RegisterRequest{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Phone:     register.Phone,
		Username:  register.Username,
		Country:   register.Country,
		Bio:       register.Bio,
		Password:  register.Password,
	}

	response, err := h.authService.Register(context.Background(), &res)
	if err != nil {
		h.log.Error("Failed to register student", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register student; " + err.Error()})
		return
	}

	reqToken := models.LoginResponse1{
		Id:       response.Id,
		Email:    response.Email,
		Username: register.Username,
		Role:     "user",
		Country:  register.Country,
	}

	accessToken, err := token.GenerateAccessToken(reqToken)
	if err != nil {
		h.log.Error("Failed to generate access token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token; " + err.Error()})
		return
	}

	refreshToken, err := token.GenerateRefreshToken(reqToken)
	if err != nil {
		h.log.Error("Failed to generate refresh token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token; " + err.Error()})
		return
	}

	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	c.JSON(http.StatusOK, response)
}

// @Summary LoginEmail Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginEmail body models.LoginEmailRequest true "register user"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/login/email [post]
func (h *authHandler) LoginEmail(c *gin.Context) {
	var auth pb.LoginEmailRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.LoginEmail(context.Background(), &auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary LoginUsername Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginUsername body models.LoginUsernameRequest true "register user"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/login/username [post]
func (h *authHandler) LoginUsername(c *gin.Context) {
	var auth pb.LoginUsernameRequest
	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.LoginUsername(context.Background(), &auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", res.AccessToken, 3600, "", "", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 3600, "", "", false, true)

	c.JSON(http.StatusOK, res)
}

// ForgotPassword godoc
// @Summary Forgot Password
// @Description it sends code to your email address
// @Tags Auth
// @Param token body models.ForgotPasswordRequest true "enough"
// @Success 200 {object} string "message"
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/forgot-password [post]
func (h *authHandler) ForgotPassword(c *gin.Context) {
	h.log.Info("ForgotPassword is working")
	var req models.ForgotPasswordRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email1 := pb.Email{
		Email: req.Email,
	}

	_, err := h.authService.GetUserByEmail(context.Background(), &email1)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not registered"})
		return
	}

	code, err := email.Email(req.Email)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email " + err.Error()})
		return
	}
	err = h.redis.SetCode(c, req.Email, code)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing codes in Redis " + err.Error()})
		return
	}
	h.log.Info("ForgotPassword succeeded")
	c.JSON(200, gin.H{"message": "Password reset code sent to your email"})
}

// RegisterAdmin godoc
// @Summary Registers user
// @Description Registers a new user`
// @Tags Auth
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/register-admin [post]
func (h *authHandler) RegisterAdmin(c *gin.Context) {
	h.log.Info("RegisterStudent handler called.")

	hash, err := hashing.HashPassword(config.Load().ADMIN_PASSWORD)

	res := pb.Message{
		Message: hash,
	}

	a, err := h.authService.RegisterAdmin(context.Background(), &res)
	if err != nil {
		h.log.Error("Error registering ADMIN", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(a)

	h.log.Info("Successfully registered user")
	c.JSON(http.StatusOK, models.Message{Massage: "FOR SURE!"})
}

// ResetPassword godoc
// @Summary Reset Password
// @Description it Reset your Password
// @Tags Auth
// @Param token body models.ResetPassReq true "enough"
// @Success 200 {object} string "message"
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/reset-password [post]
func (h *authHandler) ResetPassword(c *gin.Context) {
	h.log.Info("ResetPassword is working")
	var req models.ResetPassReq
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, err := h.redis.GetCodes(c, req.Email)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid or expired code " + err.Error()})
		return
	}
	if code != req.Code {
		h.log.Error("Invalid code")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid code "})
		return
	}

	d := pb.Email{
		Email: req.Email,
	}

	res, err := h.authService.GetUserByEmail(c, &d)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user" + err.Error()})
		return
	}
	res1 := pb.UpdatePasswordReq{
		Id:       res.UserId,
		Password: req.Password}

	_, err = h.authService.UpdatePassword(c, &res1)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password reset successfully"})
}
