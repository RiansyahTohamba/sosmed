package main

import (
	"context"
	"fmt"
	"sosmed/db"
	"time"

	"github.com/go-redis/redis/v9"
)

var ctx = context.TODO()

func main() {
	fmt.Println("sosmed")
	rdb := db.GetRedisClient()
	// readWrite(rdb.Client)
	// aquire lock?
	// printPipeline(rdb.Client)
	multiExample(rdb.Client)
}

// menggunakan set
func printPipeline(rdb *redis.Client) {

	cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 10; i++ {
			// Set(ctx context.Context, key string, value interface{}, expiration time.Duration)
			pipe.Set(ctx, fmt.Sprintf("key%d", i), "belum isi ", 0)
		}
		for i := 0; i < 10; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		// fmt.Println(cmd)
		switch cmd.(type) {
		case *redis.StatusCmd:
			fmt.Println(cmd.(*redis.StatusCmd).Val())
		case *redis.StringCmd:
			fmt.Println(cmd.(*redis.StringCmd).Val())
		default:
			fmt.Println("Error")
		}

	}
}

// ini masih fail
// panic: interface conversion: redis.Cmder is *redis.BoolCmd, not *redis.IntCmd
func readWrite(rdb *redis.Client) {

	pipe := rdb.Pipeline()

	incr := pipe.Incr(ctx, "pipeline_counter")

	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	cmds, err := pipe.Exec(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(incr.Val())

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.IntCmd).Val())
	}

}

// conversion interface
// var apapun interface{}
// kal := apapun.(string)
// fmt.Println(kal)
