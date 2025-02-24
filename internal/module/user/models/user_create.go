package user_models

type UserCreate struct {
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required,min=1,max=20"`
	LastName  string `json:"last_name" binding:"required,min=1,max=20"`
}
