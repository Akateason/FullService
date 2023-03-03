/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2022-12-04 22:20:26
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-04 00:33:41
 * @FilePath: /FullService/main.go
 * @Description: Web+DB
 *
 * Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved.
 */

package main

import (
	"FullService/dbManager"
	"FullService/webservice"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	openServices()
}

func openServices() {
	// database
	dbManager.SetupDatabase()

	// web
	webservice.SetupWebservice()
}
