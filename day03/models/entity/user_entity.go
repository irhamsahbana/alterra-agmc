package entity

type User struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	IsMarried bool    `json:"is_married"`
	Phone     *string `json:"phone"`
}
