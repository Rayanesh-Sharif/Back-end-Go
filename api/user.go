package api

import (
	"RayaneshBackend/internal/database"
	"RayaneshBackend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
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

// UserChangeProfilePic will change the profile pic of user
func (api *API) UserChangeProfilePic(c *gin.Context) {
	userID := c.MustGet(userIDAuthedContext).(uint32)
	// Limit the size of uploaded image: https://github.com/gin-gonic/gin/issues/2898#issuecomment-939976866
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxProfilePictureSize)
	// Download the file
	file, err := c.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	// At first, we read the first 512 bytes of file
	userPicLocation := api.userProfilePicLocation(userID)
	err = util.CheckMimeAndSaveFile(file, userPicLocation, "image/jpeg")
	if err != nil { // on any error, delete the file
		_ = os.Remove(userPicLocation)
	}
	if errors.Is(err, util.ErrMimeMismatch) { // on mime mismatch just tell the user that image must be jpg
		c.JSON(http.StatusBadRequest, errorResponse{errProfileMustBeJpg})
		return
	}
	if err != nil { // this is some other error
		c.JSON(http.StatusInternalServerError, errorResponse{errInternalError})
		log.WithError(err).WithField("userID", userID).Error("cannot update user's profile pic")
		return
	}
	// Done
	c.JSON(http.StatusOK, gin.H{})
}

// UserDeleteProfilePic will delete the profile pic of user
func (api *API) UserDeleteProfilePic(c *gin.Context) {
	userID := c.MustGet(userIDAuthedContext).(uint32)
	// We simply delete the file if it exists.
	// Also, fuck the errors
	_ = os.Remove(api.userProfilePicLocation(userID))
	// We good
	c.JSON(http.StatusOK, gin.H{})
}
