package rdbhoot

import (
	"context"
	"fmt"
	"strconv"
)

func (rdbHoot *HootRedisClient) CheckExperienceExists(ctx context.Context, guid int64) (bool, error) {
	key := fmt.Sprintf("experience:%d", guid)
	exists, err := rdbHoot.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (rdbHoot *HootRedisClient) SetExperience(ctx context.Context, guid int64, experience uint64) error {
	key := fmt.Sprintf("experience:%d", guid)
	err := rdbHoot.Set(ctx, key, strconv.FormatUint(experience, 10), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rdbHoot *HootRedisClient) GetExperience(ctx context.Context, guid int64) (string, error) {
	key := fmt.Sprintf("experience:%d", guid)
	val, err := rdbHoot.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (rdbHoot *HootRedisClient) IncrementExperience(ctx context.Context, guid int64, experience uint64) error {
	key := fmt.Sprintf("experience:%d", guid)
	_, err := rdbHoot.IncrBy(ctx, key, int64(experience)).Result()
	if err != nil {
		return err
	}
	return nil
}
