package cache

import (
	"context"

	"sync"

	"hash/fnv"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	shards []*redis.Client
	mu     sync.RWMutex
}

func NewShardRedis(dbCount int, addr string) *Redis {
	shards := make([]*redis.Client, 0, dbCount)

	for db := range dbCount {

		shards = append(shards, redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   db,
		}))
	}
	return &Redis{shards: shards}
}

func (r *Redis) hasher(key string) int {
	hasher := fnv.New32a()
	_, _ = hasher.Write([]byte(key))

	return int(hasher.Sum32())
}
func (r *Redis) GetShardID(key string) *redis.Client {
	shardId := r.hasher(key) % len(r.shards)

	return r.shards[shardId]
}

func (r *Redis) Set(ctx context.Context, key string, val string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.GetShardID(key).Set(ctx, key, val, 0)
}

func (r *Redis) Get(ctx context.Context, key string, ch chan string) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	defer close(ch)

	res, _ := r.GetShardID(key).Get(ctx, key).Result()

	ch <- res

}
