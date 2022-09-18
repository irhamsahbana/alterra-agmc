package dto

type Register struct {
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	IsMarried bool    `json:"is_married"`
	Phone     *string `json:"phone"`
}
