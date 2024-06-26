package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// Struct Digunakan untuk menerima input dari Login User
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Struct ini digunakan untuk cek avaibility email
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
