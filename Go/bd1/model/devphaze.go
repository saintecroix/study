package model

type DevPhaze struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CloseOpen   int    `json:"closeOpen"`
	Description string `json:"description"`
}
