package api

// loginRequest is the request sent to us when logging in
type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// loginResponse is the response of a successful login
type loginResponse struct {
	UserID       uint32 `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	TimeToLive   uint32 `json:"ttl"`
}

// registerRequest is the request body of registering the user
type registerRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

// refreshTokenRequest is the query values to refresh the token
type refreshTokenRequest struct {
	RefreshToken string `query:"refresh_token" binding:"required"`
}

// changePasswordRequest is the request body of a change password
type changePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
