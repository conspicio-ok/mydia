package service

type User struct {
	Pseudo		string	`json:"pseudo"`
	Password	string	`json:"password"`
}

type CreateUserResponse struct {
	ID		int		`json:"id"`
	Pseudo	string	`json:"pseudo"`
	Message	string	`json:"message"`
}

type UpdateUserResponse struct {
	Pseudo	string	`json:"pseudo"`
	Message	string	`json:"message"`
}

type MessageResponse struct {
	Status	string	`json:"status"`
	Message	string	`json:"message"`
}

type JWTClaims struct {
    Username string  `json:"username"`
    Admin  bool `json:"admin"`
    jwt.RegisteredClaims
}
