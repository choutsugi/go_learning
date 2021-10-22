package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 匿名：执行init()
)

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

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

	// 增删改查
	queryRowDemo()
	queryMultiRowDemo()
	insertRowDemo()
	updateRowDemo()
	deleteRowDemo()

	// 预处理
	prepareQueryDemo()
	prepareInsertDemo()

	// 输入以下字符串将引发SQL注入
	//sqlInjectDemo("xxx' or 1=1#")
	//sqlInjectDemo("xxx' union select * from user #")
	//sqlInjectDemo("xxx' and (select count(*) from user) <10 #")

	// 事务
	transactionDemo()
}
