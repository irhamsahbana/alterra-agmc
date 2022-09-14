package entity

type User struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	IsMarried bool    `json:"is_married"`
	Phone     *string `json:"phone"`
}
