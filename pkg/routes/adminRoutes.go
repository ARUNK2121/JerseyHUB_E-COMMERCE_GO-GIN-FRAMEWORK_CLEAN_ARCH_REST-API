package routes

import (
	"jerseyhub/pkg/api/handler"
	"jerseyhub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, inventoryHandler *handler.InventoryHandler, userHandler *handler.UserHandler, categoryHandler *handler.CategoryHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, offerHandler *handler.OfferHandler) {

	engine.POST("/adminlogin", adminHandler.LoginHandler)
	// api := router.Group("/admin_panel", middleware.AuthorizationMiddleware)
	// api.GET("users", adminHandler.GetUsers)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.GET("", adminHandler.GetUsers)
			usermanagement.PUT("/block", adminHandler.BlockUser)
			usermanagement.PUT("/unblock", adminHandler.UnBlockUser)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.GET("", categoryHandler.GetCategory)
			categorymanagement.POST("", categoryHandler.AddCategory)
			categorymanagement.PUT("", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("", categoryHandler.DeleteCategory)
		}

		inventorymanagement := engine.Group("/inventories")
		{
			inventorymanagement.GET("", inventoryHandler.ListProducts)
			inventorymanagement.POST("", inventoryHandler.AddInventory)
			inventorymanagement.DELETE("", inventoryHandler.DeleteInventory)
			inventorymanagement.PUT("/details", inventoryHandler.EditInventoryDetails)

			inventorymanagement.PUT("/:id/stock", inventoryHandler.UpdateInventory)
			inventorymanagement.PUT("/:id/image", inventoryHandler.UpdateProductImage)
			// inventorymanagement.PUT("/:id/details", inventoryHandler.UpdateProductDetails)
		}

		payment := engine.Group("/payment-method")
		{
			payment.POST("", adminHandler.NewPaymentMethod)
			payment.GET("", adminHandler.ListPaymentMethods)
			payment.DELETE("", adminHandler.DeletePaymentMethod)
		}

		orders := engine.Group("/orders")
		{
			orders.PUT("/status", orderHandler.EditOrderStatus)
			orders.GET("", orderHandler.AdminOrders)
		}

		coupons := engine.Group("/coupons")
		{
			coupons.GET("", couponHandler.GetAllCoupons)
			coupons.POST("", couponHandler.CreateNewCoupon)
			coupons.DELETE("", couponHandler.MakeCouponInvalid)
			//reactivation of coupons
			coupons.PUT("", couponHandler.ReActivateCoupon)
		}

		offers := engine.Group("/offers")
		{
			offers.GET("", offerHandler.GetOffers)
			offers.POST("", offerHandler.AddNewOffer)
			offers.DELETE("", offerHandler.MakeOfferExpire)
		}
	}

}
