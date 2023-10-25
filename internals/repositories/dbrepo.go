package repositories

import (
	"data-generator/utils"

	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit"
	"github.com/Kaparouita/models/myrabbit/amqp"
)

type DbRepo struct {
	channel myrabbit.Channel
}

func NewDbRepo(handler *amqp.AmqpHandler) *DbRepo {
	ch, err := handler.PubConnection.Channel() // create a new channel
	utils.FailOnError(err, "Failed to open a channel")

	_, err = ch.QueueDeclare("kati-kati", false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")
	return &DbRepo{
		channel: ch,
	}
}

func (repo *DbRepo) SaveRecipes(recipes []models.Recipe) error {
	return nil
}