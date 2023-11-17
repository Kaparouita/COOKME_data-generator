package core

import (
	"data-generator/internals/domain"
	"data-generator/internals/ports"
	"encoding/json"
	"os"

	"github.com/Kaparouita/models/models"
)

type GenerateService struct {
	db ports.DbRepo
}

func NewGenerateService(db ports.DbRepo) *GenerateService {
	return &GenerateService{
		db: db,
	}
}

func (srv *GenerateService) GenerateRecipes() ([]models.Recipe, error) {
	file, err := os.Open("./recipe_raw/recipe_raw.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create an empty Recipe struct
	var recipes []domain.RecipeTest

	// Unmarshal the JSON data into the Recipe struct
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&recipes); err != nil {
		return nil, err
	}

	return reformRecipeTests(recipes)
}

func reformRecipeTests(recipes []domain.RecipeTest) ([]models.Recipe, error) {
	var reformRecipes []models.Recipe
	for _, recipe := range recipes {
		reformRecipes = append(reformRecipes, *recipe.TransformRecipeTestToRecipe())
	}
	return reformRecipes, nil
}

func ExportIngridients(recipes []models.Recipe) (map[string]string, error) {
	ingridients := make(map[string]string)
	for _, recipe := range recipes {
		for _, ingridient := range recipe.RecipeInfo.Ingredients {
			ingridients[ingridient] = ingridient
		}
	}
	return ingridients, nil
}
