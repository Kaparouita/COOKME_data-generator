package ports

import (
	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit"
)

type DbRepo interface {
	SendToDB(*models.Recipe) error
}

type GenerateService interface {
	GenerateRecipes() ([]models.Recipe, error)
	GetRecipesFromJson(file string) ([]models.Recipe, error)
	AddImages(file string) error
}

type Handler interface {
	GetRecipes(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
}
