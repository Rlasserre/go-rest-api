package models

type Personage struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personages []Personage
