package api

import (
	"fmt"
	"path"
)

// userProfilePicLocation will return the local location of the user profile picture
func (api *API) userProfilePicLocation(userID uint32) string {
	return path.Join(api.StaticFolder, ProfilePicturesFolder, fmt.Sprintf("%d.jpg", userID))
}
