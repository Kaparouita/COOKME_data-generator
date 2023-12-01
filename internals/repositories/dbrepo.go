package repositories

import (
	"context"
	"data-generator/utils"

	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit"
	"github.com/Kaparouita/models/myrabbit/amqp"
)

type DbRepo struct {
	pubCh myrabbit.Channel
	subCh myrabbit.Channel
	msgs  <-chan myrabbit.Delivery
}

func NewDbRepo(handler *amqp.AmqpHandler) *DbRepo {
	pubCh, err := handler.PubConnection.Channel() // create a new channel
	utils.FailOnError(err, "Failed to open a channel")
	subCh, err := handler.SubConnection.Channel() // create a new channel
	utils.FailOnError(err, "Failed to open a channel")

	_, err = pubCh.QueueDeclare("db_manager.save.recipe", false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")

	_, err = subCh.QueueDeclare("db_manager.save.recipe.rcv", false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")

	return &DbRepo{
		pubCh: pubCh,
		subCh: subCh,
	}
}

func (adapter *DbRepo) SendToDB(recipe *models.Recipe) error {
	err := adapter.pubCh.PublishJSON(context.Background(), "", "db_manager.save.recipe", false, false, recipe, myrabbit.Publishing{})
	return err
}
