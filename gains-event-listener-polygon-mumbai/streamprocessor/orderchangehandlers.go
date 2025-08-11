package streamprocessor

import (
	"GainsListenerMumbai/database"
	"GainsListenerMumbai/transformers"
	"go.mongodb.org/mongo-driver/mongo"
)

func handleTpUpdated(client *mongo.Client, event transformers.TpUpdatedEventTransform) error {
	err := database.EditTp(client, event.Trader, event.PairIndex, event.Index, event.NewTp)
	if err != nil {
		return err
	}
	return nil
}

func handleSlUpdated(client *mongo.Client, event transformers.SlUpdatedEventTransform) error {
	err := database.EditSl(client, event.Trader, event.PairIndex, event.Index, event.NewSl)
	if err != nil {
		return err
	}
	return nil
}
