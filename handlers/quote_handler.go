package handlers 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"remood/models"

	"github.com/gin-gonic/gin"
)

func CreateManyQuotes(ctx *gin.Context) {
	var quotes []models.Quote
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read quotes info",
			"error":   true,
		})
		return
	}

	err = json.Unmarshal(body, &quotes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, 
			models.ErrorResponse("Fail to read quotes info"))
		return
	}

	var q models.Quote
	if quotes, err = q.CreateMany(quotes); err != nil {
		ctx.JSON(http.StatusInternalServerError, 
			models.ErrorResponse("Fail to create many quotes"))
	}

	ctx.JSON(http.StatusOK, 
		models.SuccessResponse("Create quotes succesfully", gin.H{
			"quotes": quotes,
		}))
}

func GetRandomQuotes(ctx *gin.Context) {
	param := ctx.Query("number")
	number, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, 
			models.ErrorResponse("Wrong parameter"))
	}

	var q models.Quote
	quotes, err := q.GetRandom(number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse("Fail to get quotes"))
	}

	ctx.JSON(http.StatusOK, 
		models.SuccessResponse("Get random quotes successfully", gin.H{
			"quotes": quotes, 
		}))
}