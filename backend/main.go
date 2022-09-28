package main

import (
	//"fmt"
	"net/http"
	"os"

	//"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func handler(c *gin.Context) {
	//fmt.Println(db)
	//fmt.Println(Datas)
	c.JSON(http.StatusOK, gin.H{
		//"data": db,
		"data": Datas,
	})
}

type Data struct {
	Text string `json:"text"`
}

var Datas []Data

// type db struct {
// 	data []Data
// 	sync.Mutex
// }

func (d Data) tambahData() {
	Datas = append(Datas, d)
}

// var db []string

// type DataRequest struct {
// 	Text string `json:"text"`
// }

func postHandler(c *gin.Context) {
	//var newData DataRequest
	var newData Data
	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//fmt.Println(newData)
	//db = append(db, newData.Text)
	//Datas = append(Datas, newData)
	newData.tambahData()

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": newData})
}

func main() {
	//db = make([]string, 0)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/", handler)
	r.POST("/send", postHandler)
	r.Run(":" + os.Getenv("PORT"))
}
