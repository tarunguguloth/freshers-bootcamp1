package Controllers

import (
	"fmt"
	"net/http"


	"freshers_bootcamp/Day4/Models"
	"github.com/gin-gonic/gin"
)

//Create Order
func CreateOrder(c *gin.Context) {
	var user Models.Order
	c.BindJSON(&user)
	err := Models.CreateOrder(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//getting order by id
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//getting all orders
func GetOrders(c *gin.Context) {
	var order []Models.Product
	err := Models.GetOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}