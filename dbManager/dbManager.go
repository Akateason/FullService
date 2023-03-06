/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-03 23:27:12
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-07 01:13:56
 * @FilePath: /FullService/dbManager/dbManager.go
 * @Description: ORM Database Manager
 *
 * Copyright (c) 2023 by Mamba24 akateason@qq.com, All Rights Reserved.
 */
package dbManager

import (
	"fmt"

	"github.com/xormplus/xorm"
)

var Engine *xorm.Engine

// setup db
func SetupDatabase() {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("🐒❌db engine link failed")
		fmt.Println(err)
	}
	fmt.Println("🐒db engine link success")
}

// bind model
func DbBindClass(beans ...interface{}) {
	err := Engine.Sync2(beans...)
	if err != nil {
		fmt.Println("🐒❌bind model failed")
		fmt.Println(err)
		return
	}
	fmt.Printf("🐒 model %T sync success\n", beans...)
}
