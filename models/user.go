package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserData struct {
	Id   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}
