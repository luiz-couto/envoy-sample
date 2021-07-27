package api

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetModel struct {
	Username string `json:"username"`
}
