package models

type Person struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsMarried bool   `json:"married"`
}
