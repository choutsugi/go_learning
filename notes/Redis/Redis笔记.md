## 一、NoSQL数据库

### 1.1 NoSQL数据库简介

NoSQL，即非关系型数据库，使用K-V模式存储，不依赖业务逻辑方式。

特点：

- 不遵循SQL标准。

- 不支持事务ACID。
- 性能高于SQL。

### 1.2 NoSQL数据库对比

**Memcache**

特点：

- 数据不支持持久化。
- 仅支持K-V模式，类型单一。

**Redis**

- 支持持久化。
- 除K-V模式外，支持多种数据结构：string、list、set、hash、zset。
- 支持排序。
- 周期性将数据写入硬盘。
- 主从同步。

**MongoDB**

- 文档型数据库。
- 数据在内存中，若内存不足，则将不常用数据保存在硬盘中。
- K-V模式，但对V提供丰富的查询功能。
- 支持二进制数据和大型对象。
- 可替代RDRMS，或配置RDBMS存储特点数据。

## 二、Redis

### 2.1 应用场景

配合关系型数据库做高速缓存。

- 热门数据：降低数据库IO。
- 分布式架构：共享session共享。

多种数据结构存储持久化数据。

- 最新N个数据：通过List实现自然时间排序的数据。
- 排行榜：zset有序集合。
- 手机验证码：Expire过期。
- 计数器，秒杀：原子性，自增方法INCR、DECR。
- 去除大量数据中的重复数据：利用Set集合。
- 构建队列：利用List集合。
- 发布订阅消息系统：pub、sub模式。

### 2.2 安装

地址：http://redis.io 或 http://redis.cn

安装步骤：

1. 安装gcc
2. 解压redis压缩包到`/opt`目录：`tar -zxvf redis-6.2.1.tar.gz`
3. 进入解压目录执行`make`进行编译。
4. 使用`make install`安装。

安装目录：`/usr/local/bin`

- redis-benchmark：性能测试工具  
- redis-check-aof：修复AOF文件。
- redis-check-dump：修复dump.rdb文件。
- redis-sentinel：Redis集群使用。
- redis-server：Redis服务器。
- redis-cli：Redis客户端。

redis服务后台启动：

```bash
vim /opt/redis-3.2.5/redis.conf
# 设置daemonize no 为 yes

redis-server /opt/redis-3.2.5/redis.conf
```

客户端访问：

```bash
# 指定端口
redis-cli -p 6379
# 测试
ping
```

关闭redis：

```bash
redis-cli -p 6379 shutdown
```

### 2.3 Redis特点

单线程 + 多路IO复用技术。

## 三、常用数据类型

参见：http://www.redis.cn/commands.html  

### 3.1 Redis键（Key）

示例：

```bash
# 查看当前库所有key
keys *
# 判断某个key是否存在
exists key
# 查看key的类型
type key
# 删除key
del key
# 根据value选择非阻塞删除（将key从元数据中删除，真正的删除在后续异步操作）
unlink key
# 设置过期时间（单位：s）
expire key 10
# 查看过期时间（-1：永不过期，-2：已过期）
ttl key
```

其他命令：

- select：切换数据库。
- dbsize：查看当前数据库key的数量。
- flushdb：清空当前库。
- flushall：通杀数据库。

### 3.2 Redis字符串（String）

#### 特点

- 二进制安全：可存储图片或序列化对象。

- 字符串的value最大为512M。
- 底层数据结构为动态数组，自动扩容。

#### 常用命令

```bash
# 添加键值对
set <key> <value>
# 参数NX：key不存在时添加到数据库。
# 参数XX：key存在时添加到数据库，与NX互斥。
# 参数EX：设置超时秒数。
# 参数PX：设置超时毫秒数，与EX互斥。

# 查询键值
get <key>

# 追加内容
append <key> <value>

# 获取值的长度
strlen <key>

# 仅当key不存在时设置key值。
setnx <key> <value>

# 自增，仅可操作数字，且步长为1
incr <key>

# 自减，仅可操作数字，且步长为1
decr <key>

# 自定义步长
incrby / decrby <key> <步长>

# 同时设置一个或多个键的值
mset <key1> <value1> <key2> <value2>

# 同时获取一个或多个键的值
mget <key1> <key2>

# 同时设置一个或多个键的值，仅当所有给定key都不存在
msetnx <key1> <value1> <key2> <value2>

# 设置范围
setrange <key> <起始位置> <结束位置>

# 获取范围
getrange <key> <起始位置> <结束位置>

# 设置键值的同时设置过期时间(s)
setex <key> <过期时间> <value>

# 设置新值同时获取旧值
getset <key> <value>
```

### 3.3 Redis列表（List）

#### 特点

- 单键多值，单插入顺序排序，可插入新元素到列表的头部或尾部。
- 底层数据结构为双向链表：由ziplist组成的quicklist。

#### 常用命令

```bash
# 存入数据（左/右）
lpush / rpush <key> <value1> <value2> <value3> 

# 去除数据（左/右）
lpop / rpop <key>

# 从key1列表右边取出数据存入key2列表左边
rpoplpush <key1> <key2>

# 按照索引范围获取元素（从左至右）
lrange <key> <start> <stop>
# 示例：lrange mylist 0 -1 # 取出所有值，0表头部，-1表尾部。

# 按照索引下标获取元素（从左至右）
lindex <key> <index>

# 获取列表长度
llen <key>

# 在value后面插入值
linsert <key> <value> <newValue>

# 从左边删除n个value
lrem <key> <n> <value>

# 将列表下标为index的值进行替换
lset <key> <index> <value>
```

### 3.4 Redis集合（Set）

#### 特点

- set为string类型的无序集合，功能同list，区别在于list可自动排重（去除重复数据。）
- 底层数据结构：值为null的哈希表，增删查复杂度都为O(1)。

#### 常用命令

```bash
# 将多个元素添加到集合中，已存在元素将被忽略。
sadd <key> <value1> <value2>

# 取出该集合所有值
smembers <key>

# 判断集合是否有该值：有则返回1，无则返回0
sismember <key> <value>

# 返回集合元素个数
scard <key>

# 删除集合中的某些元素
srem <key> <value1> <value2>

# 随机删除一个值
spop <key>

# 从集合中读取n个值（不删除）
srandmember <key> <n>

# 将集合中的某个值移动到另一个集合
smove <source> <destination> value

# 返回两个集合的交集元素
sinter <key1> <key2>

# 返回两个集合的并集元素
sunion <key1> <key2>

# 返回两个集合的差集元素，不包含key2
sdiff <key1> <key2>
```

### 3.5 Redis哈希（Hash）

#### 特点

- 键值对集合。
- string类型的field和value的映射表，适用于存储对象。
- 通过 key + field 操作数据，不需要序列化。
- 底层数据结构：压缩链表和hash表。

#### 常用命令

```bash
# field键赋值
hset <key> <field> <value>

# 获取field的值
hget <key> <field>

# 批量设置hash值
hmset <key> <field1> <value1> <field2> <value2>

# 查看给定field是否存在
hexists <key> <field>

# 列出所有field
hkeys <key>

# 列出所有value
hvals <key>

# 自增
hincrby <key> <field> <步长>

# 仅当域不存在时设置值
hsetnx <key> <field> <value>
```

### 3.6 Redis有序集合（Zset）

#### 特点

- 无重复元素的字符串有序集合。
- 有序集合为每个成员关联评分（score），用于排序。

#### 常用命令

```bash
# 将一个或多个member元素及其score值加入到有序集合key中
zadd <key> <score1> <value1> <score2> <value2>

# 返回有序集合中，下标在<start> 和 <stop>之间的元素，通过WITHSCORES可返回score
zrange <key> <start> <stop> WITHSCORES

# 按score从小到大的顺序返回count个元素
zrangebyscore <key> minmax WITHSCORES 要获取的元素个数

# 按score从大到小的顺序返回count个元素
zrangebyscore <key> maxmin WITHSCORES 要获取的元素个数

# 设置增量，为元素的score设置增量
zincrby <key> <步长> <value> 

# 删除指定值的元素
zrem <key> <value>

# 统计score在某个区间的元素个数
zcount <key> <min> <max>

# 返回该值在集合中的排名
zrank <key> <value> 
```

示例：文章访问量排行榜

```bash
# redis-cli
127.0.0.1:6379> zadd topn 1000 v1 2000 v2 3000 v3
(integer) 3
127.0.0.1:6379> zrevrange topn 0 9 withscores
1) "v3"
2) "3000"
3) "v2"
4) "2000"
5) "v1"
6) "1000"
127.0.0.1:6379>
```

> zrange：递增排列。
>
> zrevrange：递减排列。

## 四、Redis配置文件

配置文件：`/opt/redis-3.2.5/redis.conf`

允许远程访问：

```bash
# 注释掉 bind=127.0.0.1 # 默认只接受本机连接请求
```

关闭保护模式：

```bash
protected-mode no
```

设置端口：

```bash
port 6379
```

修改连接队列，提高backlog值避免慢客户端连接问题：

```bash
tcp-backlog 511
```

客户端超时关闭：

```bash
timeout 0 # 0表示永不关闭
```

心跳检测：

```bash
tcp-keepalive 300
```

设置为后台进程：

```bash
daemonize yes
```

存放pid文件的路径（每个实例一个pid）：

```bash
pidfile /var/run/redis_6379.pid
```

loglevel：

```bash
loglevel notice #默认
```

logfile：日志文件名称

```bash
logfile
```

设置数据库数量：

```bash
database 16
```

设置密码：

```bash
requirepass foobared
```

设置客户端连接个数：

```bash
maxclients 1000
```

## 五、发布与订阅

客户端1订阅channel：

```bash
SUBSCRIBE channel1
```

客户端2向channel1发布消息：

```bash
publish channel1 hello
```

客户端1将收到消息。

## 六、事务与锁机制

### 6.1 事务定义

Redis事务是一个单独的隔离操作：事务中的所有命令都将序列化、按顺序执行，不会被其他客户端发送的命令请求打断。

> 作用：串联多个命令，防止插队。

## 6.2 Multi、Exec、discard

从输入Multi开始后的所有命令依次存入命令队列，但不执行，直至输入Exec后，队列中的命令依次执行。

### 6.3 错误处理

若组队中的某个命令报告错误，执行时整个队列都将被取消。

若执行阶段某个命令报告错误，则只有报错的命令不会被执行，其余命令都将执行，不回滚。

### 6.4 事务冲突

#### 悲观锁

每次拿到数据后加锁。

#### 乐观锁

每次拿数据时不加锁，但在更新时判断是否有人修改过该数据（通过版本号等机制）。Redis基于check-and-set机制实现事务。

#### 监视

**WATCH key1 key2**

监视一个或多个key，若事务执行前被监视的key发生改动，则打断事务。

**UNWATCH**

若WATCH后，执行了EXEC或DISCARD，则不在需要执行UNWATCH。

#### Redis事务三特性

- 单独隔离操作。
- 无隔离级别概念。
- 不保证原子性。

## 七、Redis持久化之RDB

### 7.1 简介

在指定的时间间隔内将内存中的数据集快照写入磁盘（dump.rdb）。

### 7.2 执行过程

Redis单独创建子进程进行持久化：将数据写入到临时文件，持久化过程结束后，使用该临时文件替换上次持久化的文件；主进程不参数该过程。

### 7.3 特点

最后一次持久化之后的数据可能丢失。

## 八、Redis持久化之AOF

### 8.1 简介

以日志形式记录（追加方式）每个写操作，redis重启后读取该文件重新构建数据。

### 8.2 执行过程

客户端请求写命令被append到AOF缓冲区内；AOF缓冲区根据AOF持久化策略将操作同步到磁盘的AOF文件中；当AOF文件大小超过重写策略时，对AOF文件执行重写，压缩AOF文件容量；Redis重启时加载AOF文件中的写操作重构数据。

### 8.3 说明

AOF默认不开启，需要到配置文件中开启。

## 九、主从复制

### 9.1 主从复制简介

主机数据更新后根据配置和策略，自动同步到备机的master/slaver机制，Master以写为主，Slave以读为主。

**特点**

- 读写分离，提高性能。
- 热备容灾。

### 9.2 实现

步骤

1. 拷贝多个redis配置文件并修改其端口信息用于启动多个redis服务。
2. 从机启动后执行设置：`slaveof 127.0.0.1 6379`

说明

从机重启后，需要重新设置。

### 9.3 主从转换

**主机变为从机**

```
slaveof 127.0.0.1 6379
```

**从机变为主机**

```
slaveof no one
```

### 9.4 哨兵机制

自动监控主机状态，主机故障后根据投票确认新的主机。

#### 实现

配置哨兵

```bash
# file: sential.conf
# 1表示至少需要多少个哨兵同一迁移
sentinel monitor mymaster 127.0.0.1 6379 1 
```

启动哨兵

```bash
redis-sential /redis/sentail.conf
```

原主机重启后变为从机。

## 十、Redis集群

### 10.1 集群简介

Redis集群实现对Redis的水平扩容，即启动N个节点，将整个数据库分布存储在N个节点中，每个节点存储总数据的1/N；通过分区提供程序可用性，即使集群中的某些节点失效或无法进行通讯，集群也可以继续处理命令请求。

### 10.2 集群特点

集群特点

- 实现扩容。
- 分摊压力。
- 无中心配置。

### 10.3 实现

#### 制作6个实例

配置修改：

- 开启后台运行：daemonize。
- 修改pid文件名、log文件名、RDB持久化文件名。
- 指定端口。
- 关闭Appendonly。
- 打开集群模式：cluster-enabled
- 设置节点配置文件名：cluster-config-file
- 设置节点失联时间（毫秒）：cluster-node-timeout

配置文件示例：

```
include /home/krain/redis.conf
port 6379
pidfile "/var/run/redis_6379.pid"
dbfilename "dump6379.rdb"
dir "/home/bigdata/redis_cluster"
logfile "/home/bigdata/redis_cluster/redis_err_6379.log"
cluster-enabled yes
cluster-config-file nodes-6379.conf
cluster-node-timeout 15000
```

基于以上配置文件拷贝5份并分别修改为各自的配置。

启动6个redis服务。

#### 将6个节点合成为一个集群

应先确保每个实例的nodes-xxxx.conf文件生成正常。

合成集群（安装ruby环境）：

```bash
cd /opt/redis-6.2.1/src
# 使用ruby脚本
./redis-trib.rb create --replicas 1 192.168.11.101:6379 192.168.11.101:6380 192.168.11.101:6381 192.168.11.101:6389 192.168.11.101:6390 192.168.11.101:6391
```

#### 集群策略连接

自动切换到对应的写主机。

```bash
# -c 参数：集群策略连接，自动重定向
redis-cli -c -p 6379
```

#### 查看集群信息

```bash
cluster nodes
```

#### 插槽

一个Redis集群包含16384个插槽（hash slot），数据库中每个键都属于其中的某个插槽。集群使用`CRC16(key)%16384`计算键所属的槽；集群中的每个节点负责处理一部分插槽。

## 十一、Redis应用问题解决

### 11.1 缓存穿透

问题：key对应的数据在缓存和数据源中都不存在。

解决方案：

- 对空值缓存：若查询返回数据为空，则缓存空结果（过期时间短）。
- 设置白名单：使用bitmaps类型定义一个可访问的名单，名单id作为bitmaps的偏移量，每次访问时和bitmap里的id进行比较，若不存在，则拦截。
- 采用布隆过滤器：所有可能存在的数据hash到bitmaps，进行拦截，降低底层查询压力。
- 实时监控：Redis命中率急速降低时，排查访问对象和数据设置黑名单。

### 11.2 缓存击穿

问题：某个key对应的数据存在，但在redis缓存中过期，大量并发请求发现缓存过期后从后端加载数据并回设到缓存。

解决方案：

- 预先设置热门数据。
- 实时调整。
- 使用排他锁。

### 11.3 缓存雪崩

问题：多个key对应的数据存在，但在redis缓存中过期，大量并发请求发现缓存过期后从后端加载数据并回设到缓存。

解决方案：

- 构建多级缓存。
- 使用锁和队列。
- 设置过期标志更新缓存。
- 分散缓存失效时间，避免集群集体失效。



