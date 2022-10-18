package usersdto

type CreateUserRequest struct {
	Name string `json:"fullName" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Phone string `json:"phone" form:"phone"`
	Location string `json:"location" form:"location"`
	Image string `json:"image" form:"image"`
	Role string `json:"role" form:"role" validate:"required"`
}


type UpdateUserRequest struct {
	Name string `json:"fullName" form:"name"`
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
	Location string `json:"location" form:"location"`
	Image string `json:"image" form:"image"`
	Role string `json:"role" form:"role"`
}