package main

import (
	"data-generator/internals/domain"
	"encoding/json"
	"fmt"
	"os"
)


func main() {

	file, err := os.Open("./recipe_raw/recipe_raw.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

	// Create an empty Recipe struct
	var recipes []domain.RecipeTest

	// Unmarshal the JSON data into the Recipe struct
	decoder := json.NewDecoder(file)
    if err := decoder.Decode(&recipes); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }
i := 0
	// Now you can work with the populated Recipe struct
	for _,recipe := range recipes{
		i++
		fmt.Printf("%d : Author: %s, ",i, recipe.Page.Article.Author)
		fmt.Printf("Description: %s ,", recipe.Page.Article.Description)
		fmt.Printf("Cooking Time: %d\n", recipe.Page.Recipe.CookingTime)
	}
}
