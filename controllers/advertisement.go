package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAdvertisement(c *gin.Context) {
	type Condition struct {
		AgeStart int      `json:"ageStart"`
		AgeEnd   int      `json:"ageEnd"`
		Country  []string `json:"country"`
		Platform []string `json:"platform"`
	}

	type Body struct {
		Title      string      `json:"title"`
		StartAt    string      `json:"startAt"`
		EndAt      string      `json:"endAt"`
		Conditions []Condition `json:"conditions"`
	}

	var err error
	var body Body
	err = c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", body)

	c.JSON(http.StatusOK, gin.H{
		"message": "Advertisement created successfully",
	})
}
