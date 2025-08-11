package streamprocessor

import (
	"GainsListenerMumbai/api/orders"
	"GainsListenerMumbai/database"
	"GainsListenerMumbai/enumerate"
	"GainsListenerMumbai/transformers"
	"errors"
	"fmt"
	constants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func handleNftOrderInitiated(event transformers.NftOrderInitiatedEventTransform) error {
	//ordersArr, err := orders.GetOpenOrders(orders.PayloadOpenOrders{Chain: "polygon"})
	//if err != nil {
	//	log.Println("Error encountered retrieving the open orders.")
	//}
	//
	//convArr := orders.ConvertReceivedArrToTotalArr(ordersArr)
	//order := enumerate.FindMostRecentSingle(convArr)
	//
	//if order != nil {
	//	err := database.WriteTradeOrOrderToMongo(constants.OpenOrders, order, isArb)
	//	if err != nil {
	//		fmt.Println("Error writing order or trade to Mongo: ", err)
	//	}
	//} else {
	//	return errors.New("most recent order could not be found")
	//}
	return nil
}

func handleOpenLimitPlaced(db mongo.Database, event transformers.OpenLimitPlacedEventTransform) error {
	ordersArr, err := orders.GetOpenOrders(orders.PayloadOpenOrders{Chain: "polygon"})
	if err != nil {
		log.Println("Error encountered retrieving the open orders.")
	}

	convArr := orders.ConvertReceivedArrToTotalArr(ordersArr)
	order := enumerate.FindMostRecentSingle(convArr)

	if order != nil && *order != (orders.OpenOrdersHuman{}) {
		err := database.WriteTradeOrOrderToMongo(constants.OpenOrders, order)
		if err != nil {
			fmt.Println("Error writing order or trade to Mongo: ", err)
		}
	} else {
		return errors.New("most recent order could not be found")
	}

	return nil
}

func handleOpenLimitCanceled(db mongo.Database, event transformers.OpenLimitCanceledEventTransform, isArb bool) error {
	return nil
}
