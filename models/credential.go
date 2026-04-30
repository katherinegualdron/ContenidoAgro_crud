package models

type Credential struct {
	ID int `json:"id"`
	ID_User int `json:"id_user"`
	Password string `json:"password"`
}