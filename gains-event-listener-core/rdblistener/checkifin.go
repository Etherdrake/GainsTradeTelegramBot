package rdblistener

import (
	"context"
	"fmt"
	"log"
)

func (rdbHoot *HootRedisClient) CheckIfTradeExists(ctx context.Context, trader, tradeId string) (bool, error) {
	key := fmt.Sprintf("trades:%s:%s", trader, tradeId)
	existsCmd := rdbHoot.Exists(ctx, key)
	exists, err := existsCmd.Result()
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exists == 1, nil
}

func (rdbHoot *HootRedisClient) CheckIfOrderExists(ctx context.Context, trader, tradeId string) (bool, error) {
	key := fmt.Sprintf("orders:%s:%s", trader, tradeId)
	existsCmd := rdbHoot.Exists(ctx, key)
	exists, err := existsCmd.Result()
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exists == 1, nil
}
