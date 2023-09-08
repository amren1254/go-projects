package model

type Status string
type Account_Type string

const (
	Active   Status       = "active"
	Inactive Status       = "inactive"
	Saving   Account_Type = "saving"
	Current  Account_Type = "current"
)

func (s *Status) String() string {
	return s.String()
}

func (a *Account_Type) String() string {
	return a.String()
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	// Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Status   Status `json:"status" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type Account struct {
	Account_Number string       `json:"account_number" binding:"required"`
	Account_Type   Account_Type `json:"account_type" binding:"required"`
	Total_Amount   string       `json:"total_amount" binding:"required"`
}
