package domain

type Request struct {
	ID int `json:"id" gorm:"primaryKey"`
	TotalRecipes int `json:"total_recipes"`
}