package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {

	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database

}

func main() {
	defer Db.Close() // 注意这行代码要写在上面err判断的下面
	res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)

}
