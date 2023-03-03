/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-04 00:47:52
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-04 01:40:39
 * @FilePath: /FullService/person/User.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package person

import "time"

type Person struct {
	Id      int64
	Name    string
	Sex     string
	Age     int
	Passwd  string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
