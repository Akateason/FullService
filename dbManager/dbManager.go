/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-03 23:27:12
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-04 00:02:24
 * @FilePath: /FullService/dbManager/dbManager.go
 * @Description: ORM Database Manager
 *
 * Copyright (c) 2023 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package dbManager

import (
	"fmt"
	"time"

	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func SetupDatabase() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	fmt.Println(err)

	err = engine.Sync2(new(User))

	user := new(User)
	user.Name = "myname"
	user.Id = 1
	affected, err := engine.Insert(user)
	// INSERT INTO user (name) values (?)
	fmt.Println(affected)
	fmt.Println(err)

	users := make([]User, 0)
	err = engine.Find(&users)
	fmt.Println(err)
	fmt.Println(users)
}

type User struct {
	Id      int64
	Name    string
	Sex     string
	Age     int
	Passwd  string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
