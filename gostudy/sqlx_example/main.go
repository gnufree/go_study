package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func initDb() error {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang"
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

func testSqlxQuery() {
	sqlstr := "select id,name,age from user where id=?"
	var user User

	err := DB.Get(&user, sqlstr, 3)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}

	fmt.Printf("user:%#v\n", user)
}

func testSqlxQueryMul() {
	sqlstr := "select id,name,age from user where id>?"
	var user []User

	err := DB.Select(&user, sqlstr, 0)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}

	fmt.Printf("user:%#v\n", user)
}

func testSqlxInsert() {
	sqlstr := "insert into user (name,age) values(?,?)"
	// var user []User

	result, err := DB.Exec(sqlstr, "mark", 31)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	n, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("lastinsertid failed, err:%v\n", err)
	}

	fmt.Printf("id :%d\n", n)
}

func queryDB(name string) {
	sqlstr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
	fmt.Printf("sql: %s\n", sqlstr)
	var user []User

	err := DB.Select(&user, sqlstr)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
	}
	for _, v := range user {
		fmt.Printf("user:%#v\n", v)
	}

}

func testSqlInject() {
	// queryDB("abc' or 1 = 1 #")
	// queryDB("name=abc' and (select count(*) from user ) > 0#")
	queryDB("name=123' union select *from user #")
	// queryDBBySqlx("name=123' union select *from user #")
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	// testSqlxQuery()
	// testSqlxQueryMul()
	// testSqlxInsert()
	testSqlInject()
}
