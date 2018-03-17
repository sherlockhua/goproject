package main

import (
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"time" 
	//"math/rand"
)

type UserInfo struct {
	UserId int         `db:"user_id"`
	Username string    `db:"username"`
	Sex string	`db:"sex"`
	Email string  `db:"email"`
}

func main() {
	Db, err := sqlx.Open("mysql", "root:@tcp(10.0.0.200:3306)/godb")
	if err != nil {
		fmt.Println("connect to msyql failed, ", err)
		return
	}

	startTime := time.Now().UnixNano()
	
	conn, _ := Db.Begin()
	 
	_, err = conn.Exec("insert into user_info(username, sex, email)values(?,?,?)",
	"user01", "man", "email")
	if err != nil {
		conn.Rollback()
		fmt.Println("insert failed, err:", err)
		return
	}

	_, err = conn.Exec("insert into user_info(username, sex, email)values(?,?,?)",
	"user01", "man", "email")
	if err != nil {
		conn.Rollback()
		fmt.Println("insert failed, err:", err)
		return
	}

	conn.Commit()
	
	var userInfo UserInfo
	err = Db.Get(&userInfo, "select user_id, username, sex, email from user_info where username=?", "user1674879938132494608")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("insert succ cost:%d ms\n", (endTime - startTime)/1000/1000)

	Db.Close()
}