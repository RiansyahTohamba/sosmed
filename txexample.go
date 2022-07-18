package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

// apa itu command INCR?
func multiExample(rdb *redis.Client) {
	pipe := rdb.TxPipeline()
	// TxPipeline has wrapped with MULTI/EXEC
	incr := pipe.Incr(ctx, "tx_pipeline_counter")

	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	// sama seperti
	// //     MULTI
	// //     INCR pipeline_counter
	// //     EXPIRE pipeline_counts 3600
	// //     EXEC
	// //
	_, err := pipe.Exec(ctx)

	fmt.Println(incr.Val(), err)
}

func optimisExample() {
	const routineCount = 100
	// Increment using STRING data type
	// increment := func(key string) error {
	// 	txf := func(tx *redis.Tx) {

	// 	}
	// }
}

// const routineCount = 100

// // Transactionally increments key using GET and SET commands.
// increment := func(key string) error {
//     txf := func(tx *redis.Tx) error {
//         // get current value or zero
//         n, err := tx.Get(key).Int()
//         if err != nil && err != redis.Nil {
//             return err
//         }

//         // actual opperation (local in optimistic lock)
//         n++

//         // runs only if the watched keys remain unchanged
//         _, err = tx.TxPipelined(func(pipe redis.Pipeliner) error {
//             // pipe handles the error case
//             pipe.Set(key, n, 0)
//             return nil
//         })
//         return err
//     }

//     for retries := routineCount; retries > 0; retries-- {
//         err := rdb.Watch(txf, key)
//         if err != redis.TxFailedErr {
//             return err
//         }
//         // optimistic lock lost
//     }
//     return errors.New("increment reached maximum number of retries")
// }

// var wg sync.WaitGroup
// wg.Add(routineCount)
// for i := 0; i < routineCount; i++ {
//     go func() {
//         defer wg.Done()

//         if err := increment("counter3"); err != nil {
//             fmt.Println("increment error:", err)
//         }
//     }()
// }
// wg.Wait()

// n, err := rdb.Get("counter3").Int()
// fmt.Println("ended with", n, err)
