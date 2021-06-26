package routers

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type AirlineSearchDTO struct {
	Origin      int `json:"origin" binding:"required"`
	Destination int `json:"dest" binding:"required"`
	Page        int `json:"page"`
}