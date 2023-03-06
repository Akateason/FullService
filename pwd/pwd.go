/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-04 00:48:00
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-06 23:48:34
 * @FilePath: /FullService/pwd/pwd.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package pwd

type Secret struct {
	Id         int64  `json:"id"` // pk
	Title      string `json:"title"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	DetailInfo string `json:"detailInfo"`
	Type       int    `json:"type"`
	ReadCount  int    `json:"readCount"`
	Pinyin     string `json:"pinyin"`
	ImageUrl   string `json:"imageUrl"`
	IsDelete   bool   `json:"isDelete"`
	UserId     int64  `json:"userid"`
	Created    int64  `xorm:"created" json:"created"`
	Updated    int64  `xorm:"updated" json:"updated"`
	SyncTime   int64  `xorm:"syncTime"`
}
