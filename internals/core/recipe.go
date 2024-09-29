package core

import (
	"data-generator/internals/domain"
	"data-generator/internals/models"
	"data-generator/internals/ports"
	"encoding/json"
	"os"
)

type GenerateService struct {
	Db ports.DbRepo
}

func NewGenerateService(db ports.DbRepo) *GenerateService {
	return &GenerateService{
		Db: db,
	}
}

func (srv *GenerateService) GenerateRecipes() ([]models.RecipeJson, error) {
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

func reformRecipeTests(recipes []domain.RecipeTest) ([]models.RecipeJson, error) {
	var reformRecipes []models.RecipeJson
	for _, recipe := range recipes {
		reformRecipes = append(reformRecipes, *recipe.TransformRecipeTestToRecipe())
	}
	return reformRecipes, nil
}
