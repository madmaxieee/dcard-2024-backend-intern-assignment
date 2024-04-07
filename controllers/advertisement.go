package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request body
// '{
// "title": "AD 55",
// "startAt": "2023-12-10T03:00:00.000Z",
// "endAt": "2023-12-31T16:00:00.000Z",
// "conditions": [
// {
// "ageStart": 20,
// "ageEnd": 30,
// "country: ["TW", "JP"],
// "platform": ["android", "ios"]
// }
// ]
// }'

func CreateAdvertisement(c *gin.Context) {
	var body struct {
		Title      string `json:"title"`
		StartAt    string `json:"startAt"`
		EndAt      string `json:"endAt"`
		Conditions []struct {
			AgeStart int      `json:"ageStart"`
			AgeEnd   int      `json:"ageEnd"`
			Country  []string `json:"country"`
			Platform []string `json:"platform"`
		} `json:"conditions"`
	}

	c.BindJSON(&body)

	fmt.Printf("%+v", body)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}
