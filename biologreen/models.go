// biologreen/models.go

package biologreen

// FaceAuthResponse matches the successful API JSON response from your server.
// The `json:"..."` tags tell Go how to match these fields with the JSON data.
type FaceAuthResponse struct {
	UserID       int                    `json:"user_id"`
	IsNewUser    bool                   `json:"is_new_user"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// SignupRequest matches the JSON payload for the signup endpoint.
type SignupRequest struct {
	ImageBase64  string                 `json:"image_base64"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// LoginRequest matches the JSON payload for the login endpoint.
type LoginRequest struct {
	ImageBase64 string `json:"image_base64"`
}

// apiErrorResponse matches the JSON error response from your server.
type apiErrorResponse struct {
	Detail string `json:"detail"`
}
