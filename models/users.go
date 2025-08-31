package models

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type Order struct {
	Id    int    `json:"id"`
	Order string `json:"order"`
}

type OrderWithUser struct {
	User  Users  `json:"user"`
	Order *Order `json:"order,omitempty"`
}
