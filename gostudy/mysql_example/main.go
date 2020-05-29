package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDb() error {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int    `db:id`
	Name string `db:name`
	Age  int    `db:age`
}

func testQueryMultilRow() {
	sqlstr := "select id,name,age from user where id > ?"
	rows, err := DB.Query(sqlstr, 0)
	// 重点关注，rows一定要关闭
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

func testQueryData() {
	for i := 0; i < 101; i++ {
		fmt.Printf("query %d times\n", i)
		sqlstr := "select id,name,age from user where id=?"
		row := DB.QueryRow(sqlstr, 3)
		/*
			if row != nil {
				continue
			}*/
		var user User
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d \n", user.Id, user.Name, user.Age)
	}

	// fmt.Println(user.Id, user.Name, user.Age)
}

func testInsertData() {
	sqlstr := "insert into user (name,age) values(?,?)"
	result, err := DB.Exec(sqlstr, "tom", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	count, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert succ, %d\n", count)

}

func testUpdateData() {
	sqlstr := "update user set name=? where id=?"
	result, err := DB.Exec(sqlstr, "jim", 4)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("update  failed, err:%v\n", err)
		return
	}
	fmt.Printf("update succ, %d\n", count)

}

func testDeleteData() {
	sqlstr := "delete from  user where id=?"
	result, err := DB.Exec(sqlstr, 4)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("delete affect  failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete succ, %d\n", count)

}

func testPrepareQuery() {
	sqlstr := "select id,name,age from user where id > ?"
	stmt, err := DB.Prepare(sqlstr)
	// 重点关注，rows一定要关闭
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	rows, err := stmt.Query(0)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

func testPrepareInsertData() {
	sqlstr := "insert into user (name,age) values(?,?)"
	stmt, err := DB.Prepare(sqlstr)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("insert prepare, err:%v\n", err)
		return
	}
	result, err := stmt.Exec("zhangsan", 20)
	if err != nil {
		fmt.Printf("insert data failed, err:%v\n", err)
	}
	count, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert succ, %d\n", count)

}

func testTrans() {
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}
	sqlstr := "update user set age=1 where id=?"
	_, err = conn.Exec(sqlstr, 1)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec first sql:%s  failed, err:%v\n", sqlstr, err)
		return
	}
	sqlstr = "update user set age=2; where id=?"
	_, err = conn.Exec(sqlstr, 2)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sencod sql:%s  failed, err:%v\n", sqlstr, err)
		return
	}
	err = conn.Commit()
	if err != nil {
		fmt.Printf("commit failed ,err:%v\n", err)
		conn.Rollback()
		return
	}
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
	}
	// testQueryData()
	// testQueryMultilRow()
	// testInsertData()
	// testUpdateData()
	// testDeleteData()
	// testPrepareQuery()
	// testPrepareInsertData()
	testTrans()
}
