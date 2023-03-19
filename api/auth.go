package api

import (
	"RayaneshBackend/internal/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// AuthLogin is the endpoint with logs in a user
func (api *API) AuthLogin(c *gin.Context) {
	// Parse data
	var request loginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}
	// Get user data from database
	user, err := api.Database.UserLogin(request.Email, request.Password)
	if err == database.ErrInvalidCredentials {
		c.JSON(http.StatusUnauthorized, errorResponse{Error: "ایمیل یا پسورد شما اشتباه است."})
		return
	}
	if err != nil {
		log.WithError(err).WithField("request", request).Error("cannot get login data from database")
		c.JSON(http.StatusInternalServerError, errorResponse{Error: errInternalError})
		return
	}
	// Create a session
	var response loginResponse
	response.Token, response.RefreshToken, err = api.Session.Store(user.ID, sessionTTL)
	if err != nil {
		log.WithError(err).WithField("userID", user.ID).Error("cannot store session")
		c.JSON(http.StatusInternalServerError, errorResponse{Error: errInternalError})
		return
	}
	// Done
	response.TimeToLive = uint32(sessionTTL / time.Second)
	response.UserID = user.ID
	c.JSON(http.StatusOK, response)
}

func (api *API) AuthSignup(c *gin.Context) {
	// TODO
}

func (api *API) AuthRefresh(c *gin.Context) {
	// TODO
}

func (api *API) AuthLogout(c *gin.Context) {
	// TODO
}
