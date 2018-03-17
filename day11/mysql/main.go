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
	/*
	fmt.Println("connect to mysql succ")
	for i := 0; i < 10000000; i++ {
		username := fmt.Sprintf("user%d", rand.Int63())
		email := fmt.Sprintf("%d@qq.com", rand.Int())
		sex := "男"
		if i % 2 == 0 {
			sex = "女"
		}
		_, err := Db.Exec("insert into user_info(username, sex, email)values(?,?,?)",
		username, sex, email)
		if err != nil {
			fmt.Println("insert failed, err:", err)
			return
		}
		//user_id, _ := result.LastInsertId()
		//fmt.Println("insert succ, user_id:", user_id)
	}
	*/
	
	var userInfo UserInfo
	err = Db.Get(&userInfo, "select user_id, username, sex, email from user_info where username=?", "user1674879938132494608")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("insert succ cost:%d ms\n", (endTime - startTime)/1000/1000)
/*
	var userInfoList []*UserInfo
	err = Db.Select(&userInfoList, "select user_id, username, sex, email from user_info where user_id>?", 2)
	
	fmt.Printf("user_info_list:%#v\n", userInfoList)
	if err != nil {
		fmt.Println("select failed, err:", err)
		return
	}*/
	Db.Close()
}