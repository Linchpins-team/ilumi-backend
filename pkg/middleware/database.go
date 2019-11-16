package middleware

import (
	"github.com/Linchpins-team/ilumi-backend/pkg/database"
	"github.com/gin-gonic/gin"
)

// AddDatabase add the db to the context
func AddDatabase(db database.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
