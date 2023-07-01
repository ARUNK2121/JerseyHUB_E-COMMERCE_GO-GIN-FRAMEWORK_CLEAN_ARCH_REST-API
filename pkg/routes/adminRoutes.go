package routes

import (
	"jerseyhub/pkg/api/handler"
	"jerseyhub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, inventoryHandler *handler.InventoryHandler, userHandler *handler.UserHandler, categoryHandler *handler.CategoryHandler, orderHandler *handler.OrderHandler) {

	engine.POST("/adminlogin", adminHandler.LoginHandler)
	// api := router.Group("/admin_panel", middleware.AuthorizationMiddleware)
	// api.GET("users", adminHandler.GetUsers)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.POST("/block", adminHandler.BlockUser)
			usermanagement.POST("/unblock", adminHandler.UnBlockUser)
			usermanagement.GET("/getusers", adminHandler.GetUsers)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.POST("/add", categoryHandler.AddCategory)
			categorymanagement.PUT("/update", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("/delete", categoryHandler.DeleteCategory)
		}

		inventorymanagement := engine.Group("/inventories")
		{
			inventorymanagement.POST("/add", inventoryHandler.AddInventory)
			inventorymanagement.PUT("/update", inventoryHandler.UpdateInventory)
			inventorymanagement.DELETE("/delete", inventoryHandler.DeleteInventory)
		}

		payment := engine.Group("/payment")
		{
			payment.POST("/payment-method/new", adminHandler.NewPaymentMethod)
		}

		orders := engine.Group("/orders")
		{
			orders.PUT("/edit/status", orderHandler.EditOrderStatus)
			orders.GET("", orderHandler.AdminOrders)
		}

	}

}
