package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type numbers struct {
	Num1 float64 `json:"num1`
	Num2 float64 `json:"num2`
}
type numResponseData struct {
	Add float64 `json:add`
	Mul float64 `json:mul`
	Sub float64 `json:sub`
	Div float64 `json:div`
}

func process(numdata numbers) numResponseData {

	var numres numResponseData
	numres.Add = numdata.Num1 + numdata.Num2
	numres.Mul = numdata.Num1 * numdata.Num2
	numres.Sub = numdata.Num1 - numdata.Num2
	numres.Div = numdata.Num1 / numdata.Num2

	return numres
}

type respons struct {
	ResponseCode int    `json:responseCode`
	Message      string `json:message`
}

func main() {
	router := gin.Default()
	router.GET("/", homePage)

	router.POST("/calc", calc)

	router.Run("localhost:8088")
}

func homePage(c *gin.Context) {

	rCode := c.Writer.Status()
	rMessage := ""

	if rCode != http.StatusOK {
		rMessage = "Något gick fel"
	}
	r := respons{
		ResponseCode: rCode,
		Message:      rMessage,
	}

	c.IndentedJSON(r.ResponseCode, r)
}

func calc(c *gin.Context) {
	var numData numbers
	var numResData numResponseData

	if err := c.BindJSON(&numData); err != nil {
		return
	}

	numResData = process(numData)

	c.IndentedJSON(http.StatusCreated, numResData)
}
