package usersvc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func set(ctx context.Context, rdb *redis.ClusterClient, data User) error {
	strData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, data.Username, strData, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func get(ctx context.Context, rdb *redis.ClusterClient, key string) (*User, error) {
	var parsedValue User
	value, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, fmt.Errorf("key '%v' does not exist", key)
	} else if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(value), &parsedValue)

	return &parsedValue, nil
}

func delete(ctx context.Context, rdb *redis.ClusterClient, key string) error {
	numDeleted, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	} else if numDeleted == 0 {
		return fmt.Errorf("no items were deleted for key %v", key)
	}
	return nil
}
