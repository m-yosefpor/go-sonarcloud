package authentication

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// LogoutRequest Logout a user.
type LogoutRequest struct{}

// ValidateRequest Check credentials. Returns true for anonymous user.
type ValidateRequest struct{}

// ValidateResponse is the response for ValidateRequest
type ValidateResponse struct {
	Valid bool `json:"valid,omitempty"`
}
