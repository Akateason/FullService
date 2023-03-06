/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-12-04 22:20:26
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-06 22:47:27
 * @FilePath: /FullService/main.go
 * @Description: Web+DB
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */

package main

import (
	"FullService/dbManager"
	"FullService/person"
	"FullService/pwd"
	"FullService/webservice"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	openServices()

	forUnitTest() // TODO: delete
}

func openServices() {
	// database
	dbManager.SetupDatabase()
	dbManager.DbBindClass(new(person.Person))
	dbManager.DbBindClass(new(pwd.Secret))

	// web
	webservice.SetupWebservice()
}

func forUnitTest() {
	// user1 := new(person.Person)
	// user1.Id = 4
	// user1.Name = "yes"
	// dbManager.Upsert(*user1)

	dbManager.SelectAll()

	// web
	// webservice.SetupWebservice()
}
