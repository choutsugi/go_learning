package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 使用反射映射要求字段首字母大写。
type user struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var db *sqlx.DB

// 连接数据库
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

// 查询单行数据
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id = ?"
	var u user
	err := db.Get(&u, sqlStr, 1) // 自动映射到结构体
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

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

// NamedQuery
func namedQuery() (err error) {
	// 使用map命名查询
	sqlStr := "SELECT * FROM user WHERE name = :name"
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "love letter"})
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
	return
}

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
}

// Value 批量插入，必先实现driver.Valuer接口
func (u user) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers1 使用sqlx.In实现批量插入
func BatchInsertUsers1(users []interface{}) (err error) {
	query, args, err := sqlx.In("INSERT INTO user (name, age) VALUES (?), (?), (?)", users...)
	if err != nil {
		return err
	}
	fmt.Println(query)
	fmt.Println(args)
	_, err = db.Exec(query, args...)
	return
}

// BatchInsertUsers2 使用NamedExec实现批量插入
func BatchInsertUsers2(users []interface{}) (err error) {
	_, err = db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return
}

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

func main() {
	if err := initDB(); err != nil {
		panic(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	queryRowDemo()
	queryMultiRowDemo()
	insertRowDemo()
	updateRowDemo()
	deleteRowDemo()

	_ = namedInsert()
	_ = namedQuery()

	_ = transactionDemo()

	u1 := user{
		Name: "Teemo",
		Age:  6,
	}

	u2 := user{
		Name: "Yasuo",
		Age:  27,
	}

	u3 := user{
		Name: "ZOE",
		Age:  9999,
	}

	users := []interface{}{u1, u2, u3}
	//err := BatchInsertUsers1(users)
	err := BatchInsertUsers2(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
		return
	}

	ids := []int{2, 3, 1}
	ret, err := QueryByIDs(ids)
	if err != nil {
		fmt.Printf("QueryByIDs failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", ret)

	ret, err = QueryAndOrderByIDs(ids)
	if err != nil {
		fmt.Printf("QueryAndOrderByIDs failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", ret)
}
