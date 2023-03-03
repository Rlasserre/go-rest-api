package models

type Personage struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personages []Personage
