/*
 * @Author: Mamba24 akateason@qq.com
 * @Date: 2023-03-04 00:30:43
 * @LastEditors: Mamba24 akateason@qq.com
 * @LastEditTime: 2023-03-06 22:50:45
 * @FilePath: /FullService/webservice/webservice.go
 * @Description: gin web manager
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package webservice

import (
	"FullService/dbManager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupWebservice() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID) // 定向id

	router.GET("/users", getAllUsers)

	router.Run("localhost:8080")
}

func getAllUsers(c *gin.Context) {
	list := dbManager.SelectAll()
	c.IndentedJSON(http.StatusOK, list)
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
