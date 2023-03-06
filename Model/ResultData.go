/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-06 23:52:12
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-07 00:05:11
 * @FilePath: /FullService/Model/ResultData.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package model

type ResultData struct {
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"errorMessage"`
	ErrorCode    int         `json:"errorCode"`
	ErrorDomain  string      `json:"errorDomain"`
}
