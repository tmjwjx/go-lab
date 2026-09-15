package models

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Score  string `json:"score"`
	Enemy  string `json:"enemy"`
	Status int    `json:"status"`
}
