package api

import (
	"RayaneshBackend/internal/database"
	"RayaneshBackend/pkg/session"
	"RayaneshBackend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// AuthLogin is the endpoint with logs in a user
func (api *API) AuthLogin(c *gin.Context) {
	// Parse data
	var request loginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	// Get user data from database
	user, err := api.Database.UserLogin(request.Email, request.Password)
	if err == database.ErrInvalidCredentials {
		c.JSON(http.StatusUnauthorized, errorResponse{"ایمیل یا پسورد شما اشتباه است."})
		return
	}
	if err != nil {
		log.WithError(err).WithField("request", request).Error("cannot get login data from database")
		c.JSON(http.StatusInternalServerError, errorResponse{errInternalError})
		return
	}
	// Create a session
	var response loginResponse
	response.Token, response.RefreshToken, err = api.Session.Store(user.ID, sessionTTL)
	if err != nil {
		log.WithError(err).WithField("userID", user.ID).Error("cannot store session")
		c.JSON(http.StatusInternalServerError, errorResponse{errInternalError})
		return
	}
	// Done
	response.TimeToLive = uint32(sessionTTL / time.Second)
	response.UserID = user.ID
	c.JSON(http.StatusOK, response)
}

// AuthSignup is called when user wants to sign up
func (api *API) AuthSignup(c *gin.Context) {
	// Parse data
	var request registerRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	// Check email
	request.Email = util.CheckEmail(request.Email)
	if request.Email == "" {
		c.JSON(http.StatusBadRequest, errorResponse{"ایمیل شما نامعتبر است."})
		return
	}
	// Register user
	err := api.Database.UserRegister(request.Email, request.Username, request.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, errorResponse{"این ایمیل در سامانه وجود دارد!"})
			return
		}
		// General error
		c.JSON(http.StatusInternalServerError, errorResponse{errInternalError})
		log.WithError(err).WithField("request", request).Error("cannot signup user")
		return
	}
	// Done
	c.JSON(http.StatusOK, gin.H{})
}

func (api *API) AuthRefresh(c *gin.Context) {
	// Get the refresh token
	var request refreshTokenRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	// Refresh
	newToken, err := api.Session.Refresh(request.RefreshToken, sessionTTL)
	if err == session.ErrNotFound {
		c.JSON(http.StatusForbidden, errorResponse{"ریفرش توکن نامعتبر"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{errInternalError})
		log.WithError(err).WithField("refreshToken", request.RefreshToken).Error("cannot refresh token of user")
		return
	}
	// Send tokens
	c.JSON(http.StatusOK, loginResponse{
		Token:        newToken,
		RefreshToken: request.RefreshToken,
		TimeToLive:   uint32(sessionTTL / time.Second),
	})
}

func (api *API) AuthLogout(c *gin.Context) {
	accessToken := getAccessTokenFromHeaders(c)
	err := api.Session.Delete(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{errInternalError})
		log.WithError(err).WithField("access_token", accessToken).Error("cannot sign out user")
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
