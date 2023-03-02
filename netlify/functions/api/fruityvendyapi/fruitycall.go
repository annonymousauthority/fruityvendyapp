package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fruityvendy.com/datagenerator"
	"github.com/gin-gonic/gin"
)

type Fruit struct {
	Genus      string         `json:"genus"`
	Name       string         `json:"name"`
	Id         int            `json:"id"`
	Family     string         `json:"family"`
	Order      string         `json:"order"`
	Nutritions map[string]int `json:"nutritions"`
}

func main() {
	r := gin.Default()
	r.GET("/fruits", datagenerator.GetFruits)
	r.GET("/getFruits", extractFruitJson)
	r.Run("localhost:3000")
}

func extractFruitJson(c *gin.Context) {
	var fruitJson []Fruit
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")

	fruits, err := ioutil.ReadFile("./fruitsdatabase.json")
	if err != nil {
		// log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error Message: ": err})
	}

	_ = json.Unmarshal(fruits, &fruitJson)

	c.IndentedJSON(http.StatusFound, fruitJson)
}
