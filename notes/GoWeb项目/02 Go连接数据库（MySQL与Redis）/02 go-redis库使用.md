## 三、Go语言操作Redis

### 3.1 Redis简介

Redis开源的内存数据库，提供多种不同类型的数据结构。

#### 3.1.1 Redis支持的数据结构

Redis支持诸如字符串（strings）、哈希（hashes）、列表（lists）、集合（sets）、带范围查询的排序集合（sorted sets）、位图（bitmaps）、hyperloglogs、带半径查询和流的地理空间索引等数据结构（geospatial indexes）。 

#### 3.1.2 Redis应用场景

- 缓存系统，减轻主数据库（MySQL）的压力。
- 计数场景，比如微博、抖音中的关注数和粉丝数。
- 热门排行榜，需要排序的场景特别适合使用ZSET。
- 利用LIST可以实现队列的功能。

#### 3.1.3 准备Redis环境

Docker启动Redis server：

```bash
$ docker run --name redis -p 6379:6379 -d redis
```

Docker启动redis-cli连接redis server：

```bash
$ docker run -it --network host --rm redis redis-cli
```

### 3.2 go-redis库

#### 3.2.1 安装

安装go-redis库（ https://github.com/go-redis/redis ）：

```bash
$ go get -u github.com/go-redis/redis
```

> `go-redis`不同于`redisgo`： `go-redis`支持连接哨兵及集群模式的Redis。 

#### 3.2.2 连接

**普通连接**

```go
// 全局rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping().Result()
	return
}
```

**连接Redis哨兵模式**

Redis集群中master挂掉之后，只能手动切换master；哨兵模式监控所有组内redis，当master挂掉之后，通过Raft算法选举出新的master。

```go
var rdb *redis.Client

// 初始化连接（哨兵模式）
func initClientSentinel() (err error) {
	redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	return
}
```

**连接Redis集群**

横向扩展，多个redis服务共同工作，一个主节点，其余都为从节点，主从之间通过数据同步存储完全相同的数据，主节点发生故障后，可使某个从节点成为主节点，保证性能。

```go
var rdbCluster *redis.ClusterClient

// 初始化连接（Redis集群）
func initClientCluster() (err error) {
	rdbCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdbCluster.Ping().Result()
	return
}
```

#### 3.2.3 基本使用

**set/get示例**

设置`score`，获取`score`，获取`name`。

```go
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
```

设置`user`，字段`name`和`age`。

```go
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
```

排行榜操作。

```go
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
```

> Redis命令：很重要。

#### 3.2.4 根据前缀获取Key

```go
vals, err := rdb.Keys(ctx, "prefix*").Result()
```

#### 3.2.5 执行自定义命令

```go
res, err := rdb.Do(ctx, "set", "key", "value").Result()
```

#### 3.2.6 按通配符删除key

当通配符匹配的key的数量不多时，可以使用`Keys()`得到所有的key在使用`Del`命令删除。 如果key的数量非常多的时候，可以搭配使用`Scan`命令和`Del`命令完成删除。

```go
ctx := context.Background()
iter := rdb.Scan(ctx, 0, "prefix*", 0).Iterator()
for iter.Next(ctx) {
	err := rdb.Del(ctx, iter.Val()).Err()
	if err != nil {
		panic(err)
	}
}
if err := iter.Err(); err != nil {
	panic(err)
}
```

#### 3.2.7 Pipeline

网络优化，客户端缓冲一堆命令并一次性发送到服务器（无法保证在事务中执行），从而节省了每个命令的网络往返时间（RTT）。

示例：

```go
pipe := rdb.Pipeline()

incr := pipe.Incr("pipeline_counter")
pipe.Expire("pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
```

以上将两个命令一次性发送到redis server执行，减少了一次RTT。也可使用`Pipelined`：

```go
var incr *redis.IntCmd
_, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
	incr = pipe.Incr("pipelined_counter")
	pipe.Expire("pipelined_counter", time.Hour)
	return nil
})
fmt.Println(incr.Val(), err)
```

#### 3.2.8 事务

Redis是单线程的，故每个命令都是原子的，但来自不同客户端的命令可依次执行（交替执行）。使用`Multi/exec`可以确保在`multi/exec`两个语句之间没有其他客户端正在执行命令。使用`TxPipeline`（类似于`Pipeline`，但内部使用`MULTI/EXEC`包裹排队的命令）：

```go
pipe := rdb.TxPipeline()

incr := pipe.Incr("tx_pipeline_counter")
pipe.Expire("tx_pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
```

以上在一个RTT下执行了以下命令：

```
MULTI
INCR pipeline_counter
EXPIRE pipeline_counts 3600
EXEC
```

或可使用：

```go
var incr *redis.IntCmd
_, err := rdb.TxPipelined(func(pipe redis.Pipeliner) error {
	incr = pipe.Incr("tx_pipelined_counter")
	pipe.Expire("tx_pipelined_counter", time.Hour)
	return nil
})
fmt.Println(incr.Val(), err)
```

#### 3.2.9 Watch

某些场景下，除使用`MULTI/EXEC`命令外，还需要配合使用`WATCH`命令。在用户使用`WATCH`命令监视某个键之后，直到该用户执行`EXEC`命令的这段时间里，如果有其他用户抢先对被监视的键进行了替换、更新、删除等操作，那么当用户尝试执行 `EXEC`的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。 

```go
Watch(fn func(*Tx) error, keys ...string) error
```

Watch方法接收一个函数和一个或多个key（被监视）作为参数。基本使用示例如下： 

```go
// 监视watch_count的值，并在值不变的前提下将其值+1
key := "watch_count"
err = client.Watch(func(tx *redis.Tx) error {
	n, err := tx.Get(key).Int()
	if err != nil && err != redis.Nil {
		return err
	}
	_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Set(key, n+1, 0)
		return nil
	})
	return err
}, key)
```

示例：V8版本官方文档中使用GET和SET命令以事务方式递增Key的值的示例，仅当Key的值不发生变化时提交一个事务。 

```go
func transactionDemo() {
	var (
		maxRetries   = 1000
		routineCount = 10
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Increment 使用GET和SET命令以事务方式递增Key的值
	increment := func(key string) error {
		// 事务函数
		txf := func(tx *redis.Tx) error {
			// 获得key的当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际的操作代码（乐观锁定中的本地操作）
			n++

			// 操作仅在 Watch 的 Key 没发生变化的情况下提交
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// 最多重试 maxRetries 次
		for i := 0; i < maxRetries; i++ {
			err := rdb.Watch(ctx, txf, key)
			if err == nil {
				// 成功
				return nil
			}
			if err == redis.TxFailedErr {
				// 乐观锁丢失 重试
				continue
			}
			// 返回其他的错误
			return err
		}

		return errors.New("increment reached maximum number of retries")
	}

	// 模拟 routineCount 个并发同时去修改 counter3 的值
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()
			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(context.TODO(), "counter3").Int()
	fmt.Println("ended with", n, err)
}
```

***示例***

设置`watch_count`的值。

```bash
set watch_count 7
```

监视`watch`并对其加1。

```go
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
```

