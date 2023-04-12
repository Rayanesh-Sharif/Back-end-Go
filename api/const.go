package api

import "time"

const sessionTTL = time.Hour * 6

// ProfilePicturesFolder is the folder in the static folder which we use to store profile pictures.
// Profile pictures are simply stored as userID.jpg
const ProfilePicturesFolder = "profiles"

// maxProfilePictureSize is the maximum size of profile picture.
// Currently, this is 200KB
const maxProfilePictureSize = 200 * 1024
