package req

type CreateUser struct {
	DisplayName string  `json:"display_name"`
	Avatar      string  `json:"avatar"`
	Phone       string  `json:"phone"`
	Password    string  `json:"password"`
	Email       *string `json:"email"`
}
