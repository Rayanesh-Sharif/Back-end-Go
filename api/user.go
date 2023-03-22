package api

import (
	"RayaneshBackend/internal/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// UserChangePassword is an endpoint to change the password of a user
func (api *API) UserChangePassword(c *gin.Context) {
	userID := c.MustGet(userIDAuthedContext).(uint32)
	// Parse data
	var request changePasswordRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	// Change the password
	err := api.Database.UserChangePassword(userID, request.OldPassword, request.NewPassword)
	if err != nil {
		if err == database.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, errorResponse{"رمزعبور قبلی شما اشتباه است."})
			return
		}
		log.WithError(err).WithField("request", request).Error("cannot update user's password")
		c.JSON(http.StatusInternalServerError, errorResponse{errInternalError})
		return
	}
	// Done
	c.JSON(http.StatusOK, gin.H{})
}
