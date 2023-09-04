package models

type Password struct {
	Current string `json:"password"`
	New     string `json:"new"`
}
