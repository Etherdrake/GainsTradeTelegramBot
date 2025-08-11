package streamprocessor

import (
	"GainsListenerMumbai/transformers"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func MatchChanges(client *mongo.Client,
	TpUpdatedChan chan transformers.TpUpdatedEventTransform,
	SlUpdatedChan chan transformers.SlUpdatedEventTransform,
	multiCollat string) {
	fmt.Println("Change Matcher has been started and channels are live!")

	go func() {
		for {
			select {
			case TpUpdatedEvent, ok := <-TpUpdatedChan:
				if !ok {
					fmt.Println("Error")
				}
				go func() {
					err := handleTpUpdated(client, TpUpdatedEvent)
					if err != nil {
						log.Println(err)
					}
				}()
			case SlUpdatedEvent, ok := <-SlUpdatedChan:
				if !ok {
					fmt.Println("Error")
				}
				go func() {
					err := handleSlUpdated(client, SlUpdatedEvent)
					if err != nil {
						log.Println(err)
					}
				}()
			}
		}
	}()
}
