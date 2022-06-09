package entity

type Todo struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
