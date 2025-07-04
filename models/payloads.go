package models

type SigninPayload struct{
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`

}

type SignupPayload struct{
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`

}