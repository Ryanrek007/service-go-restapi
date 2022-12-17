package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"service-go-restapi/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var staff []models.Staff

	models.DB.Find(&staff)
	c.JSON(http.StatusOK, gin.H{
		"Data Karyawan": staff,
	})
}

func WeatherResponse(c *gin.Context) {
	const CITY = "Jakarta"
	var weather models.WeatherResponse

	APIKEY := os.Getenv("APIKEY")
	URL := "https://api.openweathermap.org/data/2.5/weather?q=" + CITY + "&APPID=" + APIKEY + "&units=metric"

	response, err := http.Get(URL)
	fmt.Println("URL that you hit: ", URL)

	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	er := json.Unmarshal(jsonBytes, &weather)
	if er != nil {
		log.Fatal(er)
	}

	c.JSON(http.StatusOK, gin.H{
		"Forecast Weather": weather,
	})

}
