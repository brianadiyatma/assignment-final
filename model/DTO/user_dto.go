package dto

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,gte=0"`
}

