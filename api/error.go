package api

// The error sent to users in case of internal server error
const errInternalError = "سامانه با مشکل مواجه شد! لطفا کمی دیگر دوباره تلاش کنید..."

// errProfileMustBeJpg is sent to users if they upload profile pics which are not jpg
const errProfileMustBeJpg = "عکس پروفایل باید به صورت jpg باشد."

// errorResponse is a generic error response json
type errorResponse struct {
	Error string `json:"error"`
}
