package models

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Article struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Author      string `json:"author"`
	Description string `json:"description"`
	RecipeId    uint   `json:"-"`
}

type RecipeJsonInfo struct {
	Id            uint           `json:"id" gorm:"primaryKey"`
	CookingTime   int            `json:"cooking_time"`
	PrepTime      int            `json:"prep_time"`
	Serves        int            `json:"serves"`
	Keywords      pq.StringArray `json:"keywords,omitempty" gorm:"type:varchar[]"`
	Ratings       int            `json:"ratings"`
	NutritionInfo pq.StringArray `json:"nutrition_info,omitempty" gorm:"type:varchar[]"`
	Ingredients   pq.StringArray `json:"ingredients,omitempty" gorm:"type:varchar[]"`
	Courses       pq.StringArray `json:"courses,omitempty" gorm:"type:varchar[]"`
	Cuisine       string         `json:"cusine"`
	SkillLevel    string         `json:"skill_level"`
	PostDates     string         `json:"post_dates"`
	RecipeId      uint           `json:"-"`
}

type RecipeJson struct {
	Id            uint            `json:"id" gorm:"primaryKey"`
	Article       *Article        `json:"article" gorm:"foreignKey:RecipeId;constraint:onDelete:CASCADE;"`
	RecipeInfo    *RecipeJsonInfo `json:"recipe" gorm:"foreignKey:RecipeId;constraint:onDelete:CASCADE;"`
	Title         string          `json:"title"`
	Image         string          `json:"image"`
	FinalRecipeID uint            `json:"-"`
}

type FinalRecipe struct {
	Id                uint               `json:"id" gorm:"primaryKey"`
	CreatedAt         time.Time          `json:"create_at"`
	MarketIngredients []MarketIngredient `json:"market_ingridients" gorm:"foreignKey:FinalRecipeID"`
	Recipe            *Recipe            `json:"recipe" gorm:"foreignKey:FinalRecipeId;constraint:onDelete:CASCADE;"`
	TotalPrice        uint               `json:"total_price"`
}

type Recipe struct {
	Id            uint           `json:"id" gorm:"primaryKey"`
	Description   string         `json:"description"`
	CookingTime   int            `json:"cooking_time"`
	PrepTime      int            `json:"prep_time"`
	Serves        int            `json:"serves"`
	Keywords      pq.StringArray `json:"keywords,omitempty" gorm:"type:varchar[]"`
	Ratings       int            `json:"ratings"`
	NutritionInfo datatypes.JSON `gorm:"type:jsonb" json:"nutrition_info,omitempty"`
	Ingredients   pq.StringArray `json:"ingredients,omitempty" gorm:"type:varchar[]"`
	Courses       pq.StringArray `json:"courses,omitempty" gorm:"type:varchar[]"`
	Cuisine       string         `json:"cuisine"`
	SkillLevel    string         `json:"skill_level"`
	PostDates     string         `json:"post_dates"`
	Title         string         `json:"title" gorm:"uniqueIndex"`
	Price         float64        `json:"price"`
	OrderStatus   string         `json:"order_status"`
	Image         string         `json:"image"`
}

type NutritionInfo struct {
	SaturatedFat float64 `json:"saturated_fat" xml:"saturated_fat"`
	Protein      float64 `json:"protein" xml:"protein"`
	Fat          float64 `json:"fat" xml:"fat"`
	Kcal         float64 `json:"kcal" xml:"kcal"`
	AddedSugar   float64 `json:"added_sugar" xml:"added_sugar"`
	Carbohydrate float64 `json:"carbohydrate" xml:"carbohydrate"`
	Salt         float64 `json:"salt" xml:"salt"`
}

func (r *RecipeJson) TranformToRecipe() *Recipe {
	nutritionInfo := make(map[string]float64)
	for _, n := range r.RecipeInfo.NutritionInfo {
		typeOfNutrition, val := extractNutritionValue(n)

		switch typeOfNutrition {
		case "Saturated":
			nutritionInfo["saturated_fat"] = val
		case "Fat":
			nutritionInfo["fat"] = val
		case "Added":
			nutritionInfo["added_sugar"] = val
		case "Carbohydrate":
			nutritionInfo["carbohydrate"] = val
		case "Kcal":
			nutritionInfo["kcal"] = val
		case "Protein":
			nutritionInfo["protein"] = val
		case "Salt":
			nutritionInfo["salt"] = val
		}
	}

	nutritionInfoJson, _ := json.Marshal(nutritionInfo)

	return &Recipe{
		Description:   r.Article.Description,
		CookingTime:   r.RecipeInfo.CookingTime,
		PrepTime:      r.RecipeInfo.PrepTime,
		Serves:        r.RecipeInfo.Serves,
		Keywords:      r.RecipeInfo.Keywords,
		Ratings:       r.RecipeInfo.Ratings,
		NutritionInfo: nutritionInfoJson,
		Ingredients:   r.RecipeInfo.Ingredients,
		Courses:       r.RecipeInfo.Courses,
		Cuisine:       r.RecipeInfo.Cuisine,
		SkillLevel:    r.RecipeInfo.SkillLevel,
		PostDates:     r.RecipeInfo.PostDates,
		Title:         r.Title,
		Image:         r.Image,
	}
}

func (r *Recipe) UnmarshalNutritionInfo() (map[string]float64, error) {
	var nutritionInfo map[string]float64
	err := json.Unmarshal(r.NutritionInfo, &nutritionInfo)
	return nutritionInfo, err
}

func extractNutritionValue(info string) (string, float64) {
	words := strings.Split(info, " ")
	if len(words) > 1 {
		if words[0] == "Added" || words[0] == "Saturated" {
			val, err := strconv.ParseFloat(strings.TrimSuffix(words[2], "g"), 64)
			if err != nil {
				return "", 0
			}
			return words[0], float64(val)
		}
		val, err := strconv.ParseFloat(strings.TrimSuffix(words[1], "g"), 64)
		if err != nil {
			return "", 0
		}
		return words[0], float64(val)
	}
	return "", 0
}

func TranformRecipes(recipes []RecipeJson) []Recipe {
	var recipe []Recipe
	for _, r := range recipes {
		recipe = append(recipe, *r.TranformToRecipe())
	}
	return recipe
}

type Course string

const (
	MainCourse   Course = "Main course"
	Dessert      Course = "Dessert"
	Treat        Course = "Treat"
	AfternoonTea Course = "Afternoon tea"
	Supper       Course = "Supper"
	Breakfast    Course = "Breakfast"
	Starter      Course = "Starter"
	Canapes      Course = "Canapes"
	Lunch        Course = "Lunch"
	Dinner       Course = "Dinner"
	SideDish     Course = "Side dish"
	Snack        Course = "Snack"
	Brunch       Course = "Brunch"
	SoupCourse   Course = "Soup course"
	Vegetable    Course = "Vegetable course"
	FishCourse   Course = "Fish Course"
	Buffet       Course = "Buffet"
	Condiment    Course = "Condiment"
	Cocktails    Course = "Cocktails"
	Drink        Course = "Drink"
	PastaCourse  Course = "Pasta course"
	CheeseCourse Course = "Cheese Course"
)
