/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-03 23:27:12
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-04 01:46:15
 * @FilePath: /FullService/dbManager/dbManager.go
 * @Description: ORM Database Manager
 *
 * Copyright (c) 2023 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package dbManager

import (
	"FullService/person"
	"fmt"

	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func SetupDatabase() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
	}

	err = engine.Sync2(new(person.Person))

	user1 := new(person.Person)
	user1.Name = "myname"
	user1.Id = 2
	affected, err := engine.Insert(user1)
	// INSERT INTO user (name) values (?)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)

	users := make([]person.Person, 0)
	err = engine.Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	fmt.Println("select all")
	fmt.Println(users)
}
