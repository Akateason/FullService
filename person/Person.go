/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-04 00:47:52
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-06 23:38:51
 * @FilePath: /FullService/person/Person.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package person

type Person struct {
	Id      int64  `json:"id"` // pk
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	Passwd  string `json:"passwd"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}
