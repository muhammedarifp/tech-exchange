package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/tech-exchange/payments/api/handlers"
	"github.com/muhammedarifp/tech-exchange/payments/api/middleware"
)

type ServeHTTP struct {
	engine *gin.Engine
}

func NewServeHTTP(user *handlers.UserPaymentHandler, admin *handlers.AdminPaymentHandler) *ServeHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())

	// user routes
	userAuthRoute := engine.Group("/api/v1/payments").Use(middleware.AuthMiddleware())
	{
		userAuthRoute.GET("/fetchplans", user.FetchPlans)
		userAuthRoute.POST("/create-subsc", user.CreateSubscription)
		userAuthRoute.DELETE("/cancel-subsc", user.CancelSubscription)
		userAuthRoute.POST("/verify-payment", user.VerifyPayment)
	}

	// admin routes
	adminAuthRoute := engine.Group("/api/v1/payments/admin").Use(middleware.AdminAuthMiddleware())
	{
		adminAuthRoute.GET("/fetch-plans", user.FetchPlans)
		adminAuthRoute.POST("/create-plan", admin.AddPlan)
		adminAuthRoute.DELETE("/remove-plan", admin.RemovePlan)
	}

	// Return engine
	return &ServeHTTP{
		engine: engine,
	}
}

func (s *ServeHTTP) Start() {
	s.engine.Run(":8003")
}
