package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 全局rdb变量
var rdb *redis.Client
var rdbCluster *redis.ClusterClient

// 初始化连接（普通连接）
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, // 连接池大小
	})

	_, err = rdb.Ping().Result()
	return
}

// 初始化连接（哨兵模式）
func initClientSentinel() (err error) {
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	return
}

// 初始化连接（Redis集群）
func initClientCluster() (err error) {
	rdbCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdbCluster.Ping().Result()
	return
}

// set/get int
func redisExample1() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	// rdb.Get().Val()：返回值（查找无结果返回零值）。
	//val1 := rdb.Get("score").Val()
	// rdb.Get().Result()：返回值和错误。
	val1, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val1)

	val2, err := rdb.Get("name").Result()
	if err != redis.Nil {
		fmt.Println("name dose not exist.")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

// set/get hash
func redisExample2() {
	info := rdb.HSet("user", "name", "lettredamour")
	fmt.Println(info)

	info = rdb.HSet("user", "age", "17")
	fmt.Println(info)

	v1, err := rdb.HGetAll("user").Result()
	if err != nil {
		fmt.Printf("hgetall failed, err:%v\n", err)
		return
	}
	fmt.Println(v1)

	v2 := rdb.HMGet("user", "name", "age").Val()
	fmt.Println(v2)

	v3 := rdb.HGet("user", "age").Val()
	fmt.Println(v3)
}

// 排行榜
func redisExample3() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 99.0, Member: "C"},
		redis.Z{Score: 97.0, Member: "Python"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "C++"},
	}

	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d success.\n", num)

	// Golang分数+1
	newScore, err := rdb.ZIncrBy(zsetKey, 1.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	result, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range result {
		fmt.Println(z.Member, z.Score)
	}

	// 取分数在95~100范围内的
	ops := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	result, err = rdb.ZRangeByScoreWithScores(zsetKey, ops).Result()
	if err != nil {
		fmt.Printf("zrangebyscorewithscores failed, err:%v\n", err)
		return
	}
	for _, z := range result {
		fmt.Println(z.Member, z.Score)
	}
}

// Watch 监视watch_count
func watchDemo() {
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 业务逻辑
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Printf("tx exec failed, err:%v\n", err)
		return
	}
	fmt.Println("tx exec success.")
}

// 事务
//func transactionDemo() {
//	var (
//		maxRetries   = 1000
//		routineCount = 10
//	)
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	// Increment 使用GET和SET命令以事务方式递增Key的值
//	increment := func(key string) error {
//		// 事务函数
//		txf := func(tx *redis.Tx) error {
//			// 获得key的当前值或零值
//			n, err := tx.Get(ctx, key).Int()
//			if err != nil && err != redis.Nil {
//				return err
//			}
//
//			// 实际的操作代码（乐观锁定中的本地操作）
//			n++
//
//			// 操作仅在 Watch 的 Key 没发生变化的情况下提交
//			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
//				pipe.Set(ctx, key, n, 0)
//				return nil
//			})
//			return err
//		}
//
//		// 最多重试 maxRetries 次
//		for i := 0; i < maxRetries; i++ {
//			err := rdb.Watch(ctx, txf, key)
//			if err == nil {
//				// 成功
//				return nil
//			}
//			if err == redis.TxFailedErr {
//				// 乐观锁丢失 重试
//				continue
//			}
//			// 返回其他的错误
//			return err
//		}
//
//		return errors.New("increment reached maximum number of retries")
//	}
//
//	// 模拟 routineCount 个并发同时去修改 counter3 的值
//	var wg sync.WaitGroup
//	wg.Add(routineCount)
//	for i := 0; i < routineCount; i++ {
//		go func() {
//			defer wg.Done()
//			if err := increment("counter3"); err != nil {
//				fmt.Println("increment error:", err)
//			}
//		}()
//	}
//	wg.Wait()
//
//	n, err := rdb.Get(context.TODO(), "counter3").Int()
//	fmt.Println("ended with", n, err)
//}

func main() {
	// 普通连接
	err := initClient()
	if err != nil {
		fmt.Println("init redis client failed, err:", err)
		return
	}

	// 集群模式
	//_ = initClientSentinel()

	// 哨兵模式
	//_ = initClientCluster()

	// 程序退出时释放相关资源
	if rdb != nil {
		defer func(rdb *redis.Client) {
			err := rdb.Close()
			if err != nil {
				panic(err)
			}
		}(rdb)
	} else if rdbCluster != nil {
		defer func(rdbCluster *redis.ClusterClient) {
			err := rdbCluster.Close()
			if err != nil {
				panic(err)
			}
		}(rdbCluster)
	}

	// get/set示例
	redisExample1()
	redisExample2()
	redisExample3()

	// watch
	watchDemo()
}
