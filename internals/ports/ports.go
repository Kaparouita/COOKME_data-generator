package ports

import "github.com/Kaparouita/models/models"


type DbRepo interface {
	SaveRecipes(recipes []models.Recipe) error
}

type GenerateService interface {
	GenerateRecipes() ([]models.Recipe, error)
}