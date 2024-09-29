package repositories

import (
	"data-generator/internals/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbRepo struct {
	*gorm.DB
}

func NewDbRepo() *DbRepo {
	db, err := connectDb()
	if err != nil {
		log.Fatal(err)
	}
	return &DbRepo{
		db,
	}
}

func connectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Mingrations")

	//Create the tables
	db.AutoMigrate(
		&models.Recipe{},
	)

	return db, nil
}

func (db *DbRepo) SendToDB(recipe *models.Recipe) error {
	var r models.Recipe
	db.Where("id = ?", recipe.Id).First(&r)
	if r.Title == "" {
		fmt.Printf("[LOG] did not find recipe with title %s\n", recipe.Title)
		return fmt.Errorf("did not find recipe with title %s", recipe.Title)
	} else {
		// update the recipe image
		r.Image = recipe.Image
		db.Save(&r)
	}
	return nil
}

func (db *DbRepo) UpdateTitleDescription(recipe *models.Recipe) error {
	var r models.Recipe
	db.Where("id = ?", recipe.Id).First(&r)
	if r.Title == "" {
		fmt.Printf("[LOG] did not find recipe with title %s\n", recipe.Title)
		return fmt.Errorf("did not find recipe with title %s", recipe.Title)
	} else {
		// update the recipe title and description
		r.Title = recipe.Title
		r.Description = recipe.Description
		db.Model(&r).Updates(models.Recipe{Title: recipe.Title, Description: recipe.Description})
	}
	return nil
}

func (db *DbRepo) GetRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe
	db.Find(&recipes)
	return recipes, nil
}
