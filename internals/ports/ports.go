package ports

import "data-generator/internals/models"

type DbRepo interface {
	SendToDB(*models.Recipe) error
}

type GenerateService interface {
	GenerateRecipes() ([]models.Recipe, error)
	GetRecipesFromJson(file string) ([]models.Recipe, error)
	AddImages(file string) error
}
