package routes

import (
	"jerseyhub/pkg/api/handler"
	"jerseyhub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler, cartHandler *handler.CartHandler, paymentHandler *handler.PaymentHandler, wishlisthandler *handler.WishlistHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)
	engine.GET("/forgot-password", userHandler.ForgotPasswordSend)
	engine.POST("/forgot-password", userHandler.ForgotPasswordVerifyAndChange)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)

	payment := engine.Group("/payment")
	{
		payment.GET("/razorpay", paymentHandler.MakePaymentRazorPay)
		payment.GET("/update_status", paymentHandler.VerifyPayment)
	}

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
			home.POST("/wishlist/add", wishlisthandler.AddToWishlist)

		}

		profile := engine.Group("/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddresses)
			profile.POST("/address/add", userHandler.AddAddress)
			profile.GET("/get-link", userHandler.GetMyReferenceLink)

			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.DELETE("", orderHandler.CancelOrder)
				orders.PUT("/return", orderHandler.ReturnOrder)
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

		wishlist := engine.Group("/wishlist")
		{
			wishlist.GET("/", wishlisthandler.GetWishList)
			wishlist.DELETE("/remove", wishlisthandler.RemoveFromWishlist)
		}

		checkout := engine.Group("/check-out")
		{
			checkout.GET("", cartHandler.CheckOut)
			checkout.POST("/order", orderHandler.OrderItemsFromCart)
		}

	}

}
