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
