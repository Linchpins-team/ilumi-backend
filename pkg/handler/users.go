package handler

import (
	"fmt"
	"github.com/Linchpins-team/ilumi-backend/pkg/database"
	"github.com/Linchpins-team/ilumi-backend/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func getDB(c *gin.Context) *gorm.DB {
	obj, ok := c.Get("db")
	if ok {
		if db, ok := obj.(database.DB); ok {
			return db.Gorm()
		}
	}
	panic(fmt.Errorf("db not found %v", obj))
}

func hashPassword(pwd string) []byte {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashed
}

// Join, the sign up function. require email, password ()
// auto login (add to session)
func Join(c *gin.Context) {
	var user model.User
	user.Email = c.PostForm("email")
	user.Password = hashPassword(c.PostForm("password"))
	db := getDB(c)
	if err := db.Create(&user).Error; err != nil {
		// Email already exist
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(403, data)
		return
	}
	data := gin.H{
		"id": user.ID,
	}
	c.JSON(200, data)
}
