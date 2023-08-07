package users

type User struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
