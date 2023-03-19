package api

// The error sent to users in case of internal server error
const errInternalError = "سامانه با مشکل مواجه شد! لطفا کمی دیگر دوباره تلاش کنید..."

// errorResponse is a generic error response json
type errorResponse struct {
	Error string `json:"error"`
}
