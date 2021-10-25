## 一、Go语言操作MySQL（database/sql）

### 1.1 环境搭建

安装Docker：https://www.docker.com/products/docker-desktop

```bash
// Windows 启动Docker出错的解决办法
$ cd "C:\Program Files\Docker\Docker"
$ .\DockerCli.exe -SwitchDaemon
```

安装MySQL：

```bash
$ docker pull mysql:latest
```

查看Docker本地镜像：

```bash
$ docker images
```

运行MySQL：容器的3306端口映射到本机13306端口。

```bash
$ docker run --name mysql -p 13306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```

> Win7首次连接必须在Docker启动一个MySQL客户端连接MySQL服务，后续才能使用Navicat连接。
>
> ```bash
> docker run -it --network host --rm mysql mysql -h127.0.0.1 -P13306 --default-character-set=utf8mb4 -uroot -p
> ```

### 1.2 连接

Go语言中的`database/sql`包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。

#### 1.2.1 下载依赖

```bash
$ go get -u github.com/go-sql-driver/mysql
```

#### 1.2.2 使用mysql驱动

```go
func Open(driverName, dataSourceName string) (*DB, error)
```

`Open`函数打开一个`driverName`指定的数据库，通过`dataSourceName`指定数据源（连接信息）。

```go
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // 匿名：执行init()
)

var db *sql.DB
var err error

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("close mysql failed!")
		}
	}(db)

}
```

#### 1.2.3 初始化连接

使用`Ping`方法检验数据源的名称是否有效。

```go
var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:13306)/sql_test?charset=utf8mb4&parseTime=True"
	// 初始化全局db对象
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	// 连接设置
	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200) // 最大连接数
	db.SetMaxIdleConns(10)  // 最大空闲连接数

	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("close mysql failed!")
		}
	}(db)
	queryRowDemo()
}
```

其中`sql.DB`表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。在内部维护着一个具有零到多个底层连接的连接池，并发安全。 

#### 1.2.4 SetMaxOpenConns

```go
func (db *DB) SetMaxOpenConns(n int)
```

`SetMaxOpenConns`设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。 如果n<=0，不会限制最大开启连接数，默认为0（无限制）。 

#### 1.2.5 SetMaxIdleConns

```go
func (db *DB) SetMaxIdleConns(n int)
```

SetMaxIdleConns设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。

### 1.3 CRUD

#### 1.3.1 建库建表

新建数据库：

```mysql
CREATE DATABASE sql_test;
```

进入数据库：

```mysql
USE sql_test;
```

建表：

```mysql
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

#### 1.3.2 查询数据

定义结构体存储user表的数据：

```go
type user struct {
	id   int
	age  int
	name string
}
```

**单行查询**

执行一次单行查询，最多返回一行结果。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果） 

```go
func (db *DB) QueryRow(query string, args ...interface{}) *Row
```

示例：

```go
// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
```

**多行查询**

多行查询`db.Query()`执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。

```go
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```

示例：

```go
// 查询多行数据
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic("close connection failed!")
		}
	}(rows)

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}
```

#### 1.3.3 插入数据

插入、更新和删除操作都使用`Exec`方法。

```go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
```

Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。 

示例：

```go
// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "lettredamour", 25)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	lastInsertId, err := ret.LastInsertId() // 最新插入数据的ID
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert success, the lastinsert id is %d.\n", lastInsertId)
}
```

#### 1.3.4 更新数据

示例：

```go
// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, "24", 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	rowsAffected, err := ret.RowsAffected() // 操作影响的行
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", rowsAffected)
}
```

#### 1.3.5 删除数据

示例：

```go
// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	rowsAffected, err := ret.RowsAffected() // 影响操作的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", rowsAffected)
}
```

### 1.4 MySQL预处理

#### 1.4.1 什么是预处理

普通SQL语句执行过程：

1. 客户端对SQL语句进行占位符替换得到完整的SQL语句。
2. 客户端发送完整SQL语句到MySQL服务端
3. MySQL服务端执行完整的SQL语句并将结果返回给客户端。

预处理执行过程：

1. 把SQL语句分成两部分，命令部分与数据部分。
2. 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
3. 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
4. MySQL服务端执行完整的SQL语句并将结果返回给客户端。

#### 1.4.2 预处理的优势

- 优化MySQL服务器重复执行SQL的方法，挺高服务器性能：使服务器提前编译，一次编译多次执行。
- 防止SQL注入。

#### 1.4.3 Go实现MySQL预处理

`database/sql`中使用下面的`Prepare`方法来实现预处理操作。

```go
func (db *DB) Prepare(query string) (*Stmt, error)
```

`Prepare`方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。 

查询操作预处理示例：

```go
// 预处理查询示例
func prepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}
```

插入、更新和删除操作的预处理相同，以插入操作为例：

```go
// 预处理插入操作
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	_, err = stmt.Exec("shinin", 24)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	_, err = stmt.Exec("shinrin", 23)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert success.")
}
```

#### 1.4.4 SQL注入

根据`name`查询`user`表示例：

```go
// SQL注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name = %s", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	fmt.Printf("user:%#v\n", u)
}
```

输入以下字符串将引发SQL注入：

```go
sqlInjectDemo("xxx' or 1=1#")
sqlInjectDemo("xxx' union select * from user #")
sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
```

不同的数据库中，SQL语句使用的占位符语法不尽相同：

|   数据库   |  占位符语法  |
| :--------: | :----------: |
|   MySQL    |     `?`      |
| PostgreSQL | `$1`, `$2`等 |
|   SQLite   |  `?` 和`$1`  |
|   Oracle   |   `:name`    |

### 1.5 MySQL事务

#### 1.5.1 什么是事务

事务：一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成。A转账给B，这里面就需要执行两次update操作。 

在MySQL中只有使用了`Innodb`数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。 

#### 1.5.2 事务的ACID

通常事务必须满足4个条件：原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）。 

|  条件  |                             解释                             |
| :----: | :----------------------------------------------------------: |
| 原子性 | 一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。 |
| 一致性 | 在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。 |
| 隔离性 | 数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。 |
| 持久性 |             事务处理结束后，对数据的修改是永久的             |

#### 1.5.3 事务相关方法

Go语言中使用以下三个方法实现MySQL中的事务操作。  

**开始事务**

```go
func (db *DB) Begin() (*Tx, error)
```

**提交事务**

```go
func (tx *Tx) Commit() error
```

**回滚事务**

```go
func (tx *Tx) Rollback() error
```

事务示例：

```go
// 事务示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			err := tx.Rollback() // 回滚
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}

	sqlStr1 := "update user set age = 25 where id = ?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	rowsAffected1, err := ret1.RowsAffected()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "update user set age = 24 where id = ?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	rowsAffected2, err := ret2.RowsAffected()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("exec ret2.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(rowsAffected1, rowsAffected2)
	if rowsAffected1 == 1 && rowsAffected2 == 1 {
		err := tx.Commit()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Transaction committed.")
	} else {
		err := tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Transaction rolled back.")
	}

	fmt.Println("exec trans success.")
}
```

## 二、sqlx库

### 2.1 sqlx简介

地址：https://github.com/jmoiron/sqlx

sqlx功能比database/sql更强大。

### 2.2 sqlx安装

指令：

```go
$ go get github.com/jmoiron/sqlx
```

> 项目中使用import后go mod tidy引入即可。

### 2.3 基本使用

#### 2.3.1 连接数据库

示例：

```go
var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:13306)/sql_test?charset=utf8mb4&parseTime=True"
	// 或可使用MustConnect：连接失败即panic
	db, err = sqlx.Connect("mysql", dsn) // Open() + Ping()
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}
```

#### 2.3.2 查询数据

定义user结构体：

```go
// 使用反射映射要求字段首字母大写。
type user struct {
	ID   int    `db:"id"`
	Name string `db:"name""`
	Age  int    `db:"age"`
}
```

查询单行数据：

```go
// 查询单行数据
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id = ?"
	var u user
	err := db.Get(&u, sqlStr, 1) // 自动映射到结构体
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d", u.ID, u.Name, u.Age)
}
```

查询多行数据：

```go
// 查询多行数据
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
```

#### 2.3.3 插入数据

示例：

```go
// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?,?)"
	ret, err := db.Exec(sqlStr, "love letter", 22)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	insertId, err := ret.LastInsertId() // 新插入数据的ID
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", insertId)
}
```

#### 2.3.4 更新数据

示例：

```go
// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age = ? where name = ?"
	ret, err := db.Exec(sqlStr, 23, "love letter")
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get ret.RowsAffected() failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, the ret.RowsAffected() is %d.\n", affected)
}
```

#### 2.3.5 删除数据

示例：

```go
// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where name = ?"
	ret, err := db.Exec(sqlStr, "love letter")
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get ret.RowsAffected() failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, the ret.RowsAffected() is %d.\n", affected)
}
```

#### 2.3.6 NamedExec

`DB.NamedExec`方法用来绑定SQL语句与结构体或map中的同名字段。 

示例：

```go
// NamedExec
func namedInsert() (err error) {
	sqlStr := "INSERT INTO user (name, age) VALUES (:name, :age)"
	_, err = db.NamedExec(sqlStr, map[string]interface{}{
		"name": "love letter",
		"age":  23,
	})
	if err != nil {
		return err
	}
	return
}
```

#### 2.3.7 NamedQuery

示例：

```go
// NamedQuery
func namedQuery() {
	// 使用map命名查询
	sqlStr := "SELECT * FROM user WHERE name = :name"
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name":"love letter"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return 
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
	
	// 使用结构体命名查询
	u := user{Name: "love letter"}
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}
```

#### 2.3.8 事务操作

`sqlx`提供的`db.Beginx()`和`tx.Exec()`方法以支持事务操作。

```go
// 事务操作
func transactionDemo() (err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err := tx.Rollback()
			if err != nil {
				panic(err)
			}
			panic(p)
		} else if err != nil {
			fmt.Printf("rollback")
			err := tx.Rollback()
			if err != nil {
				panic(err)
			}
		} else {
			err := tx.Commit()
			if err != nil {
				panic(err)
			}
			fmt.Printf("commit")
		}
	}()

	sqlStr1 := "UPDATE user SET age = 21 WHERE id = ?"
	ret1, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	affected, err := ret1.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "UPDATE user SET age = 23 WHERE id = ?"
	ret2, err := tx.Exec(sqlStr2, 7)
	if err != nil {
		return err
	}
	affected, err = ret2.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return
}// 事务操作
func transactionDemo() (err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err := tx.Rollback()
			if err != nil {
				panic(err)
			}
			panic(p)
		} else if err != nil {
			fmt.Printf("rollback")
			err := tx.Rollback()
			if err != nil {
				panic(err)
			}
		} else {
			err := tx.Commit()
			if err != nil {
				panic(err)
			}
			fmt.Printf("commit")
		}
	}()

	sqlStr1 := "UPDATE user SET age = 21 WHERE id = ?"
	ret1, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	affected, err := ret1.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "UPDATE user SET age = 23 WHERE id = ?"
	ret2, err := tx.Exec(sqlStr2, 7)
	if err != nil {
		return err
	}
	affected, err = ret2.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return
}
```

### 2.4 sqlx.In

#### 2.4.1 sqlx.In的批量导入

**bindvars（绑定变量）**

查询占位符`?`在内部称为**bindvars（查询占位符）**，使用查询占位符可以有效防止SQL注入攻击。不同的数据库有不同的`bindvars`：

- MySQL中使用`?`
- PostgreSQL使用枚举的`$1`、`$2`等bindvar语法。
- SQLite中`?`和`$1`的语法都支持。
- Oracle中使用`:name`的语法。

注：`bindvars`仅用于参数化，不允许更改SQL语句的结构，即无法对列名或表名生效。

**使用sqlx.In实现批量插入**

前提：结构体实现`driver.Valuer`接口。

```go
// Value 批量插入，必先实现driver.Valuer接口
func (u user) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}
```

批量插入：

```go
// BatchInsertUsers1 使用sqlx.In实现批量插入
func BatchInsertUsers(users []interface{}) (err error) {
	query, args, err := sqlx.In("INSERT INTO user (name, age) VALUES (?), (?), (?)", users...)
	if err != nil {
		return err
	}
	fmt.Println(query)
	fmt.Println(args)
	_, err = db.Exec(query, args...)
	return
}
```

**使用NamedExex实现批量插入**

示例：

```go
// BatchInsertUsers2 使用NamedExec实现批量插入
func BatchInsertUsers2(users []interface{}) (err error) {
	_, err = db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return
}
```

#### 2.4.2 sqlx.In查询

**in查询**

查询id在给定id集合中的数据。

```go
// QueryByIDs in查询
func QueryByIDs(ids []int) (users []user, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In返回带`?`bindvar的查询语句，使用Rebind重新绑定。
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}
```

**in查询和FIND_IN_SET函数**

查询id在给定id集合的数据并维持给定id集合的顺序。 

```go
// QueryAndOrderByIDs in查询和FIND_IN_SET函数
func QueryAndOrderByIDs(ids []int) (users []user, err error) {
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}
```

