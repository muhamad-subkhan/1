package usersdto

type UserResponse struct {
	ID int `json:"id"`
	Name string `json:"fullName" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Phone string `json:"phone" form:"phone"`
	Location string `json:"location" form:"location"`
	Image string `json:"image" form:"image"`
	Role string `json:"role" form:"role" validate:"required"`
}
