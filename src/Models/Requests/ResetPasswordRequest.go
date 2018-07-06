package requests

type ResetPasswordRequest struct {
	AccountID       int    `json:"accountId"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
