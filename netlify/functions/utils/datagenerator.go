package datagenerator

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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

func GetFruits(c *gin.Context) {
	var fruitJson []Fruit

	log.SetFlags(0)
	log.SetPrefix("Message Information: ")

	apiURL := "https://fruityvice.com/api/fruit/all"
	resp, err := http.Get(apiURL)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	errr := json.Unmarshal(body, &fruitJson)
	if errr != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	_ = ioutil.WriteFile("fruitsdatabase.json", body, 0644)

	c.IndentedJSON(http.StatusOK, fruitJson)
}
