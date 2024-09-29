package core

import (
	"context"
	"data-generator/internals/domain"
	"data-generator/internals/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func (srv *GenerateService) GetRecipesFromJson(file string) ([]models.Recipe, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var recipes []models.Recipe
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &recipes)
	if err != nil {
		return nil, err
	}

	for i := range recipes {
		srv.Db.SendToDB(&recipes[i])
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
	}
	return recipes, nil
}

func (srv *GenerateService) AddImages(file, apiKey string) error {
	// jsonFile, err := os.Open(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile.Close()

	recipes, err := srv.Db.GetRecipes()
	if err != nil {
		return err
	}
	// byteValue, _ := io.ReadAll(jsonFile)
	// json.Unmarshal(byteValue, &recipes)

	if len(recipes) == 0 {
		return fmt.Errorf("recipes is empty")
	}
	//0-800

	recipes, err = srv.GenerateImages(recipes, apiKey)
	if err != nil {
		return err
	}

	for i := range recipes {
		err = srv.Db.SendToDB(&recipes[i])
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
	}
	// err = srv.UpdateJson(recipes)
	// if err != nil {
	// 	return err
	// }
	fmt.Println("Successfully Generated Images")
	return nil
}

func (srv *GenerateService) GenerateImages(recipes []models.Recipe, apiKey string) ([]models.Recipe, error) {
	subscriptionKey := apiKey
	fmt.Printf("API Key: %s\n", subscriptionKey)

	endpoint := "https://api.pexels.com/v1/search"
	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(1)
	ctx := context.Background()
	index := 0
	total := 100

	for i := 0; i < total; i++ {

		if strings.Contains(recipes[i].Title, "&amp;") || strings.Contains(recipes[i].Description, "&amp;") {
			recipes[i].Title = strings.Replace(recipes[i].Title, "&amp;", "and", -1)
			recipes[i].Description = strings.Replace(recipes[i].Description, "&amp;", "and", -1)
			srv.Db.UpdateTitleDescription(&recipes[i])
		}

		// check if image starts with encrypted
		ok := strings.Contains(recipes[i].Image, "https://encrypted")
		if (recipes[i].Image != "" && recipes[i].Image != "Not found") && !ok {
			index++
			total++
			fmt.Printf("[LOG] Recipe %d already has an image\n", i)
			continue
		}

		sem.Acquire(ctx, 1)
		wg.Add(1)
		go func(i int, recipe *models.Recipe) {
			client := &http.Client{}

			defer func() {
				sem.Release(1)
				defer wg.Done()
			}()

			// if recipe.Image != "" && recipe.Image != "Not found" {
			// 	fmt.Printf("Recipe %d already has tried to retrieve image\n", i)
			// 	return
			// }

			query := recipe.Title
			encodedQuery := url.QueryEscape(query)
			uriQuery := endpoint + "?query=" + encodedQuery + "&per_page=1" + "&total_results=2" + "&size=medium"
			tries := 0

			fmt.Printf("[LOG] Requesting image for recipe %d\n", i)
			ok := strings.HasPrefix(recipes[i].Image, "https://encrypted")
			for (recipe.Image == "" || recipe.Image == "Not found" || ok) && tries < 5 {
				// Perform the Web request and get the response
				request, err := http.NewRequest("GET", uriQuery, nil)
				if err != nil {
					fmt.Println("[ERROR] ", err)
					return
				}

				request.Header.Add("Authorization", subscriptionKey)

				response, err := client.Do(request)
				if err != nil {
					fmt.Println("[ERROR] ", err)
					return
				}
				defer response.Body.Close()
				// Read the body of the response
				body, err := io.ReadAll(response.Body)
				if err != nil {
					fmt.Println("[ERROR] ", err)
				}

				type RespError struct {
					Message string `json:"message"`
					Status  int    `json:"status"`
				}
				respErr := RespError{}

				err = json.Unmarshal(body, &respErr)
				if err == nil && respErr.Status == 429 {
					fmt.Printf("[LOG] Rate limit exceeded for recipe %d, waiting 10 seconds, tries %d\n", i, tries)
					time.Sleep(time.Duration(10) * time.Second)
					tries++
					continue
				}
				var searchResponse domain.PexelsResponse
				err = json.Unmarshal(body, &searchResponse)
				if err != nil {
					recipe.Image = "Not found"
					fmt.Println("[ERROR] ", err)
					return
				}

				fmt.Printf("[LOG] Recipe %d has no image, trying to retrieve image, tries %d\n", i, tries)
				recipe.Image = ""
				for _, photo := range searchResponse.Photos {
					if photo.Src.Original != "" {
						fmt.Println("[LOG] Found image for recipe ", i)
						recipe.Image = photo.Src.Original
						ok = strings.HasPrefix(recipe.Image, "https://encrypted")
						break
					} else if photo.Src.Medium != "" {
						fmt.Println("[LOG] Found image for recipe ", i)
						recipe.Image = photo.Src.Medium
						ok = strings.HasPrefix(recipe.Image, "https://encrypted")
						break
					} else {
						fmt.Println("[LOG] No image found for recipe ", i)
					}
				}

			}
			fmt.Printf("[LOG] Finished retrieving image for recipe %d, image: %s\n", i, recipe.Image)
		}(i, &recipes[i])
	}
	wg.Wait()
	return recipes, nil
}

func (srv *GenerateService) UpdateJson(recipes []models.Recipe) error {
	file, err := json.MarshalIndent(recipes, "", "")
	if err != nil {
		return err
	}
	err = os.WriteFile("recipes.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
