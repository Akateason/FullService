/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-03 23:27:12
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-07 01:06:58
 * @FilePath: /FullService/dbManager/dbManager.go
 * @Description: ORM Database Manager
 *
 * Copyright (c) 2023 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package dbManager

import (
	"FullService/person"
	"fmt"

	"github.com/Akateason/GoScriptPlayground/earth"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

// setup db
func SetupDatabase() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("🐒❌db engine link failed")
		fmt.Println(err)
	}
	fmt.Println("🐒db engine link success")
}

// bind model
func DbBindClass(beans ...interface{}) {
	err := engine.Sync2(beans...)
	if err != nil {
		fmt.Println("🐒❌bind model failed")
		fmt.Println(err)
		return
	}
	fmt.Printf("🐒 model %T sync success\n", beans...)
}

// select all
func SelectAll() []person.Person {
	users := make([]person.Person, 0)
	err := engine.Find(&users)
	if err != nil {
		fmt.Println("🐒❌select all failed")
		fmt.Println(err)
	}
	fmt.Println("🐒 select all")

	var _list []interface{}
	for _, v := range users {
		_list = append(_list, v)
	}
	earth.PrintArray([]interface{}(_list))
	return users
}

// select by id
func GetByName(name string) interface{} {
	var user person.Person
	has, err := engine.Where("name = ?", name).Desc("id").Get(&user)
	if err != nil {
		fmt.Println("🐒❌ Exist failed")
		fmt.Println(err)
		return nil
	}
	if has {
		return user
	}
	return nil
}

// insert, update if exsist
func Upsert(user person.Person) error {
	has, err := engine.Where("Id = ?", user.Id).Exist(&person.Person{})
	if err != nil {
		fmt.Println("🐒❌ Exist failed")
		fmt.Println(err)
		return err
	}
	if has { // update
		affected, err := engine.Update(&user, &person.Person{Id: user.Id})
		// UPDATE user SET ... Where Id = ?
		if err != nil {
			fmt.Println("🐒❌ update failed")
			fmt.Println(err)
			return err
		}
		fmt.Printf("🐒 %d %T update success\n", affected, user)
	} else { // insert
		affected, err := engine.Insert(&user)
		if err != nil {
			fmt.Println("🐒❌ insert failed")
			fmt.Println(err)
			return err
		}
		fmt.Printf("🐒 %d %T insert success\n", affected, user)
	}
	return nil
}
