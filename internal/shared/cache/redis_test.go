package cache

import (
	"context"
	"fmt"
	"testing"

	"sync"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func newTestShardedRedis(t *testing.T, shardCount int) (*Redis, func()) {
	t.Helper()

	var shards []*redis.Client
	var servers []*miniredis.Miniredis

	for i := 0; i < shardCount; i++ {
		s := miniredis.RunT(t)
		servers = append(servers, s)

		client := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})
		shards = append(shards, client)
	}

	r := &Redis{
		shards: shards,
	}

	cleanup := func() {
		for _, s := range servers {
			s.Close()
		}
	}

	return r, cleanup
}

func TestRedis_SetGet(t *testing.T) {
	ctx := context.Background()
	ch := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup
	r, cleanup := newTestShardedRedis(t, 3)
	defer cleanup()
	wg.Add(1)
	go func() {
		r.Set(ctx, "user:1", "value1")
		r.Set(ctx, "user:2", "value2")
		wg.Done()
	}()
	wg.Wait()
	go func() {
		r.Get(ctx, "user:1", ch)
		r.Get(ctx, "user:2", ch2)
	}()

	fmt.Println("In progress")
	val := <-ch
	require.Equal(t, "value1", val)

	val2 := <-ch2
	require.Equal(t, "value2", val2)
}

func TestRedis_GetMissingKey(t *testing.T) {
	ctx := context.Background()
	ch := make(chan string)

	r, cleanup := newTestShardedRedis(t, 2)
	defer cleanup()

	go r.Get(ctx, "missing:key", ch)

	val := <-ch
	require.Empty(t, val)
}
