package routes

import (
	"jerseyhub/pkg/api/handler"
	"jerseyhub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler, cartHandler *handler.CartHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)
	engine.GET("/forgot-password", userHandler.ForgotPasswordSend)
	engine.POST("/forgot-password", userHandler.ForgotPasswordVerifyAndChange)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)

	engine.Use(middleware.UserAuthMiddleware)
	{

		search := engine.Group("/search")
		{
			search.POST("/", inventoryHandler.SearchProducts)
		}

		home := engine.Group("/home")
		{
			home.GET("/products", inventoryHandler.ListProducts)
			home.GET("/products/details", inventoryHandler.ShowIndividualProducts)
			home.POST("/add-to-cart", cartHandler.AddToCart)

		}

		profile := engine.Group("/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddresses)
			profile.POST("/address/add", userHandler.AddAddress)

			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.DELETE("", orderHandler.CancelOrder)
				// orders.GET("/view",userHandler.OrderDetails)
			}

			edit := profile.Group("/edit")
			{
				edit.PUT("/name", userHandler.EditName)
				edit.PUT("/email", userHandler.EditEmail)
				edit.PUT("/phone", userHandler.EditPhone)
			}

			security := profile.Group("/security")
			{
				security.PUT("/change-password", userHandler.ChangePassword)
			}
		}

		cart := engine.Group("/cart")
		{
			cart.GET("/", userHandler.GetCart)
			cart.DELETE("/remove", userHandler.RemoveFromCart)
			cart.PUT("/updateQuantity/plus", userHandler.UpdateQuantityAdd)
			cart.PUT("/updateQuantity/minus", userHandler.UpdateQuantityLess)
			// hello
		}

		checkout := engine.Group("/check-out")
		{
			checkout.GET("", cartHandler.CheckOut)
			checkout.POST("/order", orderHandler.OrderItemsFromCart)
		}

	}
}
