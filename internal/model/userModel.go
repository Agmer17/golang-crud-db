package model

type UserModel struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama"`
	Umur   int    `json:"umur"`
	Gender string `json:"gender"`
}
