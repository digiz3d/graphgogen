package model

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Shows    []*Show `json:"shows"`
}