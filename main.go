package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// выдает список альбомов
// gin.Context валидирует json
// Context IndentedJSON переводит структуру в json и добавляет в ответ
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// добавляет из json
func postAlbums(c *gin.Context) {
	var newAlbum album

	//вызывает BindJSON для получения json в newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//добавляет новый альбом в слайс к существующим
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//func delAlbumByID(c *gin.Context) {
//	id := c.Param("id")
//
//	for _, a := range albums {
//		if a.ID == id {
//			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album deleted"})
//			return
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found, cant delete"})
//
//}

func main() {
	router := gin.Default()          //gin.Default() поднимает роутер
	router.GET("/albums", getAlbums) //получает http метод и путь /albums через хэндлер
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	//router.DELETE("/albums/del/:id", delAlbumByID)

	router.Run("localhost:8081") //старт сервера
}
