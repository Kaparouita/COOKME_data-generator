package domain

import "github.com/Kaparouita/models/models"


type RecipeTest struct {
	Page struct {
		Article struct {
			Author      string `json:"author"`
			Description string `json:"description"`
			ID          string `json:"id"`
		} `json:"article"`
		Recipe struct {
			Collections   []string `json:"collections"`
			CookingTime   int      `json:"cooking_time"`
			PrepTime      int      `json:"prep_time"`
			Serves        int      `json:"serves"`
			Keywords      []string `json:"keywords"`
			Ratings       int      `json:"ratings"`
			NutritionInfo []string `json:"nutrition_info"`
			Ingredients   []string `json:"ingredients"`
			Courses       []string `json:"courses"`
			Cuisine       string   `json:"cusine"` 
			DietTypes     []string `json:"diet_types"`
			SkillLevel    string   `json:"skill_level"`
			PostDates     string   `json:"post_dates"`
		} `json:"recipe"`
		Channel string `json:"channel"`
		Title   string `json:"title"`
	} `json:"page"`
}

func (recipeTest *RecipeTest) TransformRecipeTestToRecipe() *models.Recipe {
    article := &models.Article{
        Author:      recipeTest.Page.Article.Author,
        Description: recipeTest.Page.Article.Description,
    }

    recipeInfo := &models.RecipeInfo{
        CookingTime:   recipeTest.Page.Recipe.CookingTime,
        PrepTime:      recipeTest.Page.Recipe.PrepTime,
        Serves:        recipeTest.Page.Recipe.Serves,
        Keywords:      recipeTest.Page.Recipe.Keywords,
        Ratings:       recipeTest.Page.Recipe.Ratings,
        NutritionInfo: recipeTest.Page.Recipe.NutritionInfo,
        Ingredients:   recipeTest.Page.Recipe.Ingredients,
        Courses:       recipeTest.Page.Recipe.Courses,
        Cuisine:       recipeTest.Page.Recipe.Cuisine,
        SkillLevel:    recipeTest.Page.Recipe.SkillLevel,
        PostDates:     recipeTest.Page.Recipe.PostDates,
    }

    recipe := &models.Recipe{
        Article:   article,
        RecipeInfo: recipeInfo,
        Title:     recipeTest.Page.Title,
    }

    return recipe
}