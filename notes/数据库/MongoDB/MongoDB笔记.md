参见：https://www.yiibai.com/mongodb

## 一、MongoDB简介

### 1.1 简介

MongoDB官网：http://www.mongodb.com/

可视化工具：https://robomongo.org/download

MongoDB特点：

- 跨平台。
- 面向文档。
- 高性能、高可用、易扩展。

### 1.2 基本概念

**数据库**

数据库是一个集合的物理容器，一个MongoDB服务器有多个数据库。

**集合**

集合是一组MongoDB的文件，等效于RDBMS表，集合中的文档可以有不同的字段（虽然可以，但不建议）。

**文档**

文档是一组键值对，具有动态模式（同一集合的文档不要求具有相同的字段或结构，且相同的字段可以保持不同类型的数据）。

文档示例：博客网站。

```json
{
   _id: ObjectId(7df78ad8902c)
   title: 'MongoDB Overview', 
   description: 'MongoDB is no sql database',
   by: 'LAB tutorial',
   url: 'http://www.lab.com',
   tags: ['mongodb', 'database', 'NoSQL'],
   likes: 100, 
   comments: [	
      {
         user:'user1',
         message: 'My first comment',
         dateCreated: new Date(2011,1,20,2,15),
         like: 0 
      },
      {
         user:'user2',
         message: 'My second comments',
         dateCreated: new Date(2011,1,25,7,45),
         like: 5
      }
   ]
}
```

### 1.3 安装

使用docker安装：

```bash
$ docker pull mongo:latest
$ docker run -itd --name mongo -p 27017:27017 mongo --auth # 开启密码验证
$ docker exec -it mongo mongo admin
> db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});
> db.auth('admin', '123456')
```

## 二、增删改查

### 2.1 数据库操作

**创建数据库**

使用`use DATABASE_NAME`创建名为`DATABASE_NAME`的数据库，若数据库不存在则创建，否则返回现有的数据库。

示例：

```
> use newdb
switched to db newdb
```

**查看当前选择的数据库**

```
> db
newdb
```

**查询数据库列表**

不显示空数据库（至少要有一个文档）

```
> show dbs
```

**默认数据库**

MongoDB默认数据库为`test`，若尚未创建任何数据库，则集合/文档存储于`test`数据库中。

**删除数据库**

使用`db.dropDatabase()`方法删除数据库。

```
> use newdb
switched to db newdb
>db.dropDatabase()
>{ "dropped" : "newdb", "ok" : 1 }
```

### 2.2 集合操作

**创建集合**

使用`db.createCollection(name, options)`方法创建集合；其中`name`为要创建集合的名称（String）；`option`为用于指定集合配置的文档（Decument），可使用的选项列表如下：

| 字段        | 类型    | 描述                                                         |
| ----------- | ------- | ------------------------------------------------------------ |
| capped      | Boolean | (可选)如果为`true`，则启用封闭的集合。上限集合是固定大小的集合，它在达到其最大大小时自动覆盖其最旧的条目。 如果指定`true`，则还需要指定`size`参数。 |
| autoIndexId | Boolean | (可选)如果为`true`，则在`_id`字段上自动创建索引。默认值为`false`。 |
| size        | 数字    | (可选)指定上限集合的最大大小(以字节为单位)。 如果`capped`为`true`，那么还需要指定此字段的值。 |
| max         | 数字    | (可选)指定上限集合中允许的最大文档数。                       |

在向集合插入文档时，MongoDB先检查`capped`字段的大小，然后检查`max`字段。

示例1：创建无选项的集合。

```
>use test
switched to db test
>db.createCollection("mycol")
{ "ok" : 1 }
>
```

示例2：创建带有选项的集合。

```
>db.createCollection("mycolWithOption", {capped: true, autoIndexId: true, size: 6142800, max: 10000})
```

**插入文档时创建集合**

使用`db.COLLECTION_NAME.insert(...)`在插入值前自动创建名为`COLLECTION_NAME`的集合。

示例：

```
>db.user.insert({"name":"okarin"})
```

**查看集合列表**

使用`show collections`查看集合。

```
>show collections
mycol
mycolWithOption
```

**删除集合**

使用`db.COLLECTION_NAME.drop()`从数据库中删除名为`COLLECTION_NAME`的集合。

示例：

```
>db.mycolWithOption.drop()
true
```

### 2.3 数据类型

MongoDB支持的数据类型如下：

| 数据类型     | 说明                                           |
| ------------ | ---------------------------------------------- |
| 字符串       | 以UTF-8编码的字符串。                          |
| 整型         | 32或64位，取决于服务器。                       |
| 布尔类型     | true或false。                                  |
| 双精度浮点数 | 浮点值。                                       |
| 最大/最小键  | 此类型用于将值与最小和最大`BSON`元素进行比较。 |
| 数组         | 存储多个值到一个键中。                         |
| 时间戳       | 记录文档修改或新建。                           |
| 对象         | 用于嵌入式文档。                               |
| Null         | 存储Null值。                                   |
| 符号         | 同字符串；通常保留用于使用特定符号类型的语言。 |
| 日期         | 以Unix时间格式存储当前日期或时间。             |
| 对象ID       | 存储文档ID。                                   |
| 二进制数据   | 存储二进制数据。                               |
| 代码         | 存储Js代码。                                   |
| 正则表达式   | 存储正则表达式。                               |

### 2.4 文档操作

#### 插入文档

> insert()方法将被弃用。

**插入一个文档**

使用`db.COLLECTION_NAME.insertOne(document)`将一条数据插入到集合中。

示例：

```
use blog
db.createCollection('users')
db.users.insertOne({
    _id: 100,
    title: 'MongoDB Overview',
    description: 'MongoDB is no sql database',
    by: 'okarin',
    url: 'http://www.okarin-lab.com',
    tags: ['mongodb', 'database', 'NoSQL'],
    likes: 100
})
```

注：如果不指定`_id`参数，则MongoDB将为文档分配唯一的`ObjectId`（十二字节十六进制数）。如果指定ID，则以`save()`方法的形式替换包含`_id`文档的全部数据。

**插入多个文档**

使用`db.COLLECTION_NAME.insertMany(documents...)`将多条数据插入集合中。

示例：

```
db.something.insertMany([
   {
        item: "journal",
        qty: 25,
        tags: ["blank", "red"],
        size: { h: 14, w: 21, uom: "cm" }
   },
   {
        item: "mat",
        qty: 85,
        tags: ["gray"],
        size: { h: 27.9, w: 35.5, uom: "cm" }
   },
   {
        item: "mousepad",
        qty: 25,
        tags: ["gel", "blue"],
        size: { h: 19, w: 22.85, uom: "cm" }
   }
])
```

#### 查询文档

**查询全部文档**

`find()`方法查询当前集合的多个文档，基本语法：

```
>db.COLLECTION_NAME.find()
```

**查询指定列**
示例：

```
>db.something.find({}, {_id:1, item:1})
```

**查询单个文档**

`findOne()`方法查询当前集合的一个文档，基本语法：

```
>db.COLLECTION_NAME.findOne()
```

**where子句**

| 操作     | 语法                  | 示例                               | RDBMS等效语句       |
| -------- | --------------------- | ---------------------------------- | ------------------- |
| 相等     | {<key>:<value>}       | db.mycol.find({"by":"okarin"})     | where by = 'okarin' |
| 小于     | {<key>:{$lt:<value>}  | db.mycol.find({"likes":{$lt:50}})  | where likes < 50    |
| 小于等于 | {<key>:{$lte:<value>} | db.mycol.find({"likes":{$lte:50}}) | where likes <= 50   |
| 大于     | {<key>:{$gt:<value>}  | db.mycol.find({"likes":{$gt:50}})  | where likes > 50    |
| 大于等于 | {<key>:{$gte:<value>} | db.mycol.find({"likes":{$gte:50}}) | where likes >= 50   |
| 不等于   | {<key>:{$ne:<value>}  | db.mycol.find({"likes":{$ne:50}})  | where likes != 50   |

**AND操作符**

语法：

```
>db.COLLECTION_NAME.find({
	$and: [{key1: value1}, {key2: value2}]
})
```

示例：

```
db.something.find({
    $and: [{item: "mousepad"}, {qty: 25}]
})

# 或者

db.something.find({item: "mousepad", qty: 25})
```

等效SQL：

```sql
SELECT * FROM something WHERE item = 'mousepad' AND qty = 25
```

**OR操作符**

语法：

```
>db.COLLECTION_NAME.find({
	$or: [{key1: value1}, {key2: value2}]
})
```

示例：

```
db.something.find({
    $or: [{item: "none"}, {qty: 85}]
})
```

等效SQL：

```sql
SELECT * FROM something WHEAR item = 'none' OR qty = 85
```

**AND与OR联用**

示例：

```
>db.something.find({qty:{$gt:50}, $or:[{item: "mat"}, {tags: ['cotton']}]})
```

#### 查询嵌入/嵌套文档

准备测试数据：

```
>use blog;
>db.inventory.insertMany( [
   { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "A" },
   { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
   { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
   { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" }
]);
```

**匹配嵌入/嵌套文档**

示例：

```
>db.inventory.find({size: {h: 14, w: 21, uom: "cm"}})
```

注：嵌入文档的字段顺序必须一致。

**查询嵌套字段**

示例：

```
>db.inventory.find({"size.h": {$lt: 15}})
```

**指定AND条件**

示例：

```
>db.inventory.find({"size.h": {$lt: 15}, "size.uom": "in", status: "D"})
```

#### 更新文档

> `save()`方法已被弃用。

`updateOne()`和`updateMany()`以及`replaceOne()`方法用于将集合中的文档更新。

- `updateOne()`：更新单个现有文档中的值。
- `updateMany()`：更新多个现有文档中的值。
- `replaceOne()`：替换现有文档。

**更新字段的值（单个文档）**

示例：

```
>db.users.updateOne({title: 'MongoDB Overview'}, {$set:{title: 'New Update MongoDB Overview'}})
```

**更新字段的值（多个文档）**

示例：

```
>db.users.updateMany({likes: {$gt: 50}}, {$set:{title: 'updateMany() method.'}})
```

**替换现有文档**

示例：替换`_id`为200的数据。

```
>db.users.replaceOne({_id: 100} , {by: 'okarin-lab', description: 'Just for test.', likes: 200, tags: ['none'], title: 'save() method', url: 'www.lab.com'})
```

#### 删除文档

> `remove()`方法已弃用。

`deleteOne()`和`deleteMany()`用于删除集合中的文档；前者删除一个，后者删除多个。

**删除一条文档**

示例：

```
>db.users.deleteOne({_id: 100})
```

**删除多条文档**

示例：

```
>db.users.deleteMany({_id: {$lt: 200}})
```

**删除所有文档**

示例：

```
>db.users.deleteMany()
```



## 三、其他操作

### 3.1 选择字段

`findOne()`和`findMany()`方法默认显示文档的所有字段，为限制显示的字段，可设置第二个可选参数指定要检索的字段列表：设置字段列表对应的值为0（隐藏字段）或1（显示字段）。

语法：

```
>db.COLLECTION_NAME.find({}, {KEY1: 1, KEY2: 1})
```

示例：

```
>db.users.find({}, {_id: 1, by: 1})
```

### 3.2 限制记录数

`limit()`方法通过指定参数`number`的值来限制返回的记录数。

示例：

```
>db.users.find({}, {_id: 1, by: 1}).limit(2)
```

`skip()`方法通过指定参数`number`的值来过滤返回的记录。

示例：只显示第三个结果（跳过前两个且限制只显示一个结果）。

```
>db.users.find({}, {_id: 1, by: 1}).limit(1).skip(2)
```

### 3.3 排序记录

使用`sort()`方法对记录进行排序，参数-1指定降序，参数1指定升序。

语法：

```
>db.COLLECTION_NAME.find().sort({KEY:1})
```

示例：

```
>db.users.find({}, {_id: 1, title: 1}).sort({title: -1})
```

### 3.4 索引

索引有效提高查询效率：作为特殊的数据结构，以易于遍历的形式存储数据集的一小部分（特定字段或一组字段的值，按照索引中指定的字段值排序）。

#### 创建索引

使用`ensureIndex()`方法创建索引。

语法：

```
>db.COLLECTION_NAME.ensureIndex({KEY:1})
```

示例：基于`title`字段和`description`字段创建索引，前者升序，后者降序。

```
>db.users.ensureIndex({title:1, description:-1})
```

`ensureIndex()`选项列表：

| 参数               | 类型     | 描述                                                         |
| ------------------ | -------- | ------------------------------------------------------------ |
| background         | Boolean  | 后台构建索引，不影响数据库其他活动。                         |
| unique             | Boolean  | 创建唯一的索引。                                             |
| name               | String   | 索引名称；默认通过字段名和排序顺序生成。                     |
| dropDups           | Boolean  | 在可能有重复的字段上创建唯一索引。                           |
| sparse             | Boolean  | 仅引用具有指定字段的文档。                                   |
| expireAfterSeconds | integer  | 指定值（单位：秒）作为TTL，控制集合中保留文档的时间          |
| v                  | 索引版本 | 索引版本号。                                                 |
| weights            | 文档     | 权重范围：1~99999，表示该字段相对于其他索引字段在分数方面的意义。 |
| default_language   | String   | 对于文本索引，确定停止词列表的语言以及句柄和分词器的规则。   |
| language_override  | String   | 对于文本索引，指定文档中包含覆盖默认语言的字段名称。         |

### 3.5 聚合

聚合操作处理数据记录并返回计算结果，相当于SQL中的`count(*)`和`group by`；MongoDB中使用`aggregate()`方法实现聚合功能。

语法：

```
>db.COLLECTION_NAME.aggregate(AGGREGATE_OPERATION)
```

构建测试数据集：

```
db.article.insertMany([
{
   _id: 100,
   title: 'MongoDB Overview',
   description: 'MongoDB is no sql database',
   by_user: 'Maxsu',
   url: 'http://www.yiibai.com',
   tags: ['mongodb', 'database', 'NoSQL'],
   likes: 100
},
{
   _id: 101,
   title: 'NoSQL Overview',
   description: 'No sql database is very fast',
   by_user: 'Maxsu',
   url: 'http://www.yiibai.com',
   tags: ['mongodb', 'database', 'NoSQL'],
   likes: 10
},
{
   _id: 102,
   title: 'Neo4j Overview',
   description: 'Neo4j is no sql database',
   by_user: 'Kuber',
   url: 'http://www.neo4j.com',
   tags: ['neo4j', 'database', 'NoSQL'],
   likes: 750
},
{
   _id: 103,
   title: 'MySQL Overview',
   description: 'MySQL is sql database',
   by_user: 'Curry',
   url: 'http://www.yiibai.com/mysql/',
   tags: ['MySQL', 'database', 'SQL'],
   likes: 350
}])
```

聚合操作：

```
>db.article.aggregate({$group:{_id: "$by_user", num_tutorial:{$sum:1}}})
```

等效SQL：

```
SELECT by_user, count(*) AS num_tutorial FROM `artical` GROUP BY by_user;
```

#### 聚合表达式列表

| 表达式    | 描述                                     | 示例                                                         |
| --------- | ---------------------------------------- | ------------------------------------------------------------ |
| $sum      | 求和                                     | `db.article.aggregate([{$group:{_id:"$by_user", num_tutorail: {$sum:"$likes"}}}])` |
| $avg      | 求平均值                                 | `db.article.aggregate([{$group:{_id:"$by_user", num_tutorail: {$avg:"$likes"}}}])` |
| $min      | 求最小值                                 | `db.article.aggregate([{$group:{_id:"$by_user", num_tutorail: {$min:"$likes"}}}])` |
| $max      | 求最大值                                 | `db.article.aggregate([{$group:{_id:"$by_user", num_tutorail: {$max:"$likes"}}}])` |
| $push     | 插入新值到生成的文档数组中               | `db.article.aggregate([{$group:{_id:"$by_user", url: {$push:"$url"}}}])` |
| $addToSet | 插入新值到生成的文档数组中，不创建重复项 | `db.article.aggregate([{$group:{_id:"$by_user", url: {$addToSet:"$url"}}}])` |
| $first    | 根据分组从源文档获取第一个文档           | `db.article.aggregate([{$group:{_id:"$by_user", first_url: {$first:"$url"}}}])` |
| $last     | 根据分组从源文档获取最后一个文档         | `db.article.aggregate([{$group:{_id:"$by_user", last_url: {$last:"$url"}}}])` |

#### 管道

集合输出作为管道输入，进行再处理；阶段操作符如下：

- `$project`：选择特定字段。
- `$match`：匹配条件。
- `$group`：分组。
- `$sort`：排序。
- `$skip`：跳过部分结果。
- `$limit`：限制个数。
- `$unwind`：展开正在使用数组的文档。

**示例：分组并排序。**

```
>db.article.aggregate([{$group:{_id:"$by_user", num_tutorail: {$sum:"$likes"}}}, {$sort:{num_tutorail:-1}}])
```

## 四、MongoDB相关

### 4.1 复制集

#### 复制集简介

一组MongoDB复制集，即一组MongoDB进程，这些进程维护同一个数据集合；复制集提供了数据冗余和高等级的可靠性。

#### 复制集的优势

- 规避单点损坏带来的风险，提高容灾能力。
- 读写分离，提高系统负载。
- 提供数据冗余，提高可用性（维护无停机）。

#### 复制集的基本架构

> 早期版本包括主节点、从节点和仲裁节点（不存储数据，只参与投票，当前版本已不建议）。

##### 通过选举完成故障恢复

- 具有投票权的节点之间两两互相发送心跳。
- 当连续5次心跳未收到时判定为节点失联。
- 若主节点失联，则由具有投票权的从节点在从节点中选举出新的主节点。
- 若从节点失联，则不会产生选举。
- 选举基于`RAFT`一致性算法实现，选举成功的必要条件是大多数投票节点存活。
- 复制集中最多可以有50个节点，但具有投票权的节点最多7个。
- 若原主节点从故障中恢复，将以从节点的身份加入到复制集。

#### 复制集搭建

搭建拥有3个成员的复制集。

步骤：

1. 创建节点目录：每个节点目录相同，包含数据目录`data`、日志目录`logs`、Js文件目录`js`、启动脚本`mongod-server.sh`，节点配置文件`mongod.conf`。

   ```bash
   $ mkdir instance_1
   $ mkdir instance_2
   $ mkdir instance_3
   ```

2. 节点配置文件：路径中的`instance`分别替换为`instance_1`、`instance_2`、`instance_3`；每个实例的端口分别配置为`27017`，`27018`，`27019`。

   ```
   # 日志配置
   logpath=/mnt/mongodb/instance/logs/mongod.log #日志目录
   logappend=true
   timeStampFormat=iso8601-utc
   
   # 存储配置
   dbpath=/mnt/mongodb/instance/data/db #数据存储目录
   directoryperdb=true
   
   # 进程管理配置
   fork=true
   pidfilepath=/mnt/mongodb/instance/var/run/mongod.pid # 进程PID文件保存目录
   
   # 网络配置
   bind_ip=127.0.0.1
   port=27017 
   
   # 安全配置
   #auth=true
   
   # 复制集
   replSet=test-set    # 复制集名称
   ```

3. 启动节点，脚本如下，参数为：`status`、`start`、`stop`。

   ```shell
   #!/bin/bash
   
   source /etc/profile
   source ~/.bash_profile
   
   # MongoDB 需要将 LC_ALL 设置为 C
   export LC_ALL=C
   
   # 判断是否开启 debug 模式。如果开启了 debug 模式，则通过 set -x 设置显示执行的命令
   [[ -n "$DEBUG" ]] && set -x
   
   #ANSI Colors
   echoRed() { echo $'\e[0;31m'"$1"$'\e[0m'; }
   echoGreen() { echo $'\e[0;32m'"$1"$'\e[0m'; }
   echoYellow() { echo $'\e[0;33m'"$1"$'\e[0m'; }
   
   mongodexe="/DATA/mongodb-3.4.5/rhel-3.4.5/bin/mongod"
   
   workdir=$(pwd)
   cd $(dirname $0)
   mongo_conf=$(pwd)/mongod.conf
   cd $workdir
   
   [[ -f "$mongo_conf" ]] && source $mongo_conf
   
   [[ -n "$dbpath" ]] || dbpath=/data/db
   [[ -n "$logpath" ]] || logpath=/var/logs/mongo.log
   [[ -n "$pidfilepath" ]] || pidfilepath=/var/run/mongo.pid
   
   # 检查权限
   checkPermission() {
     # 只有 test 命令可以和多种系统运算符一起使用
     test -d $dbpath -a -r $dbpath -a -w $dbpath  || { echoRed "[$dbpath] not found or permission denied"; return 1; }
     test -f $logpath -a -r $logpath -a -w $logpath  || { echoRed "[$logpath] not found or permission denied"; return 1; }
     test -f "$pidfilepath" -a -r $pidfilepath -a -w $pidfilepath || { echoRed "[$pidfilepath] not found or permission denied"; return 1; }
     return 0
   }
   
   # 并获取 PID
   pid=$(cat "$pidfilepath")
   
   
   # 检查进程是否存在
   isRunning() {
     ps -p "$1" &> /dev/null
   }
   
   
   status() {
     # 判断进程是否存在  
     isRunning "$pid" || { echoRed "Stoped [$pid]"; return 0; }
     echoGreen "Running [$pid]"
     return 0
   }
   
   do_start() {
     if [[ -f "$mongo_conf" ]]; then
       /bin/sh -c "$mongodexe -f $mongo_conf"
     else
       /bin/sh -c "$mongodexe"
     fi
   
     pid=$(cat "$pidfilepath")
   
     
     # 判断进程是否存在  
     isRunning "$pid" || { echoRed "Failed to start"; return 1; }
     echoGreen "Started [$pid]"
     return 0
   }
   
   start() {
     # 判断进程是否存在  
     isRunning "$pid" && { echoYellow " Running [$pid]"; return 1; }
     # 检查权限
     checkPermission || return 1
     # 启动
     do_start
   }
   
   do_stop() {
    if [[ -f "$mongo_conf" ]]; then
      /bin/sh -c "$mongodexe -f $mongo_conf --shutdown"
    else
      /bin/sh -c "$mongodexe --shutdown"
    fi
      
    # 判断进程是否存在  
    isRunning "$pid" && { echoRed "Unable to kill process $pid"; return 1; }
    echoGreen "Stopped [[$pid]]"
    return 0
   }
   
   stop() {
    # 判断进程是否存在  
    isRunning "$pid" || { echoYellow "Not running (process ${pid} not found)"; return 1;}
    do_stop
   }
   
   # Call the appropriate action function
   case "$1" in
   start)
     start ; exit $?;;
   stop)
     stop ; exit $?;;
   status)
     status ; exit $?;;
   restart)
     stop; start; exit $?;;
   *)
     echoRed "Usage: $0 {start|stop|status|restart}"; exit 1;
   esac
   
   exit 0
   ```

4. 初始化复制集，任意节点的Js目录，脚本`initiate_rs.js`如下：

   ```javascript
   rs.initiate()
   rs.add("hostname:27017")
   rs.add("hostname:27019")
   rs.conf()
   ```

   执行初始化：

   ```bash
   $ mongo --port 27018 ./js/initiate_rs.js
   ```

   连接到其中一个节点，查看初始化是否成功：

   ```
   >rs.status()
   ```

5. 连接主节点，插入数据。

6. 连接从节点，查看数据是否同步。

7. 默认情况下，从节点需执行`rs.slaveOk()`后才能读写。

> 检查实例是否连接到主服务器，使用`db.isMaster()`命令。

### 4.2 分片

分片是指在多台机器之间存储数据的过程，随着数据量的增加，单个机器不足以存储所有数据，也不能提供可接受的读写吞吐量。解决方法：水平扩展，使用分片，增加机器。

#### 分片的基本组件

**碎片**

碎片用于存储数据，每个碎片都是一个独立的复制集（生产环境下）。

**配置服务器**

配置服务器存储集群的元数据，包含集群的数据集与分片的映射，查询路由器通过使用此元数据将操作定位到特定分片。分片集群拥有3个配置服务器（生产环境下）。

**查询路由器**

查询路由器为mongo实例，负责通过配置服务器将操作定位到碎片，然后将结果返回给客户端；分片集群使用多个查询路由器分割客户端请求负载。

### 4.3 备份与恢复

#### 备份

MongoDB使用`mongodump`命令创建数据库备份：导出转储服务器的整个数据到转储目录。

示例：在mongod服务器上转到mongodb实例的`bin`目录下，执行以下命令备份。

```
>mongodump
```

默认情况下，将在执行目录下生成`dump`目录，并为每个数据库创建子目录后存储备份。备份文件的目录可在`/etc/mongod.conf`中设定：

| 语法                                           | 描述                       |
| ---------------------------------------------- | -------------------------- |
| mongodump --host HOST_NAME —port PORT_NUMBER   | 备份指定实例的所有数据库。 |
| mongodump --out BACKUP_DIRECTORY               | 在指定路径上备份数据库。   |
| mongodump --collection COLLECTION --db DB_NAME | 备份指定数据库的指定集合。 |

#### 恢复

MongoDB使用`mongorestore`命令从备份目录中恢复所有数据。

### 4.4 部署

监测实例状态：`mongostat`

跟踪读写活动：`mongotp`

## 五、高端操作

待续。

## 六、用户与安全

待续。























































