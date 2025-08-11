package rdblocker

import (
	"HootCore/rdbhoot"
	"context"
	"log"
	"strconv"
)

// RemoveUserID removes a user ID from the active users set in Redis
func RemoveUserID(ctx context.Context, client *rdbhoot.HootRedisClient, userID int64) error {
	err := client.SRem(ctx, "active_users", strconv.FormatInt(userID, 10)).Err()
	if err != nil {
		log.Printf("Error removing user ID %d from active users set: %v", userID, err)
		return err
	}
	return nil
}
