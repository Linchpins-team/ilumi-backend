package router

import (
	"fmt"

	"github.com/Linchpins-team/ilumi-backend/pkg/config"
	"github.com/Linchpins-team/ilumi-backend/pkg/database"
	"github.com/Linchpins-team/ilumi-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// RunServer starts a Gin Engine
func RunServer(conf config.Config, db database.DB) {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.Use(middleware.AddDatabase(db))
		users := v1.Group("/users")
		{
			users.GET("/:uid")
			users.POST("/:uid")
			users.DELETE("/:uid")
		}
		missions := v1.Group("/missions")
		{
			missions.GET("/")
			missions.POST("/")
			missions.POST("/:mid")
			missions.DELETE("/:mid")
			missions.GET("/:mid")
			answers := missions.Group("/:mid/answers")
			{
				answers.GET("/")
				answers.POST("/")
				answers.POST("/:aid")
				answers.DELETE("/:aid")
			}
		}
		notifications := v1.Group("/notifications")
		{
			notifications.GET("/")
			notifications.POST("/")
		}
		v1.POST("/join")
		v1.POST("/login")
		v1.GET("/logout")
	}
	fmt.Printf("Server has started on http://%s:%s", conf.Host, conf.Port)
	if err := router.Run(":" + conf.Port); err != nil {
		panic(err)
	}
}
