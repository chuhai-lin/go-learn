package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Database *sqlx.DB

type Roles struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	DefaultPath string `db:"default_path"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func init() {
	user := "XXX"
	password := "XXX"
	host := "XXX"
	port := 3306
	db := "XXX"

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, db)

	d, err := sqlx.Open("mysql", url)
	if err != nil {
		fmt.Println("open error: ", err)
		return
	}
	Database = d
}

func main() {
	defer Database.Close()

	var roles []Roles
	err := Database.Select(&roles, "select * from roles where name=?", "局长")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(roles)
}
