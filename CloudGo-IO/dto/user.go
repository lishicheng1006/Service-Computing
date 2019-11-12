package dto

type UserSignUp struct {
	StuId    string `json:"stu_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
