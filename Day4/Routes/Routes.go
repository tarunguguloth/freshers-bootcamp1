package Routes

import (
	"freshers_bootcamp/Day4/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	cust := r.Group("/Ecommerce")
	{

		cust.POST("createcustomers", Controllers.CreateCustomer)
		cust.DELETE("deletecustomer/:id", Controllers.DeleteCustomer)
		cust.GET("customerssbyid/:id", Controllers.GetCustomerByID)
	}
	return r

	prod := r.Group("/Ecommerce")
	{
		prod.GET("allproducts", Controllers.GetProducts)
		prod.POST("createproducts", Controllers.CreateProduct)
		prod.DELETE("deleteproduct/:id", Controllers.DeleteProduct)
		prod.GET("productsbyid/:id", Controllers.GetProductByID)
		prod.PATCH("updateproduct/:id", Controllers.UpdateProduct)
	}
	return r


	ord := r.Group("/Ecommerce")
	{
		ord.POST("createorders", Controllers.CreateOrder)
		ord.GET("ordersbyid/:id", Controllers.GetOrderByID)
		ord.GET("allorders", Controllers.GetOrders)
	}
	return r
}