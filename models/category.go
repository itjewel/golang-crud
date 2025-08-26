package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryCreateRequest struct {
	Name string `json:"name" db:"name"`
}

type Response struct {
	Message string      `json:message`
	Status  int         `json:status`
	Data    interface{} `json:data`
}
