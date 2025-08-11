package rdbhoot

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (rdbHoot *HootRedisClient) GetExperienceTimer(ctx context.Context, guid int64) (int64, error) {
	key := fmt.Sprintf("timers:%d", guid)
	unixTime, err := rdbHoot.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	unixTimeUint, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		return 0, err
	}
	return unixTimeUint, nil
}

func (rdbHoot *HootRedisClient) CheckExperienceTimerExpired(ctx context.Context, guid int64) (bool, error) {
	key := fmt.Sprintf("timers:%d", guid)
	unixTime, err := rdbHoot.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if unixTime == "0" {
		return true, nil
	} else {
		return false, nil
	}
}

func (rdbHoot *HootRedisClient) CheckExperienceTimerExists(ctx context.Context, guid int64) (bool, error) {
	key := fmt.Sprintf("timers:%d", guid)
	exists, err := rdbHoot.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (rdbHoot *HootRedisClient) InitExperienceTimer(ctx context.Context, guid int64) error {
	key := fmt.Sprintf("timers:%d", guid)
	err := rdbHoot.Set(ctx, key, 0, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetExperienceTimer(ctx context.Context, guid int64, unixTime int64) error {
	key := fmt.Sprintf("timers:%d", guid)
	err := rdbHoot.Set(ctx, key, strconv.FormatInt(unixTime, 10), 12*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rdbHoot *HootRedisClient) DeleteExperienceTimer(ctx context.Context, guid int64) error {
	key := fmt.Sprintf("timers:%d", guid)
	err := rdbHoot.Set(ctx, key, 0, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
