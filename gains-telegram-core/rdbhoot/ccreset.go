package rdbhoot

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

func (rdbHoot *HootRedisClient) ResetUser(client *mongo.Client, guid int64) error {
	_, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	// Delete the data
	_, err := rdbHoot.JSONDel(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".").Result()
	if err != nil {
		return fmt.Errorf("error deleting coreCache struct in Redis: %v", err)
	}

	err = rdbHoot.InitUser(client, guid)
	if err != nil {
		return fmt.Errorf("error initializing user in Redis during ResetUser: %v", err)
	}

	return nil
}
