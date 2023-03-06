/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-04 00:47:52
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-07 01:17:44
 * @FilePath: /FullService/person/Person.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package person

import (
	"FullService/dbManager"
	"fmt"

	"github.com/Akateason/GoScriptPlayground/earth"
)

type Person struct {
	Id      int64  `json:"id"` // pk
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	Passwd  string `json:"passwd"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}

// db, select all
func SelectAll() []Person {
	users := make([]Person, 0)
	err := dbManager.Engine.Find(&users)
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

// db, select by id
func GetByName(name string) interface{} {
	var user Person
	has, err := dbManager.Engine.Where("name = ?", name).Desc("id").Get(&user)
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

// db, insert, update if exsist
func Upsert(user Person) error {
	has, err := dbManager.Engine.Where("Id = ?", user.Id).Exist(&Person{})
	if err != nil {
		fmt.Println("🐒❌ Exist failed")
		fmt.Println(err)
		return err
	}
	if has { // update
		affected, err := dbManager.Engine.Update(&user, &Person{Id: user.Id})
		// UPDATE user SET ... Where Id = ?
		if err != nil {
			fmt.Println("🐒❌ update failed")
			fmt.Println(err)
			return err
		}
		fmt.Printf("🐒 %d %T update success\n", affected, user)
	} else { // insert
		affected, err := dbManager.Engine.Insert(&user)
		if err != nil {
			fmt.Println("🐒❌ insert failed")
			fmt.Println(err)
			return err
		}
		fmt.Printf("🐒 %d %T insert success\n", affected, user)
	}
	return nil
}
