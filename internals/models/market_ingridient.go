package models

import "time"

// MarketIngredient is a struct that contains the information of a recipe's ingredient
// that is needed to be bought from a supermarket
type MarketIngredient struct {
	Id          uint        `json:"ingridient_id" gorm:"primaryKey"`
	CreatedAt   time.Time   `json:"create_at"`
	Brand       string      `json:"brand"` //after
	Name        string      `json:"name"`
	Price       float64     `json:"price"` //random
	ExtraInfo   string      `json:"extra_info"`
	Supermarket Supermarket `json:"market"` //random
	Url         string      `json:"url"`    //after
	Image       string      `json:"image"`  //after
}

// Supermarket is a type that contains the names of the supermarkets
type Supermarket string

const (
	Xalkiadakis  Supermarket = "Xalkiadakis"
	Masoutis     Supermarket = "Masoutis"
	AB           Supermarket = "AB"
	Lidl         Supermarket = "Lidl"
	Carrefour    Supermarket = "Carrefour"
	Sklavenitis  Supermarket = "Sklavenitis"
	Galaxias     Supermarket = "Galaxias"
	Kritikos     Supermarket = "Kritikos"
	Vasilopoulos Supermarket = "Vasilopoulos"
)
