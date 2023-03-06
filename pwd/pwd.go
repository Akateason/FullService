/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-04 00:48:00
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-05 18:57:45
 * @FilePath: /FullService/pwd/pwd.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package pwd

import "time"

type Secret struct {
	Id         int64 // pk
	Title      string
	Account    string
	Password   string
	DetailInfo string
	Type       int
	ReadCount  int
	Pinyin     string
	ImageUrl   string
	IsDelete   bool
	Created    time.Time `xorm:"created"`
	Updated    time.Time `xorm:"updated"`
	SyncTime   time.Time `xorm:"syncTime"`
}
