/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-04 00:30:43
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-07 01:19:54
 * @FilePath: /FullService/webservice/webservice.go
 * @Description: gin web manager
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package webservice

import (
	model "FullService/Model"
	"FullService/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupWebservice() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	// router.GET("/albums/:id", getAlbumByID) // 定向id

	router.GET("/users", getAllUsers)
	router.GET("/user", getUserByName)
	router.POST("/user", postUser)

	router.Run("localhost:8080")
}

func getAllUsers(c *gin.Context) {
	list := person.SelectAll()
	data := interface{}(list)
	result := model.ResultData{Data: data, Success: true}
	c.IndentedJSON(http.StatusOK, result)
}

func postUser(c *gin.Context) {
	var newUser person.Person
	if err := c.BindJSON(&newUser); err != nil {
		result := model.ResultData{Success: false, ErrorMessage: "参数错误:" + err.Error()}
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}

	if err := person.Upsert(newUser); err != nil {
		result := model.ResultData{Success: false, ErrorMessage: "表插入错误:" + err.Error()}
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}

	result := model.ResultData{Success: true}
	c.IndentedJSON(http.StatusCreated, result)
}

func getUserByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		result := model.ResultData{Success: false, ErrorMessage: "参数错误:"}
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}

	person := person.GetByName(name)
	result := model.ResultData{Data: person, Success: true}
	c.IndentedJSON(http.StatusOK, result)
}
