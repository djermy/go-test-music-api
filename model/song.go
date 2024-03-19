package model

type Song struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
}
