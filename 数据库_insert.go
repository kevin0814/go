package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Person struct {
	UserId     int     `db:"user_id"`
	UserName   string  `db:"username"`
	Sex        string  `db:"sex"`
	Email      string  `db:"email"`
}

type Place struct {
    Country string `db:"country"`
    City    string `db:"city"`
    TelCode string `db:"tel_code"`
}
var Db *sqlx.DB
func init(){
	datebase ,err :=sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err !=nil {
		fmt.Println("open mysql failed",err)
	}else {
		Db = datebase
		//创建person 表
		CreateTableWithUser()
		////创建文章表
		CreateTableWithPlace()
		//defer Db.Close()
	}


}

func main() {
	//result,err := ModifyDB("insert into person(username, sex, email)values(?, ?, ?)",
	//	"stu001", "man", "stu01@qq.com")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//fmt.Println("insert succ:",result)
	r, err := Db.Exec("insert into person(username, sex, email) values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := Db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
func CreateTableWithUser(){
	sql := `CREATE TABLE IF NOT EXISTS person(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		user_id INT(64),
		username VARCHAR(32),
		sex VARCHAR(10),
		email VARCHAR(10)
		);`
	ModifyDB(sql)
}

func CreateTableWithPlace(){
	sql := `CREATE TABLE IF NOT EXISTS place(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		country VARCHAR(64),
		city VARCHAR(32),
		tel_code VARCHAR(32)
		);`
	ModifyDB(sql)
}
