package http

import (
	"github.com/gin-gonic/gin"

	handler "jerseyhub/pkg/api/handler"
	"jerseyhub/pkg/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, otpHandler *handler.OtpHandler, orderHandler *handler.OrderHandler, cartHandler *handler.CartHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// engine.POST("users/signup", userHandler.UserSignUp)
	// engine.POST("users/login", userHandler.LoginHandler)

	// engine.GET("inventory/productdetails", inventoryHandler.ShowIndividualProducts)
	// engine.GET("inventory/productlist", inventoryHandler.ListProducts)

	// engine.POST("otp/send-otp", otpHandler.SendOTP)
	// engine.POST("otp/verify-otp", otpHandler.VerifyOTP)

	routes.UserRoutes(engine.Group("/user"), userHandler, otpHandler, inventoryHandler, orderHandler, cartHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, inventoryHandler, userHandler, categoryHandler, orderHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
